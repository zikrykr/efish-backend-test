package middleware

import (
	"fetch/api/token"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func DecodeTokenAdmin(t token.Decoder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			tokenString := c.Request().Header.Get("Authorization")
			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized Access")
			}

			claims, err := t.DecodeToken(tokenString)
			if err != nil {
				return err
			}

			if claims.Role != "admin" {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden role")
			}

			return next(c)
		}
	}
}

func DecodeToken(t token.Decoder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

			if tokenString == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized Access")
			}

			claims, err := t.DecodeToken(tokenString)
			if err != nil {
				return err
			}

			if claims.Role == "" {
				return echo.NewHTTPError(http.StatusForbidden, "invalid role")
			}

			return next(c)
		}
	}
}
