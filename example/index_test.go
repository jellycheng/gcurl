package example

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"testing"
)

// go test -run="TestGetLog"
func TestGetLog(t *testing.T)  {
	cli := gcurl.NewClient()
	o := cli.GetOptions()
	o.Debug = true
	o.Log = gcurl.NewDefaultLogger()
	//o.Log = nil
	cli.SetOptions(o)
	fmt.Println(o.Debug)
	resp, err := cli.Get("http://devapi.nfangbian.com/test.php?a=1&b=hi123")
	if err != nil {
		fmt.Println(err)
	} else {
		respBody,_ := resp.GetBody()
		fmt.Println("接口响应结果： ", respBody.GetContents())
	}
	cli.Logf("请求完毕")
}

// go test -run="TestGet2"
func TestGet2(t *testing.T)  {
	cli := gcurl.NewClient()
	resp, err := cli.Get("http://devapi.nfangbian.com/test.php?a=1&b=hi123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T \r\n", resp) // *gcurl.Response
		respBody,_ := resp.GetBody()
		fmt.Println(respBody.GetContents())
	}

	resp2, err2 := gcurl.Get("http://devapi.nfangbian.com/test.php?a=1&b=hi123")
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%T \r\n", resp2) // *gcurl.Response
		respBody,_ := resp.GetBody()
		fmt.Println(respBody.GetContents())
	}

	if resp, err := cli.Get("http://devapi.nfangbian.com/test.php?a=100&b=b200", gcurl.Options{
							Query: map[string]interface{}{
								"user": "123",
								"tags[]": []string{"学习力", "tagN"},
								"nickname": "大大",
							},
						}); err == nil {
		fmt.Printf("%s \r\n", resp.GetRequest().URL.RawQuery) // a=100&b=b200&nickname=%E5%A4%A7%E5%A4%A7&tags%5B%5D=%E5%AD%A6%E4%B9%A0%E5%8A%9B&tags%5B%5D=tagN&user=123

		respBody,_ := resp.GetBody()
		fmt.Println(respBody.GetContents())

	} else {
		fmt.Println(err)
	}

	if resp, err := cli.Get("http://devapi.nfangbian.com/test.php?a=100&b=b200", gcurl.Options{
		Query: map[string]interface{}{
			"user": "123",
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "大大",
		},
		Headers: map[string]interface{}{
			"User-Agent": "gcurl/1.0",
			"Accept":     "application/json",
			"X-USERID":   123456,
			"X-Tag":      []string{"go", "php"},
		},
	}); err == nil {
		fmt.Printf("%s \r\n", resp.GetRequest().URL.RawQuery)

		respBody,_ := resp.GetBody()
		fmt.Println(respBody.GetContents())

	} else {
		fmt.Println(err)
	}

}

// go test -run="TestMap2XML"
func TestMap2XML(t *testing.T)  {
	mapData := map[string]string{
			"hello":"world",
			"123":"yes",
			"AaBbC": "789",
	}
	// <xml><hello>world</hello><123>yes</123><AaBbC>789</AaBbC></xml>
	con, _ := gcurl.Map2XML(mapData, "xml")
	fmt.Println(string(con))

}

// go test -run="TestPost"
func TestPost(t *testing.T)  {
	resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
			"User-Agent":    "gcurl/1.0",
			"Authorization": "Bearer access_token1234",
		},
		Query: map[string]interface{}{
			"user": 123,
			"tags[]": []string{"学习力", "tagN"},
			"nickname": "大大",
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
		fmt.Println(body)
	}

}

// go test -run="TestPostjson"
func TestPostjson(t *testing.T)  {
	resp, err := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
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
		fmt.Println(body)
	}

	strJson := `{"age":26,"name":"admin123"}`
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
		fmt.Println(body)
	}

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
		}{"val1", []string{"val2-1", "val2-2"}, 333,true},
	}); err != nil {
		fmt.Println(err)
	} else {
		body, _ := resp.GetBody()
		fmt.Println(body)
	}

}

// go test -run=TestPostxml
func TestPostxml(t *testing.T)  {
	xmlStr := `
<xml>
<ToUserName><![CDATA[ww0ca05641227fc2e0]]></ToUserName>
<FromUserName><![CDATA[sys]]></FromUserName>
<CreateTime>1652515415</CreateTime>
<MsgType><![CDATA[event]]></MsgType>
<AgentID>1000019</AgentID>
<Event><![CDATA[change_app_admin]]></Event></xml>
`
	resp1, err1 := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
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
		XML: xmlStr,
	})
	if err1 != nil {
		fmt.Println(err1)
	} else {
		body, _ := resp1.GetBody()
		fmt.Println(body)
	}

	fmt.Println("==================")

	resp2, err2 := gcurl.Post("http://devapi.nfangbian.com/test.php?a=2&b=say123", gcurl.Options{
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
		XML: map[string]string{
			"name":      "admin",
			"age":       "24",
			"interests": "篮球,旅游,听音乐",
		},
	})
	if err2 != nil {
		fmt.Println(err2)
	} else {
		body, _ := resp2.GetBody()
		fmt.Println(body)
	}

}
