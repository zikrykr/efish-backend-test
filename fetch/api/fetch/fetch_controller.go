package fetch

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FetchController struct {
	e            *echo.Echo
	fetchService FetchService
}

func NewController(e *echo.Echo, fetchService FetchService) *FetchController {
	return &FetchController{
		e,
		fetchService,
	}
}

func (ctl *FetchController) HandleGetResources(c echo.Context) error {
	fmt.Println("Getting Resources...")
	// TODO: Validate JWT
	ctx := c.Request().Context()
	result, err := ctl.fetchService.GetResources(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
