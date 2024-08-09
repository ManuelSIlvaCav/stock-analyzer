package main

import (
	"context"
	"fmt"
	"stockanalyzer/internal"
	"stockanalyzer/internal/container"
	"stockanalyzer/internal/container/config"
	"stockanalyzer/internal/router"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func PrintHello() {
	fmt.Println("Hello, Modules! This is api command speaking!")
}

func NewServer() *echo.Echo {
	return echo.New()
}

func registerHooks(lifecycle fx.Lifecycle, e *echo.Echo, container *container.Container, router *router.Router, moduleContainer *internal.ModulesContainer) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// logger := container.GetLogger()

			// logger.Info(fmt.Sprintf("Server started on :%s asd", container.GetConfig().Port))

			go e.Start(fmt.Sprintf(":%s", container.GetConfig().Port))

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
}

func main() {
	fmt.Println("Hello, Modules!")

	newConfig := config.NewConfig()

	fmt.Println("name is a host v3  = ", newConfig.Postgres.Host)

	PrintHello()

	fx.New(fx.Options(
		container.Modules,
		internal.Modules,
		router.Modules,
		fx.Provide(NewServer),
		fx.Invoke(registerHooks),
	)).Run()
}
