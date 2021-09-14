package gcurl

import "time"

type Options struct {
	Debug      bool
	Timeout    time.Duration
	Query      interface{}
	Headers    map[string]interface{}
	Cookies    interface{}
	FormParams map[string]interface{}
	JSON       interface{}
	XML        interface{}
	Proxy      string
}

func NewOptions() Options {
	return Options{Debug: false, Timeout: 10 * time.Second,}
}
