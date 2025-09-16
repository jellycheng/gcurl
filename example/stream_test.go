package example

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"os"
	"testing"
)

// go test -run="TestStream"
func TestStream(t *testing.T) {
	urlStr := "http://127.0.0.1:18688/hd/drp?app=1128&haku=1"
	binFile := "./xxx.bin"
	byteCon, _ := os.ReadFile(binFile) // 读取文件内容，返回[]byte
	resCont, err := gcurl.SendPostData4StreamV1(urlStr, byteCon)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(resCont))
	}
}
