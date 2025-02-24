package main

import (
	"strings"
	"time"
)

// 入库白名单
func endsWithAny(s string, suffix []string) bool {
	for _, str := range suffix {
		if str != "" && strings.HasSuffix(s, str) {
			return true
		}
	}
	return false
}

// 入库白名单
func containsString(target string, slice []string) bool {
	for _, s := range slice {
		if strings.Contains(strings.ToLower(target), strings.ToLower(s)) {
			// log.Println(target)
			return true
		}
	}

	return false
}

// host去掉端口
func removePortFromHost(host string) string {
	// 使用 strings.Split 函数来根据冒号分割字符串
	parts := strings.Split(host, ":")

	// 如果字符串中包含冒号，则返回分割后的第一部分，否则返回原始字符串
	if len(parts) > 1 {
		return parts[0]
	}
	return host
}

func isWithinWorkingHours() bool {
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 0, 0, 0, now.Location())
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 24, 0, 0, 0, now.Location())
	return now.After(startTime) && now.Before(endTime)
}
