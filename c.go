package gcurl

import "time"

const (
	ContentTypeJson      = "application/json"
	ContentTypeForm      = "application/x-www-form-urlencoded"
	ContentTypeOctet     = "application/octet-stream"
	ContentTypeMultipart = "multipart/form-data"
	ContentTypeXml       = "application/xml"
	ContentTypeTexthtml  = "text/html"
	ContentTypeTextxml   = "text/xml"
)

const (
	TraceIdHeader = "X-Trace-Id"
	ELLIPSES      = "..."
)

const (
	// DefaultTimeout 默认超时值
	DefaultTimeout = 15 * time.Second
)
