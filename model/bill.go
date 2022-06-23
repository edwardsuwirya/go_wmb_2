package model

import (
	"enigmacamp.com/gowmb/config"
	"fmt"
	"time"
)

type Bill struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	TransDate   time.Time
	CustomerID  uint
	TableID     uint
	TransTypeID string
	TransType   TransType
	BillDetails []BillDetail
}

func (Bill) TableName() string {
	return config.PrefixedTableName("t_bill")
}
func (b Bill) String() string {
	return fmt.Sprintf("Id: %d, Trans.Date: %v, Customer.ID: %d, Table.ID: %d, TransType.ID: %s, TransType: %v", b.ID, b.TransDate, b.CustomerID, b.TableID, b.TransTypeID, b.TransType)
}
