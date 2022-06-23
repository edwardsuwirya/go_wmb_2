package model

import (
	"enigmacamp.com/gowmb/util"
	"fmt"
)

type BillDetail struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	BillID      uint
	MenuPriceID uint
	MenuPrice   MenuPrice
	Qty         float32
}

func (BillDetail) TableName() string {
	return util.PrefixedTableName("t_bill_detail")
}
func (b BillDetail) String() string {
	return fmt.Sprintf("Id: %d, Bill.ID: %d, MenuPrice.ID: %d, MenuPrice: %v, Qty: %f", b.ID, b.BillID, b.MenuPriceID, b.MenuPrice, b.Qty)
}
