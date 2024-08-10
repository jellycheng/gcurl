package gcurl

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// 上传文件
func SendUploadFile(headers map[string]string, apiUrl string, upfile string) (string, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(upfile)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("upfile", filepath.Base(upfile))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return "", errFile1
	}
	err := writer.Close()
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiUrl, payload)
	if err != nil {
		return "", err
	}
	for headerKey, headerVal := range headers {
		if headerKey != "" {
			req.Header.Add(headerKey, headerVal)
		}
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
