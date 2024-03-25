package logic

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"iotPlatform/api"
	"iotPlatform/helper"
	"iotPlatform/models"
	"log"

	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceCreateLogic {
	return &DeviceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceCreateLogic) DeviceCreate(req *types.DeviceCreateRequest) (resp string, err error) {
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//1.数据库中新增设备
		deviceBasic := &models.DeviceBasic{
			Id:     req.Pid,
			Name:   req.Name,
			Key:    uuid.New().String(),
			Secret: uuid.New().String(),
		}
		err := tx.Create(deviceBasic).Error
		if err != nil {
			log.Println("ERR: admin.api.logic.device_create_logic: ERR1: ", err)
			return err
		}
		// 2. EMQX中新增认证设备
		err = api.CreateAuthUser(&api.CreateAuthUserRequest{
			UserId:   deviceBasic.Key,
			Password: helper.Md5(deviceBasic.Key + deviceBasic.Secret),
		})
		if err != nil {
			logx.Error("ERR: admin.api.logic.device_create_logic: ERR2: ", err)
			return err
		}
		return nil
	})
	return
}
