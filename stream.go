package gcurl

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func SendPostData4StreamV1(urlStr string, byteCon []byte) (content []byte, h http.Header, err error) {
	payload := strings.NewReader(string(byteCon))
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	req, err := http.NewRequest(http.MethodPost, urlStr, payload)
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/octet-stream")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	h = res.Header
	content, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
