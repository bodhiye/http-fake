package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct {
	DataId string `json:"dataId"` // 数据Id，需要保证在一次请求中所有的Id不重复。
	Url    string `json:"url"`    // 待检测图像的URL
}

type fakeReq struct {
	BizType string   `json:"bizType"` // 标识业务场景
	Scenes  []string `json:"scenes"`  // 指定图片检测的应用场景
	Tasks   []Task   `json:"tasks"`   // 指定检测对象
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
	fmt.Printf("url=%s\n", req.Tasks)

	fmt.Fprintf(w, `{"code":200,"data":[{"code":200,"dataId":"3fc2ac30-5ed3-4bb3-a0d5-e3d919858a2f","extras":{},"msg":"OK","results":[{"label":"sexy","rate":99.9,"scene":"porn","suggestion":"pass"}],"taskId":"img7$NtJ82sHRH7gY9QTn7eP5-1sCSXn","url":"http://cdn.yeqiongzhou.top/pulp.jpg"}],"msg":"OK","requestId":"BECA8CF0-14A6-4E0D-AF80-B2ABEEAE23AD"}`)
}

func main() {
	http.HandleFunc("/fake", fake)
	http.ListenAndServe(":2333", nil)
}
