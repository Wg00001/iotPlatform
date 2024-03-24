package models

import "gorm.io/gorm"

type ProductBasic struct {
	gorm.Model
	Id   uint64 `gorm:"column:id" json:"id"`
	Key  string `gorm:"column:key; type:varchar(50);" json:"key"`
	Name string `gorm:"column:name;type:varchar(50);" json:"name"`
	Desc string `gorm:"column:desc;type:varchar(50);" json:"desc"` //产品描述信息
}

func (table ProductBasic) TableName() string {
	return "product_basic"
}

func ProductList(name string) *gorm.DB {
	tx := DB.Debug().Model(new(ProductBasic)).Select("identity, name, `desc`, `key`, created_at")
	if name != "" {
		tx.Where("name LIKE ?", "%"+name+"%")
	}
	return tx
}
