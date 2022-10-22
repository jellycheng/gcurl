package gcurl

import (
	"net"
	"net/http"
	"strings"
)

type ResponseBody []byte

func (me ResponseBody) GetContents() string {
	return string(me)
}

func (me ResponseBody) String() string {
	return string(me)
}

func (me ResponseBody) ToString() string {
	return string(me)
}

func (me ResponseBody) ToByte() []byte {
	return me
}

type Response struct {
	resp *http.Response
	req  *http.Request
	body []byte
	err  error
}

func (me *Response) GetRequest() *http.Request {
	return me.req
}

func (me *Response) GetBody() (ResponseBody, error) {
	return ResponseBody(me.body), me.err
}

func (me *Response) GetStatusCode() int {
	return me.resp.StatusCode
}

func (me *Response) IsTimeout() bool {
	if me.err == nil {
		return false
	}
	if netErr, ok := me.err.(net.Error); !ok {
		return false
	} else if netErr.Timeout() {
		return true
	}
	return false
}

// GetHeaders 获取所有请求头，也可通过 对象.GetRequest().Header获取原生http.Header类型
func (me *Response) GetHeaders() map[string][]string {
	return me.resp.Header
}

func (me *Response) GetHeaderSlice(name string) []string {
	headers := me.resp.Header
	for k, v := range headers {
		if strings.ToLower(name) == strings.ToLower(k) {
			return v
		}
	}
	return nil
}

// GetHeader 不区分大小写获取请求头内容
func (me *Response) GetHeader(name string) string {
	header := me.GetHeaderSlice(name)
	if len(header) > 0 {
		return header[0]
	}
	return ""
}

// HasHeader 不区分大小写判断请求头是否存在
func (me *Response) HasHeader(name string) bool {
	headers := me.resp.Header
	for k := range headers {
		if strings.ToLower(name) == strings.ToLower(k) {
			return true
		}
	}
	return false
}
