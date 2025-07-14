package server

import (
	"net/http"

	"github.com/nishant007-tech/GoDig/internal/config"
	"github.com/nishant007-tech/GoDig/internal/handler"
	"github.com/nishant007-tech/GoDig/internal/logger"
	"go.uber.org/dig"
)

// ServerParams defines inputs for NewServer.
type ServerParams struct {
	dig.In
	Cfg     *config.Config
	Log     *logger.Logger
	Handler *handler.UserHandler
}

// ServerOut defines multiple outputs from NewServer.
type ServerOut struct {
	dig.Out
	Handler http.Handler
	Cleanup func() error
}

// NewServer constructs the HTTP handler and a cleanup function.
func NewServer(p ServerParams) ServerOut {
	p.Log.Info("Setting up HTTP server")
	mux := http.NewServeMux()
	mux.HandleFunc("/users", p.Handler.Handle)
	return ServerOut{
		Handler: mux,
		Cleanup: func() error {
			p.Log.Info("Cleaning up server resources")
			return nil
		},
	}
}
