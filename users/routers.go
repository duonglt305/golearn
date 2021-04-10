package users

import "github.com/gin-gonic/gin"

func Routers(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", Login)
}

func Login(c *gin.Context) {

}
