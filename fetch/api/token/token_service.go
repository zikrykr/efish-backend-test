package token

import (
	"fetch/config"
	"fetch/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type token struct{}

type Decoder interface {
	DecodeToken(tokenStr string) (*model.TokenClaims, error)
}

func NewDecoder() Decoder {
	return &token{}
}

func (t *token) DecodeToken(tokenStr string) (*model.TokenClaims, error) {
	conf := config.GetConfig()
	token, _ := jwt.ParseWithClaims(tokenStr, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.SecretKey), nil
	})

	if token == nil || token.Claims == nil {
		return nil, echo.NewHTTPError(http.StatusForbidden, "invalid token")
	}

	claims := token.Claims.(*model.TokenClaims)

	return claims, nil
}
