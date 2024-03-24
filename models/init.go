package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iotPlatform/common"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := common.DSN + "/iotPlatform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("ERR: models.init.NewDB: Open mysql ERR: ", err)
		return
	}
	err = db.AutoMigrate(&DeviceBasic{}, &ProductBasic{}, &UserBasic{})
	if err != nil {
		log.Fatalln("ERR: models.init.NewDB: AutoMigrate ERR", err)
	}
	DB = db
}
