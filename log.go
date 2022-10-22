package gcurl

import "fmt"

// WriterLogger 日志接口
type WriterLogger interface {
	Printf(string, ...interface{})
}

type DefaultLogger struct{}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

func (l DefaultLogger) Printf(format string, values ...interface{}) {
	val := fmt.Sprintf(format, values...)
	fmt.Println(val)
}
