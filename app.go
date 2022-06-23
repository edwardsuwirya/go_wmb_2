package main

import (
	"enigmacamp.com/gowmb/model"
	"enigmacamp.com/gowmb/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := util.CreateDataSourceName()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: util.TablePrefix(),
		}})
	if err != nil {
		panic(err)
	}
	enigmaDb, err := db.DB()
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Menu{}, &model.MenuPrice{}, &model.Discount{}, &model.Customer{}, &model.TransType{}, &model.Table{}, &model.Bill{}, &model.BillDetail{})

}
