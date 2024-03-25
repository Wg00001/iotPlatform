package handler

import (
	"iotPlatform/admin/api/internal/logic"
	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iotPlatform/common"
)

func ProductModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductModifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewProductModifyLogic(r.Context(), svcCtx)
		resp, err := l.ProductModify(&req)
		common.Response(r, w, resp, err)
	}
}
