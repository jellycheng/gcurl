package example

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"testing"
)

// go test -run="TestGet"
func TestGet(t *testing.T)  {
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
