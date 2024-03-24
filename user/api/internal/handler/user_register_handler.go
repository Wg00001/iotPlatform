package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"iotPlatform/common"
	"iotPlatform/user/api/internal/logic"
	"iotPlatform/user/api/internal/svc"
	"iotPlatform/user/api/internal/types"
	"net/http"
)

func UserRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewUserRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UserRegister(&req)
		common.Response(r, w, resp, err)
	}
}
