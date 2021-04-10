package users

import (
	"github.com/gin-gonic/gin"
	"golearn/common"
)

type LoginSerializer struct {
	c *gin.Context
}
type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func (serializer *LoginSerializer) Response() LoginResponse {
	u := serializer.c.MustGet("user_id").(uint)
	user := LoginResponse{
		AccessToken: common.GenToken(u),
	}
	return user
}
