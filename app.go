package main

import (
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
	//db.AutoMigrate(&model.Menu{}, &model.MenuPrice{}, &model.Discount{}, &model.Customer{}, &model.TransType{}, &model.Table{}, &model.Bill{}, &model.BillDetail{})

	//1. Daftarkan Discount 10% di master discount
	//dbRepo := repository.NewDiscountRepository(db)
	//disc := model.Discount{
	//	Description: "Discount 10%",
	//	Pct:         10,
	//}
	//err = dbRepo.Create(&disc)
	//if err != nil {
	//	panic(err)
	//}

	//2. Daftarkan Customer dengan informasi Nama Kadir, No Handphone 0877123333, dan bukan member
	//customerRepo := repository.NewCustomerRepository(db)
	//cust := model.Customer{
	//	CustomerName:  "Kadir",
	//	MobilePhoneNo: "0877123333",
	//	IsMember:      false,
	//}
	//err = customerRepo.Create(&cust)
	//if err != nil {
	//	panic(err)
	//}

	//3. Daftarkan Customer dengan informasi Nama Munaroh, No Handphone 0888901920, dan member. Karena member, berhak
	// mendapatkan discount 10%
	//discRepo := repository.NewDiscountRepository(db)
	//disc, err := discRepo.FindById(1)
	//if err != nil {
	//	panic(err)
	//}
	//customerRepo := repository.NewCustomerRepository(db)
	//cust := model.Customer{
	//	CustomerName:  "Munaroh",
	//	MobilePhoneNo: "0888901920",
	//	IsMember:      true,
	//	Discounts: []model.Discount{
	//		disc,
	//	},
	//}
	//err = customerRepo.Create(&cust)
	//if err != nil {
	//	panic(err)
	//}
	//4. Berapa Total Customer
	//customerRepo := repository.NewCustomerRepository(db)
	//count, err := customerRepo.Count()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(count)
	//5. Cari informasi member berserta discountnya dengan nama munaroh
	//customerRepo := repository.NewCustomerRepository(db)
	//customers, err := customerRepo.FindAllBy("Discounts", "customer_name ILIKE ?", "%munaroh%")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(customers)
}
