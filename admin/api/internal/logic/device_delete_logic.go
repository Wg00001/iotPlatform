package logic

import (
	"context"
	"gorm.io/gorm"
	"iotPlatform/api"
	"iotPlatform/models"
	"log"

	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteRequest) (resp string, err error) {
	device := new(models.DeviceBasic)
	err = l.svcCtx.DB.Model(new(models.DeviceBasic)).
		Select("key").
		Where("id = ?", req.Id).
		Find(device).
		Error
	if err != nil {
		log.Println("ERR: admin.api.logic.device_delete_logic: ERR1: ", err)
		return
	}
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//1.数据库中删除设备
		err = tx.Where("id = ?", req.Id).
			Delete(new(models.DeviceBasic)).
			Error
		if err != nil {
			log.Println("ERR: admin.api.logic.device_delete_logic: ERR2: ", err)
			return err
		}

		// 2. EMQX中同步删除认证设备
		err = api.DeleteAuthUser(device.Key)
		if err != nil {
			log.Println("ERR: admin.api.logic.device_delete_logic: ERR3: ", err)
			return err
		}
		return nil
	})
	return
}
