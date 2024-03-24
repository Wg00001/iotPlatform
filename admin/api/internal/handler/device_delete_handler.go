package handler

import (
	"iotPlatform/admin/api/internal/logic"
	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"iotPlatform/common"
)

func DeviceDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewDeviceDeleteLogic(r.Context(), svcCtx)
		err := l.DeviceDelete(&req)
		common.Response(r, w, nil, err)
	}
}
