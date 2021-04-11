package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

type JwtClaims struct {
	ID uint
	jwt.StandardClaims
}

func GenToken(id uint) string {
	lifetime, _ := strconv.Atoi(os.Getenv("JWT_LIFETIME"))
	now := time.Now().UTC()
	claims := &JwtClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Hour * time.Duration(lifetime)).Unix(),
			NotBefore: now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(os.Getenv("APP_KEY")))
	return signedToken
}

func VerifyToken(c *gin.Context) uint {
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
	}, request.WithClaims(&JwtClaims{}))
	if err != nil {
		_ = c.Error(NewUnauthenticatedError(err))
		c.Abort()
	}
	if token != nil && token.Valid {
		claims := token.Claims.(*JwtClaims)
		return claims.ID
	}
	return 0
}
