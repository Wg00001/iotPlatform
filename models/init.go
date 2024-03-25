package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iotPlatform/common"
	"log"
)

func NewDB() *gorm.DB {
	dsn := common.DSN + "/iotPlatform?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalln("ERR: models.init.NewDB: Open mysql ERR: ", err)
		return nil
	}
	err = db.AutoMigrate(&DeviceBasic{}, &ProductBasic{}, &UserBasic{})
	if err != nil {
		log.Fatalln("ERR: models.init.NewDB: AutoMigrate ERR", err)
	}
	fmt.Println("连接数据库成功")
	return db
}
