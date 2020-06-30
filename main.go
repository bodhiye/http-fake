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

	str := `{"code":0,"message":"","result":{"detections":[{"class":"guns","pts":[[234,127],[730,127],[730,313],[234,313]],"score":0.9959627687931061},{"class":"guns","pts":[[711,548],[1093,548],[1093,779],[711,779]],"score":0.9943256378173828}]}}`

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(str)))
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/fake", fake)
	http.ListenAndServe(":2333", nil)
}
