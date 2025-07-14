package main

import (
	"fmt"
	"net/http"

	"github.com/nishant007-tech/GoDig/internal/config"
	"github.com/nishant007-tech/GoDig/internal/database"
	"github.com/nishant007-tech/GoDig/internal/handler"
	"github.com/nishant007-tech/GoDig/internal/logger"
	"github.com/nishant007-tech/GoDig/internal/plugin"
	"github.com/nishant007-tech/GoDig/internal/repository"
	"github.com/nishant007-tech/GoDig/internal/server"
	"github.com/nishant007-tech/GoDig/internal/service"
	"go.uber.org/dig"
)

type PluginParams struct {
	dig.In
	Plugins []plugin.Plugin `group:"plugins"`
}

func main() {
	container := dig.New()

	// Register providers (constructors)
	container.Provide(config.NewConfig)
	container.Provide(logger.NewLogger)

	// Named DB instances
	container.Provide(database.NewPostgres, dig.Name("primaryDB"))
	container.Provide(database.NewMySQL, dig.Name("secondaryDB"))

	// Core layers
	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewUserService)
	container.Provide(handler.NewUserHandler)

	// Server builder (uses dig.Out)
	container.Provide(server.NewServer)

	// Grouped plugins
	container.Provide(plugin.NewPaymentPlugin, dig.Group("plugins"))
	container.Provide(plugin.NewAnalyticsPlugin, dig.Group("plugins"))

	// Invoke the app
	if err := container.Invoke(run); err != nil {
		panic(err)
	}
}

func run(
	cfg *config.Config,
	log *logger.Logger,
	handler http.Handler,
	cleanup func() error,
	pluginParams PluginParams,
) error {
	plugins := pluginParams.Plugins
	// Initialize all plugins (dig.Group)
	for _, p := range plugins {
		log.Info("Initializing plugin:", p.Name())
		if err := p.Init(); err != nil {
			return err
		}
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Info("Server listening on", addr)

	defer cleanup()

	// Start HTTP server with injected Handler
	return http.ListenAndServe(addr, handler)
}
