package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func SendHTTPToXray(t TestTrafficData) (string, *http.Response, error) {

	proxyURL, _ := url.Parse("http://127.0.0.1:3234") // 安全扫描器代理地址
	var ProxyClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
	u := url.URL{Scheme: "http", Host: t.Host, Path: t.Path}
	if t.Query != "" {
		u.RawQuery = t.Query
	}
	// 发送 HTTP 请求
	req, err := http.NewRequest(t.Method, u.String(), strings.NewReader(t.ReqBody))
	if err != nil {
		return u.String(), nil, err
	}
	headers, _ := jsonToHeader(t.ReqHeader)
	req.Header = headers
	// 发送请求
	resp, err := ProxyClient.Do(req)
	if err != nil {
		return u.String(), nil, err
	}
	// 使用 defer 来确保响应体被关闭
	defer resp.Body.Close()
	// fmt.Println("发送成功")
	return u.String(), resp, nil
}

func jsonToHeader(jsonStr string) (http.Header, error) {
	// 解析JSON字符串
	var jsonData map[string]string
	err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&jsonData)
	if err != nil {
		return nil, err
	}
	// 创建一个空的HTTP标头
	header := http.Header{}
	header.Add("dj-rc", "test")
	// 将JSON数据添加到HTTP标头中
	for key, value := range jsonData {
		header.Add(key, value)
	}
	return header, nil
}
