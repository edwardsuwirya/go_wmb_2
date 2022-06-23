package model

import (
	"enigmacamp.com/gowmb/util"
	"fmt"
)

type TransType struct {
	ID          string `gorm:"primaryKey"`
	Description string
}

func (TransType) TableName() string {
	return util.PrefixedTableName("m_trans_type")
}
func (c TransType) String() string {
	return fmt.Sprintf("Id: %d, Description: %s", c.ID, c.Description)
}
