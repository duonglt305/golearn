package users

import (
	"fmt"
	"golearn/common"
)

type User struct {
	ID       uint   `gorm:"column:id;primaryKey;not null"`
	Email    string `gorm:"column:email;uniqueIndex;size:100;not null"`
	Name     string `gorm:"column:name;not null"`
	*common.Auth
	Photo    string `gorm:"column:photo"`
	*common.Model
}

func Migrate() {
	db := common.GetDB()
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
