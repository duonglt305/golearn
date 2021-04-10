package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golearn/common"
	"net/http"
)

func Routers(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	auth.POST("/login", Login)
	auth.Use(JWTMiddleware(false))
	auth.GET("profile", func(c *gin.Context) {
		println(c.MustGet("user_id").(uint))
		c.JSON(http.StatusOK, gin.H{"user": "ahihi"})
	})
}

func Login(c *gin.Context) {
	validator := NewLoginValidator()
	if err := validator.Bind(c); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, common.NewValidationError(err))
	}
	var u User
	_, err := u.FindOne(&User{Email: validator.Email})
	if err == nil && u.VerifyPassword(validator.Password) == nil {
		SetContextUser(c, u.ID)
		u.SetLoggedTime()
		serializer := LoginSerializer{c}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"user": serializer.Response()})
	}
	c.AbortWithStatusJSON(http.StatusForbidden, common.NewError("email", errors.New("not registered email or invalid password")))
}
