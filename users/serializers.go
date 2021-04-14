package users

import (
	"github.com/gin-gonic/gin"
	"time"
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
		AccessToken: GenToken(u),
	}
	return user
}

type ProfileSerializer struct {
	c *gin.Context
}

type ProfileResponse struct {
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	Photo          string    `json:"photo"`
	LatestLoggedAt time.Time `json:"latest_logged_at"`
}

func (p ProfileSerializer) Response() ProfileResponse {
	user := p.c.MustGet("user").(User)
	profile := ProfileResponse{
		Email:          user.Email,
		Name:           user.Name,
		Photo:          user.Photo,
		LatestLoggedAt: user.LatestLoggedAt,
	}
	return profile
}
