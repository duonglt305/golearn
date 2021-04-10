package users

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golearn/common"
	"time"
)

type User struct {
	ID             uint      `gorm:"column:id;primaryKey;not null"`
	Email          string    `gorm:"column:email;uniqueIndex;size:100;not null"`
	Name           string    `gorm:"column:name;not null"`
	Password       string    `gorm:"column:password;size:255;not null"`
	Photo          string    `gorm:"column:photo"`
	LatestLoggedAt time.Time `gorm:"column:latest_logged_at"`
	common.Model
}

func (u *User) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should be not empty")
	}
	bytes := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) VerifyPassword(password string) error {
	bytes := []byte(password)
	hashedBytes := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(hashedBytes, bytes)
}
func (u *User) SetLoggedTime() {
	db := common.GetDB()
	logged := User{LatestLoggedAt: time.Now()}
	db.Where(&User{ID: u.ID}).Updates(logged)
}
func (u *User) FindOne(condition interface{}) (*User, error) {
	db := common.GetDB()
	err := db.Where(condition).First(&u).Error
	return u, err
}

func Migrate() {
	db := common.GetDB()
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
