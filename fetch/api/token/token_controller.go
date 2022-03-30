package token

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type TokenController struct {
	e            *echo.Echo
	tokenDecoder Decoder
}

func NewController(e *echo.Echo, tokenDecoder Decoder) *TokenController {
	return &TokenController{
		e,
		tokenDecoder,
	}
}

func (ctl *TokenController) HandleVerifyJWT(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

	if tokenString == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized Access")
	}

	claims, err := ctl.tokenDecoder.DecodeToken(tokenString)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid token")
	}

	return c.JSON(http.StatusOK, claims)
}
