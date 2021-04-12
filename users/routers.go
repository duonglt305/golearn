package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(r *gin.RouterGroup) {
	auth := r.Group("auth")
	auth.POST("login", Login)
	auth.Use(JWTMiddleware())
	auth.GET("profile", ProfileRetrieve)
	auth.PUT("profile", ProfileUpdate)
}

func ProfileRetrieve(c *gin.Context) {
	serializer := ProfileSerializer{c}
	c.JSON(http.StatusOK, serializer.Response())
}

func ProfileUpdate(c *gin.Context) {
	validator := ProfileValidator{}
	if err := validator.Bind(c); err != nil {
		return
	}
	u := c.MustGet("user").(User)
	_ = Update(u.ID, validator.User)
	SetUserContext(c, u.ID)
	serializer := ProfileSerializer{c}
	c.JSON(http.StatusOK, serializer.Response())
}

func Login(c *gin.Context) {
	validator := NewLoginValidator()
	if err := validator.Bind(c); err != nil {
		_ = c.Error(err)
		return
	}
	u, err := FindOne(&User{Email: validator.Email})
	if err == nil && u.VerifyPassword(validator.Password) == nil {
		SetUserContext(c, u.ID)
		u.SetLoggedTime()
		serializer := LoginSerializer{c}
		c.JSON(http.StatusOK, serializer.Response())
		return
	}
	_ = c.Error(NewLoginError(err))
}
