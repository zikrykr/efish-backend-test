package application

import (
	"context"
	"fetch/api/fetch"
	"fetch/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

type App struct {
	E      *echo.Echo
	config *config.Config
}

func New(config *config.Config) *App {
	app := &App{
		config: config,
		E:      echo.New(),
	}

	app.initRoutes()

	return app
}

func (app *App) initRoutes() {
	fetchService := fetch.NewService(&http.Client{Timeout: 10 * time.Second})

	fetchController := fetch.NewController(app.E, fetchService)

	v1 := app.E.Group("/v1/fetch")
	v1.GET("/resources", fetchController.HandleGetResources)
	v1.GET("/resources/aggregate", fetchController.HandleGetResourceAggregate)
}

// Start the server and handle graceful shutdown
func (app *App) Start() {
	app.E.HideBanner = true

	// Start server
	go func() {
		if err := app.E.Start(":" + app.config.AppPort); err != nil {
			app.E.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.E.Shutdown(ctx); err != nil {
		app.E.Logger.Fatal(err)
	}
}
