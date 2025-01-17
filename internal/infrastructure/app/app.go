package app

import (
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/config"
	"github.com/josepdcs/go-proposal-hexagonal-arch/internal/infrastructure/server/di"
	"github.com/pkg/errors"
)

func Start() error {
	cfg, err := config.Load()
	if err != nil {
		return errors.Wrapf(err, "cannot load config")
	}

	server, err := di.InitializeAPI(cfg.DB)
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
