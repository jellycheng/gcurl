package gcurl

import (
	"io"
	"time"
)

type Options struct {
	Debug       bool
	Log         WriterLogger
	Timeout     time.Duration
	Query       interface{}
	Headers     map[string]interface{}
	Cookies     interface{}
	FormParams  map[string]interface{}
	JSON        interface{}
	XML         interface{}
	DefaultBody io.Reader
	Proxy       string
}

func NewOptions() Options {
	return Options{Debug: false, Timeout: 15 * time.Second}
}
