package users

import (
	"github.com/gin-gonic/gin"
	"golearn/common"
)

type LoginValidator struct {
	Email    string `form:"email" json:"email" binding:"email"`
	Password string `form:"password" json:"password" binding:"min=8,max=255"`
}

func (v *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	return nil
}
func NewLoginValidator() LoginValidator {
	validator := LoginValidator{}
	return validator
}
