package middleware

import (
	"fetch/config"
	"fetch/model"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func DecodeTokenAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := config.GetConfig()

			tokenString := c.Request().Header.Get("Authorization")
			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized Access")
			}

			token, _ := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(conf.SecretKey), nil
			})

			if token == nil || token.Claims == nil {
				return echo.NewHTTPError(http.StatusForbidden, "invalid token")
			}

			claims := token.Claims.(*model.TokenClaims)

			if claims.Role != "admin" {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden role")
			}

			return next(c)
		}
	}
}

func DecodeToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := config.GetConfig()

			tokenString := c.Request().Header.Get("Authorization")
			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized Access")
			}

			token, _ := jwt.ParseWithClaims(tokenString, &model.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(conf.SecretKey), nil
			})

			if token == nil || token.Claims == nil {
				return echo.NewHTTPError(http.StatusForbidden, "invalid token")
			}

			claims := token.Claims.(*model.TokenClaims)

			if claims.Role == "" {
				return echo.NewHTTPError(http.StatusForbidden, "invalid role")
			}

			return next(c)
		}
	}
}
