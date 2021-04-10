package common

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"os"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenToken(id uint) string {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	jwtToken.Claims = jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token, _ := jwtToken.SignedString([]byte(os.Getenv("APP_KEY")))
	return token
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

type ErrorMessage struct {
	Errors map[string]interface{} `json:"errors"`
}

func NewValidationError(err error) ErrorMessage {
	resp := ErrorMessage{}
	resp.Errors = make(map[string]interface{})
	errors := err.(validator.ValidationErrors)
	for _, v := range errors {
		if v.Param() != "" {
			resp.Errors[v.Field()] = fmt.Sprintf("{%v: %v}", v.Tag(), v.Param())
		} else {
			resp.Errors[v.Field()] = fmt.Sprintf("{key: %v}", v.Tag())
		}
	}
	return resp
}

func NewError(key string, err error) ErrorMessage {
	res := ErrorMessage{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}
