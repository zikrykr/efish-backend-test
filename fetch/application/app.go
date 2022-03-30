package application

import (
	"context"
	"fetch/api/currency"
	"fetch/api/fetch"
	"fetch/api/token"
	"fetch/cache"
	"fetch/config"
	customMiddleware "fetch/middleware"
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

	app.initCache()
	app.initRoutes()

	return app
}

func (app *App) initRoutes() {
	tokenDecoder := token.NewDecoder()

	currencyService := currency.NewService(&http.Client{Timeout: 10 * time.Second})
	fetchService := fetch.NewService(&http.Client{Timeout: 10 * time.Second}, currencyService)

	fetchController := fetch.NewController(app.E, fetchService)
	tokenController := token.NewController(app.E, tokenDecoder)

	v1 := app.E.Group("/v1/fetch")
	v1.GET("/resources", fetchController.HandleGetResources, customMiddleware.DecodeToken(tokenDecoder))
	v1.GET("/resources/aggregate", fetchController.HandleGetResourceAggregate, customMiddleware.DecodeTokenAdmin(tokenDecoder))

	v1.GET("/verify-token", tokenController.HandleVerifyJWT)
}

func (app *App) initCache() {
	cache := cache.NewCache()
	duration := time.NewTicker(time.Duration(app.config.CacheDuration) * time.Minute)

	go cache.Start(duration)
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
