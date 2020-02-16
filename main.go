package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type fakeReq struct {
	Data struct {
		URI string `json:"uri"`
	} `json:"data"`
	Params struct {
		Scenes []string `json:"scenes"`
	} `json:"params"`
}

// 处理application/json类型的POST请求
func fake(w http.ResponseWriter, r *http.Request) {
	// 根据请求body创建一个json解析器实例
	decoder := json.NewDecoder(r.Body)
	// 用于存放参数数据
	var req fakeReq
	// 解析参数 存入map
	decoder.Decode(&req)
	// 打印日志
	fmt.Printf("url=%s\n", req.Data.URI)

	fmt.Fprintf(w, `{"code": 200, "data": [{"code": 200, "msg": "OK", "results": [{"label": "qrcode", "rate": 99.91, "scene": "ad", "suggestion": "block"}], "taskId": "img1gkLHlcQAM44mg2JCKEcz@-1s3ymX"}], "msg": "OK", "requestId": "CF782CC1-CC21-4333-8D5C-57EDD7BBD917"}`)
}

func main() {
	http.HandleFunc("/fake", fake)
	http.ListenAndServe(":2333", nil)
}
