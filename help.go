package gcurl

import (
	"fmt"
	"strings"
)

// 字符串分割转切片
func AllowFileExtStr2Slice(allowExt string) []string {
	temp := strings.Split(allowExt, "|")
	exts := make([]string, len(temp))
	i := 0
	for _, item := range temp {
		if item != "" {
			exts[i] = item
			i++
		}
	}
	return exts
}

// 判断是否允许的文件类型,不区分大小写，ext=扩展名，allowExt支持的扩展名如png|jpg|jpeg|gif|txt|doc|docx|pdf|mp4
func IsAllowFileExt(ext string, allowExt string) bool {
	if strings.HasPrefix(ext, ".") {
		ext = string(ext[1:])
	}
	exts := AllowFileExtStr2Slice(allowExt)
	for _, item := range exts {
		if item == "*" {
			return true
		}
		if strings.EqualFold(item, ext) {
			return true
		}
	}
	return false
}

// 拼接阿里云oss和腾讯云cos的图片地址
func PingImgUrl(endpoint, bucketKey string) string {
	return fmt.Sprintf("%s/%s", strings.TrimRight(endpoint, "/"), bucketKey)
}
