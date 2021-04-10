package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"golearn/common"
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
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		SetContextUser(c, 0)
		token, err := request.ParseFromRequest(c.Request, &request.MultiExtractor{
			&request.PostExtractionFilter{
				Extractor: request.HeaderExtractor{"Authorization"},
				Filter: func(token string) (string, error) {
					if len(token) > 5 && strings.ToUpper(token[0:6]) == "TOKEN " {
						return token[6:], nil
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

		}
		if token != nil && token.Valid {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				userId := uint(claims["id"].(float64))
				SetContextUser(c, userId)
			}
		}
	}
}
