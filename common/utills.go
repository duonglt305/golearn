package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

type ValidationErrorMessages map[string]interface{}

func GetValidationMessages(err error) ValidationErrorMessages {
	vErrors := ValidationErrorMessages{}
	errors := err.(validator.ValidationErrors)
	for _, v := range errors {
		field := strings.ToLower(v.Field())
		if v.Param() != "" {
			vErrors[field] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			vErrors[field] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}
	return vErrors
}

func Pagination(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}
		limit, _ := strconv.Atoi(c.Query("limit"))
		if limit > 100 {
			limit = 100
		} else if limit <= 0 {
			limit = 15
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}
