package models

import (
	"gorm.io/gorm"
	"time"
)

// DeviceBasic 设备
type DeviceBasic struct {
	gorm.Model
	Id             uint64 `gorm:"column:id" json:"id"`
	ProductId      uint64 `gorm:"column:pid" json:"product_id"`
	Name           string `gorm:"column:name;type:varchar(50);" json:"name"`
	Key            string `gorm:"column:key;type:varchar(50);" json:"key"`
	Secret         string `gorm:"column:secret;type:varchar(50);" json:"secret"` //连接的密钥
	LastOnlineTime int64  `gorm:"column:last_online_time; type:int(11);" json:"last_online_time"`
}

func (table DeviceBasic) TableName() string {
	return "device_basic"
}

// GetDeviceList 获取设备列表
func GetDeviceList(db *gorm.DB, name string) *gorm.DB {
	tx := db.Model(new(DeviceBasic)).Select("device_basic.id, device_basic.name," +
		"device_basic.key, device_basic.secret, pb.name product_name, device_basic.last_online_time").
		Joins("LEFT JOIN product_basic pb ON pb.identity = device_basic.product_identity")
	if name != "" {
		tx.Where("device_basic.name LIKE ?", "%"+name+"%")
	}
	return tx
}

// UpdateDeviceOnlineTime 更新设备上线时间
//
// productKey 产品 key
// deviceKey 设备 key
func UpdateDeviceOnlineTime(db *gorm.DB, productKey, deviceKey string) error {
	var productIdentity string
	err := db.Model(new(ProductBasic)).Select("identity").Where("`key` = ?", productKey).Scan(&productIdentity).Error
	if err != nil {
		return err
	}
	err = db.Model(new(DeviceBasic)).
		Where("`key` = ? AND product_identity = ?", deviceKey, productIdentity).
		Update("last_online_time", time.Now().Unix()).Error
	return err
}
