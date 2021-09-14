# gcurl
```
用go封装http请求库，jsonrpc请求
```

## get请求示例
```
resp, err := gcurl.Get("http://devapi.nfangbian.com/test.php?a=1&b=hi123")
if err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("%T \r\n", resp) // *gcurl.Response
    respBody,_ := resp.GetBody()
    // 获取接口响应内容
    fmt.Println(respBody.GetContents())
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

```

