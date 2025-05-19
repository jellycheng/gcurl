# gcurl
```
封装http请求，支持get、post、put、delete等请求方式，支持上传文件
封装jsonrpc协议发起post请求
封装sse协议

```

## get请求示例1
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {
	resp, err := gcurl.Get("http://devapi.nfangbian.com/test.php?a=1&b=hi123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T \r\n", resp) // *gcurl.Response
		respBody,_ := resp.GetBody()
		// 获取接口响应内容
		fmt.Println(respBody.GetContents())
	}
}

```

## get请求示例2
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {
	if resp, err := gcurl.Get("http://devapi.nfangbian.com/test.php?a=100&b=b200", gcurl.Options{
		// 追加和覆盖get参数
		Query: map[string]interface{}{
			"user": "123",
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "小小",
			"b":"bxxx",
		},
		Headers: map[string]interface{}{
			"User-Agent": "gcurl/1.0",
			"Accept":     gcurl.ContentTypeJson,
			"X-USERID":   123456,
			"X-Tag":      []string{"go", "php", "java"},
			gcurl.TraceIdHeader: "traceid-abc-123-xyz",
		},
	}); err == nil {
		fmt.Printf("请求参数：%s \r\n", resp.GetRequest().URL.RawQuery)
		respBody,_ := resp.GetBody()
		fmt.Println("响应结果：", respBody.GetContents())

	} else {
		fmt.Println(err)
	}

}
```

## post请求示例1
```
参数优先级FormParams > JSON > XML

package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {

	resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
		Headers: map[string]interface{}{
			//"Content-Type": "application/x-www-form-urlencoded",
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":    "gcurl/1.0",
			"Authorization": "Bearer access_token1234",
			gcurl.TraceIdHeader: "trace-id-123x",
		},
		Query: map[string]interface{}{
			"user": 123,
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "大大",
			"a": 99,
			"isok":false,
		},
		FormParams: map[string]interface{}{
			"name":      "admin",
			"age":       24,
			"interests[]": []string{"篮球", "旅游", "听音乐"},
			"isAdmin":   true,
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := resp.GetBody()
		fmt.Println("响应结果：", body)
	}

}


```

## post json示例1
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {
	resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php", gcurl.Options{
		Query: map[string]interface{}{
			"user": 123,
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "大大",
			"a": 108,
			"isok":false,
		},
		JSON: map[string]interface{}{
			"name":      "admin",
			"age":       24,
			"interests": []string{"篮球", "旅游", "听音乐"},
			"isAdmin":   true,
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := resp.GetBody()
		fmt.Println("响应内容", body)
	}

}
```

## post json示例2
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {
	resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
		Headers: map[string]interface{}{
			"User-Agent":    "gcurl/1.0",
			"Authorization": "Bearer access_token1234",
			gcurl.TraceIdHeader: "me-trace-id123",
		},
		Query: map[string]interface{}{
			"user": 123,
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "大大",
			"a": 108,
			"isok":false,
		},
		JSON: map[string]interface{}{
			"name":      "admin",
			"age":       24,
			"interests": []string{"篮球", "旅游", "听音乐"},
			"isAdmin":   true,
		},
	})
	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := resp.GetBody()
		fmt.Println("响应内容", body)
	}

}
```

## post json示例3
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {
	// json字符串
	strJson := `{"age":26,"name":"账号admin123"}`
	if resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
		Headers: map[string]interface{}{
			"User-Agent":    "gcurl/1.0",
			"Authorization": "Bearer access_token1234",
		},
		Query: map[string]interface{}{
			"user": 123,
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "大大",
			"isok":false,
		},
		JSON: strJson,
	}); err != nil {
		fmt.Println(err)
	} else {
		body, _ := resp.GetBody()
		fmt.Println("响应结果：", body)
	}

}
```

## post json示例4
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {

	if resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
		Headers: map[string]interface{}{
			"User-Agent":    "gcurl/1.0",
			"Authorization": "Bearer access_token1234",
		},
		Query: map[string]interface{}{
			"user": 123,
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "大大123",
			"isok":false,
		},
		JSON: struct {
			Key1 string   `json:"key1"`
			Key2 []string `json:"key2"`
			Key3 int      `json:"key3"`
			Key4 bool      `json:"key4"`
		}{"val1，结构体方式作为参数", []string{"val2-1", "val2-2"}, 333,true},
	}); err != nil {
		fmt.Println(err)
	} else {
		body, _ := resp.GetBody()
		fmt.Println("响应结果：", body)
	}

}

```

