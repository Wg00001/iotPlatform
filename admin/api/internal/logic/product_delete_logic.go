package logic

import (
	"context"
	"errors"
	"iotPlatform/models"
	"log"

	"iotPlatform/admin/api/internal/svc"
	"iotPlatform/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteRequest) (resp string, err error) {
	var cnt int64
	err = l.svcCtx.DB.Model(new(models.DeviceBasic)).
		Where("pid = ?", req.Id).
		Count(&cnt).Error
	if err != nil {
		log.Println("ERR: admin.api.logic.product_delete_logic: ERR1: ", err)
		return
	}
	if cnt > 0 {
		err = errors.New("设备已试炼，无法删除")
		return
	}
	err = l.svcCtx.DB.Debug().
		Where("id = ?", req.Id).
		Delete(new(models.ProductBasic)).
		Error
	if err != nil {
		log.Println("ERR: admin.api.logic.product_delete_logic: ERR2: ", err)
	}
	return
}
