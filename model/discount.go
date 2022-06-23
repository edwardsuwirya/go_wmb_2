package model

import (
	"enigmacamp.com/gowmb/util"
	"fmt"
)

type Discount struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Description string
	Pct         uint
}

func (Discount) TableName() string {
	return util.PrefixedTableName("m_discount")
}
func (d Discount) String() string {
	return fmt.Sprintf("Id: %d, Description: %s, Pct: %d", d.ID, d.Description, d.Pct)
}
