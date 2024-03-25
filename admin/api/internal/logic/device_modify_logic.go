package logic

import (
	"context"
	"iotPlatform/models"
	"log"

	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceModifyLogic {
	return &DeviceModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeviceModify 更新、更改设备
func (l *DeviceModifyLogic) DeviceModify(req *types.DeviceModifyRequest) (resp string, err error) {
	err = l.svcCtx.DB.Debug().
		Where("id = ?", req.Id).
		Updates(&models.DeviceBasic{
			ProductId: req.Pid,
			Name:      req.Name,
		}).Error
	if err != nil {
		log.Println("ERR: admin.api.logic.device_modify_logic: ERR1: ", err)
	}
	return
}
