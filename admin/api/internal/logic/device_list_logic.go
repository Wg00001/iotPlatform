package logic

import (
	"context"
	"iotPlatform/helper"
	"iotPlatform/models"
	"log"

	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceListLogic {
	return &DeviceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceListLogic) DeviceList(req *types.DeviceListRequest) (resp *types.DeviceListReply, err error) {
	req.Size = helper.If(req.Size == 0, 20, req.Size).(int)
	req.Page = helper.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)
	var count int64
	resp = new(types.DeviceListReply)
	data := make([]*types.DeviceListBaisc, 0)
	err = models.GetDeviceList(l.svcCtx.DB, req.Name).
		Count(&count).
		Limit(req.Size).
		Offset(req.Page).
		Find(&data).
		Error
	if err != nil {
		log.Println("ERR: admin.api.logic.device_list_logic: ERR1: ", err)
		return
	}
	resp.Count = count
	resp.List = data
	return
}
