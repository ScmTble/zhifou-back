package result

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// HttpResult response
func HttpResult(_ *http.Request, w http.ResponseWriter, data interface{}, err error) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type") //header的类型
	//w.Header().Set("content-type", "application/json")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	if err == nil {
		res := Success(data)
		httpx.WriteJson(w, http.StatusOK, res)
	} else {
		res := Fail(err.Error())
		httpx.WriteJson(w, http.StatusOK, res)
	}
}
