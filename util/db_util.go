package util

import (
	"enigmacamp.com/gowmb/config"
	"fmt"
)

func CreateDataSourceName() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
}
func TablePrefix() string {
	return fmt.Sprintf("%s.", config.DbSchema)
}
func PrefixedTableName(tableName string) string {
	return fmt.Sprintf("%s%s", TablePrefix(), tableName)
}
