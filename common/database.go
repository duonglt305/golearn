package common

import (
	"fmt"
	"golearn/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		if config.Config.Debug {
			DB = DB.Debug()
		}
	}
	return DB
}

func getDataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.Name,
	)
}
