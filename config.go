package main

var suffixes = []string{
	".wav",
	".mp3",
	".pdf",
	".svg",
	".ico",
	".php",
	".png",
	".woff",
	"woff2",
	".ttf",
	".js",
	".map",
	".jpeg",
	".jpg",
	".gif",
	".css",
	".csv",
	".html"} //检测后缀
var allowedRespHeaders = []string{
	"image/png",
	"text/html",
	"application/pdf",
	"text/css",
	"audio/mpeg",
	"audio/wav",
	"video/mp4",
	"application/grpc",
} // 检测响应头

// 检测请求头
var allowedReqHeaders = []string{
	"xxx",
}
var allowedHosts = []string{
	"baidu.com",
}
var allowedPaths = []string{
	"/api/xxxx/xxxx/xxxxxx",

	// "api",
}
