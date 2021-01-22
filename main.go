package main

import (
	"fmt"
	"net/http"
	"time"
)

// 处理application/json类型的POST请求
func fake(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_ = query.Get("url")

	normalRes := `{"code":0,"msg":"请求成功","data":{"tags":[{"confidence":55,"value":"动物"},{"confidence":32,"value":"食物"},{"confidence":8,"value":"人物"}]}}`

	pornRes := `{"code":0,"msg":"请求成功","data":{"tags":[{"confidence":55,"value":"植物"},{"confidence":32,"value":"零食"},{"confidence":8,"value":"七牛云"}]}}`

	var (
		str  string
		time = time.Now().Unix()
	)
	if time%10 < 2 {
		str = pornRes
	} else {
		str = normalRes
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(str)))
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/handler", fake)
	http.ListenAndServe(":2048", nil)
}
