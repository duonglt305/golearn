package users

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"golearn/common"
	"net/http"
	"os"
	"strings"
)

func SetContextUser(c *gin.Context, id uint) {
	var user User
	if id != 0 {
		db := common.GetDB()
		db.First(&user, id)
	}
	c.Set("user_id", id)
	c.Set("user", user)
}
func JWTMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		SetContextUser(c, 0)
		token, err := request.ParseFromRequest(c.Request, &request.MultiExtractor{
			&request.PostExtractionFilter{
				Extractor: request.HeaderExtractor{"Authorization"},
				Filter: func(token string) (string, error) {
					if len(token) > 5 && strings.ToUpper(token[0:6]) == "BEARER" {
						return strings.TrimSpace(token[6:]), nil
					}
					return token, nil
				},
			},
			request.ArgumentExtractor{"access_token"},
		}, func(token *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("APP_KEY"))
			return b, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewError("email", errors.New("unauthenticated user")))
		}
		if token != nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				userId := uint(claims["user_id"].(float64))
				SetContextUser(c, userId)
			}
		}
	}
}
