package handler

import (
	"iotPlatform/open/api/internal/logic"
	"iotPlatform/open/api/internal/svc"
	"iotPlatform/open/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iotPlatform/common"
)

func SendMessageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewSendMessageLogic(r.Context(), svcCtx)
		resp, err := l.SendMessage(&req)
		common.Response(r, w, resp, err)
	}
}
