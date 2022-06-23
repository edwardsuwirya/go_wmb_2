package model

import (
	"enigmacamp.com/gowmb/config"
	"fmt"
)

type MenuPrice struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	MenuID uint
	Price  float32
}

func (MenuPrice) TableName() string {
	return config.PrefixedTableName("m_menu_price")
}
func (m MenuPrice) String() string {
	return fmt.Sprintf("Id: %d, Menu.ID: %d, Price: %f", m.ID, m.MenuID, m.Price)
}
