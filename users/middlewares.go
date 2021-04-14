package users

import (
	"github.com/gin-gonic/gin"
)

func SetUserContext(c *gin.Context, id uint) {
	var u User
	if id != 0 {
		u, _ = FindOne(id)
	}
	c.Set("user_id", id)
	c.Set("user", u)
}

func GetUserContext(c *gin.Context) *User {
	u, _ := c.MustGet("user").(User)
	return &u
}
func JWTMiddleware() gin.HandlerFunc {
	return VerifyToken
}
