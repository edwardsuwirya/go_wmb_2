package model

import (
	"enigmacamp.com/gowmb/util"
	"fmt"
)

type Menu struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	MenuName   string
	MenuPrices []MenuPrice
}

func (Menu) TableName() string {
	return util.PrefixedTableName("m_menu")
}
func (m Menu) String() string {
	return fmt.Sprintf("Id: %d, Name: %s, MenuPrice: %v", m.ID, m.MenuName, m.MenuPrices)
}
