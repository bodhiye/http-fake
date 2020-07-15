package main

import (
	"fmt"
	"net/http"
	"strings"
)

// 处理application/json类型的POST请求
func fake(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query.Get("url")

	normalRes := `{"code":200,"data":[{"code":200,"msg":"OK","results":[{"label":"normal","rate":99.91,"scene":"normal","suggestion":"pass","details":[{"startTime":0,"endTime":30,"text":"啊。","label":"normal"}]}],"taskId":"vc_f_5M$CW$2MQNI5nL1aAkWzfR-1sR4@n"}],"msg":"OK","requestId":"3612842B-0F3A-4AD4-A6B5-7F7AD4BE1956"}`

	pornRes := `{"code":200,"data":[{"code":200,"msg":"OK","results":[{"label":"porn","rate":99.91,"scene":"antispam","suggestion":"block","details":[{"startTime":0,"endTime":2,"text":"嗯。","label":"normal"},{"startTime":2,"endTime":2,"text":"啊。","label":"normal"},{"startTime":4,"endTime":5,"text":"哎。","label":"normal"},{"startTime":7,"endTime":10,"text":"哎。","label":"normal"},{"startTime":10,"endTime":12,"text":"哎。","label":"normal"},{"startTime":12,"endTime":12,"text":"是。","label":"normal"},{"startTime":13,"endTime":14,"text":"死。","label":"normal"},{"startTime":16,"endTime":17,"text":"啊。","label":"normal"},{"startTime":17,"endTime":18,"text":"啊。","label":"normal"},{"startTime":22,"endTime":23,"text":"是。","label":"normal"},{"startTime":26,"endTime":27,"text":"哎。","label":"normal"},{"startTime":28,"endTime":28,"text":"嗯。","label":"normal"},{"startTime":33,"endTime":34,"text":"怎么了呢？","label":"normal"},{"startTime":35,"endTime":37,"text":"你的身体。","label":"normal"},{"startTime":37,"endTime":40,"text":"开始在战斗了呢。","label":"normal"},{"startTime":40,"endTime":44,"text":"吸收不了了吗？","label":"normal"},{"startTime":45,"endTime":47,"text":"真的受不了了。","label":"normal"},{"startTime":48,"endTime":54,"text":"嗯，嗯，姐姐知道了。那。","label":"normal"},{"startTime":55,"endTime":60,"text":"姐姐还不停下来。","label":"normal"},{"startTime":60,"endTime":69,"text":"嗯。嗯。","label":"porn"},{"startTime":69,"endTime":85,"text":"去。嗯，嗯。是。啊。","label":"porn"},{"startTime":87,"endTime":88,"text":"哎。","label":"normal"},{"startTime":89,"endTime":91,"text":"哎。","label":"normal"},{"startTime":94,"endTime":95,"text":"可以可以。","label":"normal"},{"startTime":95,"endTime":103,"text":"哎。嗯。死。嗯。","label":"porn"},{"startTime":105,"endTime":107,"text":"哎。","label":"normal"},{"startTime":107,"endTime":108,"text":"是。","label":"normal"},{"startTime":109,"endTime":111,"text":"啊。","label":"normal"},{"startTime":111,"endTime":112,"text":"世界。","label":"normal"},{"startTime":112,"endTime":115,"text":"哎。","label":"normal"},{"startTime":115,"endTime":116,"text":"是。","label":"normal"},{"startTime":116,"endTime":118,"text":"嗯。","label":"normal"},{"startTime":120,"endTime":121,"text":"嗯。","label":"normal"},{"startTime":125,"endTime":127,"text":"s。","label":"normal"},{"startTime":132,"endTime":133,"text":"啊。","label":"normal"},{"startTime":133,"endTime":134,"text":"嗯。","label":"normal"},{"startTime":139,"endTime":140,"text":"啊。","label":"normal"},{"startTime":144,"endTime":145,"text":"对。","label":"normal"},{"startTime":145,"endTime":146,"text":"对。","label":"normal"},{"startTime":148,"endTime":148,"text":"啊。","label":"normal"},{"startTime":149,"endTime":151,"text":"哎。","label":"normal"},{"startTime":158,"endTime":159,"text":"嗯。","label":"normal"},{"startTime":160,"endTime":161,"text":"去。","label":"normal"},{"startTime":161,"endTime":162,"text":"嗯。","label":"normal"},{"startTime":164,"endTime":166,"text":"操。","label":"normal"},{"startTime":173,"endTime":174,"text":"是。","label":"normal"},{"startTime":180,"endTime":181,"text":"看。","label":"normal"},{"startTime":181,"endTime":182,"text":"是。","label":"normal"},{"startTime":183,"endTime":184,"text":"是。","label":"normal"},{"startTime":184,"endTime":185,"text":"是。","label":"normal"},{"startTime":186,"endTime":187,"text":"去。","label":"normal"},{"startTime":189,"endTime":190,"text":"是。","label":"normal"},{"startTime":191,"endTime":192,"text":"7。","label":"normal"},{"startTime":196,"endTime":197,"text":"嗯。","label":"normal"},{"startTime":199,"endTime":209,"text":"哎。啊。嗯。哎。哎。","label":"porn"},{"startTime":212,"endTime":213,"text":"满足了。","label":"normal"},{"startTime":213,"endTime":215,"text":"嗯。","label":"normal"},{"startTime":215,"endTime":216,"text":"嗯。","label":"normal"},{"startTime":217,"endTime":226,"text":"对不起啊，姐姐把你的耳朵弄的这么多口水。","label":"normal"},{"startTime":226,"endTime":227,"text":"哎，熊二。","label":"normal"},{"startTime":227,"endTime":229,"text":"嗯。","label":"normal"},{"startTime":229,"endTime":237,"text":"用姐姐的小毛巾帮你把耳朵擦干净吧。","label":"normal"},{"startTime":237,"endTime":238,"text":"哎，哎。","label":"normal"},{"startTime":239,"endTime":246,"text":"再调小么经验，感谢接qq姐姐。身体很多地方。","label":"normal"},{"startTime":246,"endTime":248,"text":"嗯。","label":"normal"},{"startTime":249,"endTime":253,"text":"不过你应该不会在意这个吧。","label":"normal"},{"startTime":253,"endTime":256,"text":"毕竟解决现在。","label":"normal"},{"startTime":257,"endTime":260,"text":"可是压在你的身体上。你。","label":"normal"},{"startTime":260,"endTime":262,"text":"55。","label":"normal"},{"startTime":263,"endTime":269,"text":"真是的，才不会让你得意忘形。哎，哎。","label":"normal"},{"startTime":269,"endTime":270,"text":"嗯。","label":"normal"},{"startTime":271,"endTime":287,"text":"我就知道你的耳朵是弱点，可不可以告诉姐姐，我也喜欢你的耳朵这么可爱又这么好玩呢。","label":"normal"},{"startTime":288,"endTime":300,"text":"为什么啊，我也什么呢，怎么样再来啊看谁比较厉害啊。","label":"normal"},{"startTime":300,"endTime":303,"text":"姐姐也开始累了。","label":"normal"}]}],"taskId":"vc_f_6XTF3O3VHxJ4JXHWhKIHMm-1sR4Yh"}],"msg":"OK","requestId":"A04E052C-3619-4825-9EA6-A3FB837B505B"}`

	var str string
	if strings.Contains(url, "300000-330000.aac") {
		str = pornRes
	} else {
		str = normalRes
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", fmt.Sprint(len(str)))
	w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/fake", fake)
	http.ListenAndServe(":2333", nil)
}
