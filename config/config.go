package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	DbHost     = "167.172.69.254"
	DbPort     = "5432"
	DbUser     = "smm2"
	DbPassword = "5mmB2"
	DbName     = "smm2"
	DbSchema   = "restowmb"
)

type Config struct {
	db *gorm.DB
}

func (c *Config) initDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DbHost, DbUser, DbPassword, DbName, DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: TablePrefix(),
		}})
	if err != nil {
		panic(err)
	}
	enigmaDb, err := db.DB()
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	}
	c.db = db
}

func (c *Config) DbConn() *gorm.DB {
	return c.db
}
func NewConfig() Config {
	cfg := Config{}
	cfg.initDb()
	return cfg
}
func TablePrefix() string {
	return fmt.Sprintf("%s.", DbSchema)
}
func PrefixedTableName(tableName string) string {
	return fmt.Sprintf("%s%s", TablePrefix(), tableName)
}
