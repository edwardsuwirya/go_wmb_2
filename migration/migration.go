package migration

import (
	"enigmacamp.com/gowmb/model"
	"gorm.io/gorm"
)

func DbMigration(db *gorm.DB) {
	db.AutoMigrate(&model.Menu{}, &model.MenuPrice{}, &model.Discount{}, &model.Customer{}, &model.TransType{}, &model.Table{}, &model.Bill{}, &model.BillDetail{})

}
