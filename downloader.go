package gcurl

import (
	"compress/flate"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"
	"time"
)

type Downloader struct {
	DownloadUrl     string            // 下载地址
	Size            int64             // 文件大小
	SaveFileName    string            // file name to save
	TempFileDownExt string            // 下载文件临时文件名后缀
	ChunkSizeMB     int               // 分块大小，单位MB, 默认0-不使用分块下载
	RetryTimes      int               // 重试次数, 默认0-不重试
	RequestTimeout  time.Duration     // 超时时间
	Headers         map[string]string // 请求头,User-Agent、Referer、Content-Type、Cookie、Authorization等
	RawCookie       string            // cookie
	ExtInfo         interface{}       // 扩展信息

	DownSizeCallbackFunc func(*Downloader, int64)                                   // 下载大小回调
	SaveCallbackFunc     func(*Downloader, *http.Response, *os.File) (int64, error) // 保存文件回调

}

func NewDownloader() *Downloader {
	ret := &Downloader{
		TempFileDownExt: ".download",
		ChunkSizeMB:     0,
		RetryTimes:      0,
		RequestTimeout:  15 * time.Minute,
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",
		},
		SaveCallbackFunc: DefaultSaveCallback,
	}
	return ret
}

func (m *Downloader) Save() error {
	fileSize, exists, err := FileExistAndSize(m.SaveFileName)
	if err != nil {
		return err
	}
	if exists && fileSize == m.Size { // 文件已经下载完成
		if m.DownSizeCallbackFunc != nil {
			m.DownSizeCallbackFunc(m, fileSize)
		}
		return nil
	}
	// 继续下载文件
	tempFilePath := m.SaveFileName + m.TempFileDownExt
	tempFileSize, _, err := FileExistAndSize(tempFilePath)
	if err != nil {
		return err
	}
	var (
		osFileObj *os.File
		fileError error
	)
	if tempFileSize > 0 {
		m.Headers["Range"] = fmt.Sprintf("bytes=%d-", tempFileSize)
		osFileObj, fileError = os.OpenFile(tempFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		if m.DownSizeCallbackFunc != nil {
			m.DownSizeCallbackFunc(m, tempFileSize)
		}
		//downloader.bar.Add64(tempFileSize)
	} else {
		osFileObj, fileError = os.Create(tempFilePath)
	}
	if fileError != nil {
		return fileError
	}
	defer func() {
		_ = osFileObj.Close()
		if err == nil {
			_ = os.Rename(tempFilePath, m.SaveFileName) //重命名文件
		}
	}()

	if m.ChunkSizeMB > 0 {
		// 分块下载
		var start, end, perChunkSize int64
		perChunkSize = int64(m.ChunkSizeMB) * 1024 * 1024 // 每个分块大小
		if m.Size == 0 {
			s1, _ := m.GetFileSize()
			m.Size = s1
		}
		remainingSize := m.Size // 剩下大小
		if tempFileSize > 0 {
			start = tempFileSize
			remainingSize -= tempFileSize
		}
		chunk := remainingSize / perChunkSize // 块数
		if remainingSize%perChunkSize != 0 {
			chunk++
		}
		errorSlice := make([]error, 0)
		var i int64 = 1
		for ; i <= chunk; i++ {
			end = start + perChunkSize - 1
			m.Headers["Range"] = fmt.Sprintf("bytes=%d-%d", start, end)
			written, err := m.writeFile(osFileObj)
			_ = written
			if err != nil {
				errorSlice = append(errorSlice, err)
				break
			}
			start = end + 1
		}
		if len(errorSlice) != 0 {
			return errorSlice[0]
		}
	} else {
		// 不分块下载
		written, err := m.writeFile(osFileObj)
		_ = written
		return err
	}
	return nil
}

func (m *Downloader) writeFile(osFileObj *os.File) (int64, error) {
	resp, err := m.Request(http.MethodGet, nil)
	if err != nil {
		return 0, err
	}
	if m.SaveCallbackFunc != nil {
		return m.SaveCallbackFunc(m, resp, osFileObj)
	}
	defer resp.Body.Close()
	written, err := io.Copy(osFileObj, resp.Body)
	if m.DownSizeCallbackFunc != nil {
		m.DownSizeCallbackFunc(m, written)
	}
	return written, err
}

// 发起请求
func (m *Downloader) Request(method string, body io.Reader) (*http.Response, error) {
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DisableCompression:  true,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   m.RequestTimeout,
		Jar:       jar,
	}

	req, err := http.NewRequest(method, m.DownloadUrl, body)
	headers := m.Headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if _, ok := headers["Referer"]; !ok {
		req.Header.Set("Referer", m.DownloadUrl)
	}
	if m.RawCookie != "" {
		cookies, _ := ParseCookie4String(m.RawCookie)
		if len(cookies) > 0 {
			for _, c := range cookies {
				req.AddCookie(c)
			}
		} else {
			req.Header.Set("Cookie", m.RawCookie)
		}
	}
	var (
		resp         *http.Response
		requestError error
	)
	for i := 1; ; i++ {
		resp, requestError = client.Do(req)
		if requestError == nil && resp.StatusCode < 400 {
			break
		} else if i >= m.RetryTimes {
			var err error
			if requestError != nil {
				err = fmt.Errorf("request error: %v", requestError)
			} else {
				err = fmt.Errorf("%s request error: status code %d", m.DownloadUrl, resp.StatusCode)
			}
			return nil, err
		}
		time.Sleep(1 * time.Second)
	}
	return resp, nil
}

func (m *Downloader) GetResponseHeaders() (http.Header, error) {
	res, err := m.Request(http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res.Header, nil
}

func (m *Downloader) GetFileSize() (int64, error) {
	headerObj, err := m.GetResponseHeaders()
	if err != nil {
		return 0, err
	}
	s := headerObj.Get("Content-Length")
	if s == "" {
		return 0, errors.New("Content-Length不存在")
	}
	size, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func (m *Downloader) GetContentType() (string, string, error) {
	headerObj, err := m.GetResponseHeaders()
	if err != nil {
		return "", "", err
	}
	s := headerObj.Get("Content-Type")
	// text/html; charset=utf-8
	return strings.Split(s, ";")[0], s, nil
}

func DefaultSaveCallback(m *Downloader, resp *http.Response, osFileObj *os.File) (int64, error) {
	var respReader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		respReader, _ = gzip.NewReader(resp.Body)
	case "deflate":
		respReader = flate.NewReader(resp.Body)
	default:
		respReader = resp.Body
	}
	defer respReader.Close()
	written, err := io.Copy(osFileObj, respReader)
	if m.DownSizeCallbackFunc != nil {
		m.DownSizeCallbackFunc(m, written)
	}
	return written, err
}
