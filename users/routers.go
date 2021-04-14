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
	v, err := NewProfileValidator(c)
	if err != nil {
		return
	}
	u := c.MustGet("user").(User)
	_ = Update(u.ID, v.User)
	SetUserContext(c, u.ID)
	serializer := ProfileSerializer{c}
	c.JSON(http.StatusOK, serializer.Response())
}

func Login(c *gin.Context) {
	v, err := NewLoginValidator(c)
	if err != nil {
		return
	}
	u, err := FindOne(&User{Email: v.Email})
	if err == nil && u.VerifyPassword(v.Password) == nil {
		SetUserContext(c, u.ID)
		u.SetLoggedTime()
		serializer := LoginSerializer{c}
		c.JSON(http.StatusOK, serializer.Response())
		return
	}
	_ = c.Error(NewLoginError(err))
}
