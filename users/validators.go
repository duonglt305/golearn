package users

import "github.com/gin-gonic/gin"

type LoginValidator struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=255"`
}

func NewLoginValidator(c *gin.Context) (*LoginValidator, error) {
	v := &LoginValidator{}
	if err := c.ShouldBind(v); err != nil {
		_ = c.Error(err)
		return v, err
	}
	return v, nil
}

type ProfileValidator struct {
	Name string `form:"name" json:"name" binding:"required,max=255"`
	User User   `json:"-"`
}

func NewProfileValidator(c *gin.Context) (*ProfileValidator, error) {
	v := &ProfileValidator{}
	if err := c.ShouldBind(v); err != nil {
		_ = c.Error(err)
		return v, err
	}
	v.User.Name = v.Name
	return v, nil
}
