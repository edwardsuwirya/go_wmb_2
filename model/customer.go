package model

import (
	"enigmacamp.com/gowmb/config"
	"fmt"
)

type Customer struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	CustomerName  string
	MobilePhoneNo string
	IsMember      bool
	Bills         []Bill
	Discounts     []Discount `gorm:"many2many:m_customer_discount;"`
}

func (Customer) TableName() string {
	return config.PrefixedTableName("m_customer")
}
func (c Customer) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, Phone: %s, Member: %v, Bills: %v, Disc: %v", c.ID, c.CustomerName, c.MobilePhoneNo, c.IsMember, c.Bills, c.Discounts)
}
