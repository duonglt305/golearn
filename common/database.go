package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func NewConnection() *gorm.DB {
	if DB == nil {
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                      getDataSourceName(),
			DefaultStringSize:        255,
			DisableDatetimePrecision: true,
		}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			println(err.Error())
		}
		DB = db
	}
	return DB
}

func getDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
