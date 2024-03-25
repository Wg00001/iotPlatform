package handler

import (
	"iotPlatform/admin/api/internal/logic"
	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iotPlatform/common"
)

func ProductCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewProductCreateLogic(r.Context(), svcCtx)
		resp, err := l.ProductCreate(&req)
		common.Response(r, w, resp, err)
	}
}
