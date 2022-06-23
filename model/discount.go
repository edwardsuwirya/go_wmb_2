package model

import "fmt"

type Discount struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Description string
	Pct         uint
}

func (Discount) TableName() string {
	return "restowmb.m_discount"
}
func (d Discount) String() string {
	return fmt.Sprintf("Id: %d, Description: %s, Pct: %d", d.ID, d.Description, d.Pct)
}
