package main

type RequestInfo struct {
	Request struct {
		URI     string            `json:"uri"`
		Method  string            `json:"method"`
		Body    string            `json:"body"`
		Headers map[string]string `json:"headers"`
		URL     string            `json:"url"`
	} `json:"request"`
	Response struct {
		Status  int               `json:"status"`
		Body    string            `json:"body"`
		Headers map[string]string `json:"headers"`
	} `json:"response"`
}

type TestTrafficData struct {
	I_ID       int64  `json:"I_ID"`
	Path       string `json:"path"`
	Host       string `json:"host"`
	Method     string `json:"method"`
	ReqHeader  string `json:"reqheader"`
	Query      string `json:"query"`
	ReqBody    string `json:"reqbody"`
	RespBody   string `json:"respbody"`
	RespStatus string `json:"respstatus"`
	TestResult string `json:"testresult"`
	HashValue  string `json:"hashvalue"`
	RespHeader string `json:"respheader"`
	Timestamp  string `json:"timestamp"`
}
