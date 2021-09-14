package gcurl

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type Request struct {
	options Options
	cli  *http.Client
	req  *http.Request
	body io.Reader
}

func (r *Request) Get(uri string, opts ...Options) (*Response, error) {
	return r.Request(http.MethodGet, uri, opts...)
}

func (r *Request) Post(uri string, opts ...Options) (*Response, error) {
	return r.Request(http.MethodPost, uri, opts...)
}

func (r *Request) Put(uri string, opts ...Options) (*Response, error) {
	return r.Request(http.MethodPut, uri, opts...)
}

func (r *Request) Patch(uri string, opts ...Options) (*Response, error) {
	return r.Request(http.MethodPatch, uri, opts...)
}

func (r *Request) Delete(uri string, opts ...Options) (*Response, error) {
	return r.Request(http.MethodDelete, uri, opts...)
}

func (r *Request) Options(uri string, opts ...Options) (*Response, error) {
	return r.Request(http.MethodOptions, uri, opts...)
}

func (r *Request) Request(method, uri string, opts ...Options) (*Response, error) {
	if len(opts) > 0 {
		r.options = opts[0]
	}
	if r.options.Headers == nil {
		r.options.Headers = make(map[string]interface{})
	}

	switch method {
	case http.MethodGet, http.MethodDelete:
		req, err := http.NewRequest(method, uri, nil)
		if err != nil {
			return nil, err
		}
		r.req = req
	case http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions:
		r.parseBody()
		req, err := http.NewRequest(method, uri, r.body)
		if err != nil {
			return nil, err
		}
		r.req = req
	default:
		return nil, errors.New("invalid request method")
	}

	r.parseOptions()
	r.parseClient()
	r.parseQuery()
	r.parseHeaders()
	r.parseCookies()

	if r.options.Debug {
		dump, err := httputil.DumpRequest(r.req, true)
		if err == nil {
			fmt.Printf("\n%s\n\n", dump)
		}
	}

	_resp, err := r.cli.Do(r.req)

	resp := &Response{
					resp: _resp,
					req:  r.req,
					err:  err,
				}

	if err == nil {
		body, err := ioutil.ReadAll(_resp.Body)
		_resp.Body.Close()

		resp.body = body
		resp.err = err
	}

	if err != nil {
		if r.options.Debug {
			fmt.Println(err)
		}

		return resp, err
	}

	if r.options.Debug {
		body, _ := resp.GetBody()
		fmt.Println(string(body))
	}

	return resp, nil
}

func (r *Request) parseOptions() {
	if r.options.Timeout == 0 {
		r.options.Timeout = 15 * time.Second
	}

}

func (r *Request) parseClient() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if r.options.Proxy != "" {
		proxy, err := url.Parse(r.options.Proxy)
		if err == nil {
			tr.Proxy = http.ProxyURL(proxy)
		}
	}

	r.cli = &http.Client{
		Timeout:   r.options.Timeout,
		Transport: tr,
	}
}

func (r *Request) parseQuery() {
	switch r.options.Query.(type) {
	case string:
		str := r.options.Query.(string)
		r.req.URL.RawQuery = str
	case map[string]interface{}:
		q := r.req.URL.Query()
		for k, v := range r.options.Query.(map[string]interface{}) {
			if vv, ok := v.(string); ok {
				q.Set(k, vv)
				continue
			}
			if vv, ok := v.([]string); ok {
				for _, vvv := range vv {
					q.Add(k, vvv)
				}
			}
		}
		r.req.URL.RawQuery = q.Encode()
	}
}

func (r *Request) parseCookies() {
	switch r.options.Cookies.(type) {
	case string:
		cookies := r.options.Cookies.(string)
		r.req.Header.Add("Cookie", cookies)
	case map[string]string:
		cookies := r.options.Cookies.(map[string]string)
		for k, v := range cookies {
			r.req.AddCookie(&http.Cookie{
				Name:  k,
				Value: v,
			})
		}
	case []*http.Cookie:
		cookies := r.options.Cookies.([]*http.Cookie)
		for _, cookie := range cookies {
			r.req.AddCookie(cookie)
		}
	}
}

func (r *Request) parseHeaders() {
	if r.options.Headers != nil {
		for k, v := range r.options.Headers {
			if vv, ok := v.(string); ok {
				r.req.Header.Set(k, vv)
				continue
			}
			if vv, ok := v.([]string); ok {
				for _, vvv := range vv {
					r.req.Header.Add(k, vvv)
				}
				continue
			}
			r.req.Header.Set(k, fmt.Sprintf("%v", v))
		}
	}
}

func (r *Request) parseBody() {
	// application/x-www-form-urlencoded
	if r.options.FormParams != nil {
		if _, ok := r.options.Headers["Content-Type"]; !ok {
			r.options.Headers["Content-Type"] = CONTENT_TYPE_FORM
		}

		values := url.Values{}
		for k, v := range r.options.FormParams {
			if vv, ok := v.(string); ok {
				values.Set(k, vv)
			}
			if vv, ok := v.([]string); ok {
				for _, vvv := range vv {
					values.Add(k, vvv)
				}
			}
		}
		r.body = strings.NewReader(values.Encode())

		return
	}

	// application/json
	if r.options.JSON != nil {
		if _, ok := r.options.Headers["Content-Type"]; !ok {
			r.options.Headers["Content-Type"] = CONTENT_TYPE_JSON
		}

		b, err := json.Marshal(r.options.JSON)
		if err == nil {
			r.body = bytes.NewReader(b)

			return
		}
	}

	// application/xml
	if r.options.XML != nil {
		if _, ok := r.options.Headers["Content-Type"]; !ok {
			r.options.Headers["Content-Type"] = CONTENT_TYPE_XML
		}

		switch r.options.XML.(type) {
		case map[string]string:
			// 请求参数转换成xml结构
			b, err := Map2XML(r.options.XML.(map[string]string))
			if err == nil {
				r.body = bytes.NewBuffer(b)
				return
			}
		default:
			b, err := xml.Marshal(r.options.JSON)
			if err == nil {
				r.body = bytes.NewBuffer(b)
			}
		}
	}

	return
}
