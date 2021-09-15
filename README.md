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

## post请求示例
```
参数优先级FormParams > JSON > XML

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

```

## post json示例
```
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

```
