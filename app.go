package main

import (
	"enigmacamp.com/gowmb/config"
	"enigmacamp.com/gowmb/repository"
	"log"
)

func main() {
	cfg := config.NewConfig()
	db := cfg.DbConn()
	//migration.DbMigration(db)
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

	//1. Daftarkan 4 table (Bulk Insert)
	//dbRepo := repository.NewTableRepository(db)
	//tables := []model.Table{
	//	{TableDescription: "T01"},
	//	{TableDescription: "T02"},
	//	{TableDescription: "T03"},
	//	{TableDescription: "T04"},
	//}
	//err := dbRepo.CreateBulk(tables)
	//if err != nil {
	//	panic(err)
	//}

	//2. Update Table T03 karena ada pemakaian
	dbRepo := repository.NewTableRepository(db)
	table, err := dbRepo.FindById(3)
	if err != nil {
		log.Println(err)
	}
	err = dbRepo.Update(&table, map[string]interface{}{"is_available": false})
	if err != nil {
		log.Println(err)
	}
}
