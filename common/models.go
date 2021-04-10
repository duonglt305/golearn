package common

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Auth struct {
	Password string `gorm:"column:password;size:255;not null"`
}

func (auth Auth) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should be not empty")
	}
	bytes := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	auth.Password = string(hashedPassword)
	return nil
}

func (auth Auth) VerifyPassword(password string) error {
	bytes := []byte(password)
	hashedBytes := []byte(auth.Password)
	return bcrypt.CompareHashAndPassword(hashedBytes, bytes)
}

type Model struct {
	CreatedAt time.Time `gorm:"column:created_at;sql:DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"column:updated_at;sql:DEFAULT:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
