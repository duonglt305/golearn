package users

import "os/user"

type LoginValidator struct {
	User struct {
		Email    string `form:"email" json:"email" binding:"exists,email"`
		Password string `form:"password" json:"password" binding:"exists,min:8,max:255"`
	} `json:"user"`
	model user.User `json:"-"`
}
