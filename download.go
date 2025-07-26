package gcurl

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func SimpleDownload(downloadUrl string, saveFilename string) error {
	out, err := os.Create(saveFilename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
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
	//defer resp.Body.Close()

	body, err := ioutil.ReadAll(respReader)
	if err != nil {
		return err
	}
	if _, err = out.Write(body); err != nil {
		return err
	}
	return nil
}

func SimpleDownloadV2(downloadUrl string, saveFilename string) error {
	out, err := os.Create(saveFilename)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
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
	_, err = io.Copy(out, respReader)
	if err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}
	return nil
}
