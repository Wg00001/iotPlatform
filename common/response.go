package common

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type ResponseBody struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func Response(r *http.Request, w http.ResponseWriter, res any, err error) {
	if err != nil {
		//可以根据不同错误码返回不同的错误信息
		body := ResponseBody{
			Code: 1,
			Data: nil,
			Msg:  "ERR",
		}
		httpx.WriteJson(w, http.StatusOK, body)
		return
	}
	body := ResponseBody{
		Code: 0,
		Data: res,
		Msg:  "successful",
	}
	httpx.WriteJson(w, http.StatusOK, body)
}
