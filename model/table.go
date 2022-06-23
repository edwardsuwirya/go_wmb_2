package model

import (
	"enigmacamp.com/gowmb/config"
	"fmt"
)

type Table struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	TableDescription string
	Bills            []Bill
}

func (Table) TableName() string {
	return config.PrefixedTableName("m_table")
}
func (t Table) String() string {
	return fmt.Sprintf("Id: %d, Desc: %s, Bills: %v", t.ID, t.TableDescription, t.Bills)
}
