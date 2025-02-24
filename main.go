package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/segmentio/kafka-go"
)

// github.com/segmentio/kafka-go
var (
	reader *kafka.Reader
	topic  = "secapisixtest"
)

func readKafka(ctx context.Context) {
	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"1.1.1.1:9092", "2.2.2.2:9092", "3.3.3.3:9092"}, // kafka 地址
		Topic:          topic,
		CommitInterval: 1 * time.Second,
		GroupID:        "asdfg11111",
		StartOffset:    kafka.FirstOffset,
	})
	// offset, err := reader.SetOffset()
	for {
		if isWithinWorkingHours() { //只有不在09:00 - 19:00之间才消费
			if message, err := reader.ReadMessage(ctx); err != nil {
				fmt.Println("读取失败")
				break
			} else {
				// fmt.Println("a")
				parseJsonData(string(message.Value))
			}
		} else {
			log.Println("sleep")
			time.Sleep(3600 * time.Second)
		}
	}
}

// 输出Headers字段为JSON
func headersAsJSON(headers map[string]string) (string, error) {
	headersJSON, err := json.Marshal(headers)
	if err != nil {
		return "", err
	}
	return string(headersJSON), nil
}

func parseJsonData(jsonData string) {
	var requestInfoSlice RequestInfo
	// 使用json.Unmarshal解析JSON数据到RequestInfo结构体
	err := json.Unmarshal([]byte(jsonData), &requestInfoSlice)
	if err != nil {
		fmt.Println("解析JSON时发生错误:", err)
		return
	}
	// 访问解析后的数据
	// for _, requestInfo := range requestInfoSlice {
	reqHeadersJSON, err := headersAsJSON(requestInfoSlice.Request.Headers)
	if err != nil {
		fmt.Println("无法转换Headers为JSON:", err)
		return
	}
	respHeadersJSON, err := headersAsJSON(requestInfoSlice.Response.Headers)
	if err != nil {
		fmt.Println("无法转换Headers为JSON:", err)
		return
	}
	parsedURL, err := url.Parse(requestInfoSlice.Request.URL)
	if err != nil {
		fmt.Println("URL解析出错:", err)
		return
	}
	host := parsedURL.Host
	host = removePortFromHost(host)
	path := parsedURL.Path
	query := parsedURL.RawQuery
	h := md5.New()
	h.Write([]byte(host + path + requestInfoSlice.Request.Method))
	hashValue := hex.EncodeToString(h.Sum(nil))
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	t := TestTrafficData{
		Timestamp:  timestamp,
		HashValue:  hashValue,
		Host:       host,
		Path:       path,
		Query:      query,
		ReqHeader:  reqHeadersJSON,
		Method:     requestInfoSlice.Request.Method,
		ReqBody:    requestInfoSlice.Request.Body,
		RespStatus: fmt.Sprint(requestInfoSlice.Response.Status),
		RespBody:   requestInfoSlice.Response.Body,
		RespHeader: respHeadersJSON,
	}
	if !queryRowBool(t.Path, t.RespHeader, t.RespStatus, t.ReqHeader, t.Host) {
		// fmt.Println(t.Host)
		SendHTTPToXray(t)
	}
	// }
}

func queryRowBool(path, respHeadersJSON, status, reqheader, host string) bool {
	if endsWithAny(path, suffixes) {
		return true
	} else if containsString(respHeadersJSON, allowedRespHeaders) {
		return true
	} else if status != "200" {
		return true
	} else if containsString(reqheader, allowedReqHeaders) {
		return true
	} else if containsString(host, allowedHosts) {
		log.Println("host白名单host:", host)
		return true
	} else if containsString(path, allowedPaths) {
		log.Println("path白名单接口:", path)
		return true
	}
	return false
}

func main() {
	ctx := context.Background()
	readKafka(ctx)
}
