package app

import (
	"time"

	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/di"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/http"
	"github.com/pkg/errors"
)

// server is the application server. It can be used to start, stop, and access the Fiber app instance.
var server *http.Server

// Start starts the application.
func Start() error {
	cfg, err := config.Load()
	if err != nil {
		return errors.Wrapf(err, "cannot load config")
	}

	server, err = di.InitializeAPI(cfg.DB)
	if err != nil {
		return errors.Wrapf(err, "cannot initialize server")
	} else {
		err = server.Start()
		if err != nil {
			return errors.Wrapf(err, "cannot start server")
		}
	}

	return nil
}

// Shutdown stops the application.
func Shutdown() error {
	if server == nil {
		// The server is not initialized, so there is nothing to shut down.
		return nil
	}
	return server.Shutdown()
}

// ShutdownWithTimeout stops the application with a timeout.
func ShutdownWithTimeout(timeout time.Duration) error {
	if server == nil {
		// The server is not initialized, so there is nothing to shut down.
		return nil
	}
	return server.ShutdownWithTimeout(timeout)
}
