package common

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"math/rand"
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