## jsonrpc请求示例1
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
)

func main() {
	jsonrpcReqDto := gcurl.NewJsonRpcReqDto()
	jsonrpcReqDto.Method = `user\info`
	jsonrpcReqDto.Params = "hello"
	if resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
		Headers: map[string]interface{}{
			"User-Agent":    "gcurl/1.0/jsonrpc2.0",
			gcurl.TraceIdHeader: "traceid-123",
		},
		JSON: jsonrpcReqDto,
	}); err != nil {
		fmt.Println(err)
	} else {
		body, _ := resp.GetBody()
		fmt.Println("响应结果：", body)
	}

}

```

## 并发请求示例1
```
使用协程发起并发请求示例
package main

import (
	"context"
	"fmt"
	"github.com/jellycheng/gcurl"
	"sync"
)

func main() {
	wg := gcurl.NewWg()
	result := sync.Map{}
	ctx1, _ := context.WithCancel(context.Background())
	wg.RunApi(ctx1, func(ctx2 context.Context) {
		// 接口1
		resp, err := gcurl.Get("http://devapi.nfangbian.com/test.php?a=1&b=hi123")
		if err != nil {
			result.Store("api_1", err.Error())
		} else {
			respBody, _ := resp.GetBody()
			// 获取接口响应内容
			result.Store("api_1", respBody.GetContents())
		}

	})
	
	wg.RunApi(ctx1, func(ctx2 context.Context) {
		// 接口2
		resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
			Headers: map[string]interface{}{
				"Content-Type":      gcurl.ContentTypeForm,
				"User-Agent":        "gcurl/1.0",
				"Authorization":     "Bearer access_token1234",
				gcurl.TraceIdHeader: "trace-id-123x",
			},
			Query: map[string]interface{}{
				"user":     123,
				"tags[]":   []string{"学习力", "tagN"},
				"nickname": "大大",
				"a":        99,
				"isok":     false,
			},
			FormParams: map[string]interface{}{
				"name":        "admin",
				"age":         24,
				"interests[]": []string{"篮球", "旅游", "听音乐"},
				"isAdmin":     true,
			},
		})
		if err != nil {
			result.Store("api_2", err.Error())
		} else {
			respBody, _ := resp.GetBody()
			result.Store("api_2", respBody.GetContents())
		}
	})
	
	wg.Wait()
	// 统一处理api结果
	result.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}

```

## sse示例1
```
package main

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"net/http"
	"time"
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// 设置 SSE 所需的 HTTP 头
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 允许跨域

	// 检查是否支持流式响应
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	token := r.Header.Get("token")
	if token == "" {
		token = r.URL.Query().Get("token")
	}
	id := r.URL.Query().Get("id")
	content := r.URL.Query().Get("content")
	// 模拟实时数据推送
	for {

		date := time.Now().Format("2006-01-02 15:04:05")
		sseMsg := gcurl.NewSseMsg()
		sseMsg.SetData(fmt.Sprintf("收到参数token：%s,id=%s,content=%s,响应时间：%s", token, id, content, date))
		// 响应内容
		w.Write(sseMsg.FormatMsg())
		flusher.Flush()

		time.Sleep(2 * time.Second)
	}
}

func main() {
	http.HandleFunc("/sse/events", sseHandler)
	http.ListenAndServe(":8989", nil)
	// 访问地址示例：http://127.0.0.1:8989/sse/events?id=123&token=xxxx&content=您好
}

```
