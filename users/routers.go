package users

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golearn/common"
	"net/http"
)

func Routers(r *gin.RouterGroup) {
	auth := r.Group("auth")
	auth.POST("login", Login)
	auth.Use(JWTMiddleware(false))
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
		c.JSON(http.StatusUnprocessableEntity, common.NewValidationError(err))
		return
	}
	userId := c.MustGet("user_id").(uint)
	_ = Update(userId, validator.User)
	serializer := ProfileSerializer{c}
	c.JSON(http.StatusOK, serializer.Response())
}

func Login(c *gin.Context) {
	validator := NewLoginValidator()
	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidationError(err))
		return
	}
	u, err := FindOne(&User{Email: validator.Email})
	if err == nil && u.VerifyPassword(validator.Password) == nil {
		SetContextUser(c, u.ID)
		u.SetLoggedTime()
		serializer := LoginSerializer{c}
		c.JSON(http.StatusOK, serializer.Response())
		return
	}
	c.AbortWithStatusJSON(http.StatusForbidden, common.NewError("email", errors.New("not registered email or invalid password")))
}
