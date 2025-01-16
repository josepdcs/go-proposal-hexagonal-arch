package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
)

const (
	InMemoryDB = "in_memory"
)

type Config struct {
	DB DB `koanf:"db"`
}

type DB struct {
	Type     string `koanf:"type"`
	Host     string `koanf:"host"`
	Name     string `koanf:"name"`
	User     string `koanf:"user"`
	Port     string `koanf:"port"`
	Password string `koanf:"password"`
}

var envs = []string{
	"DB_HOST", "DB_NAME", "DB_USER", "DB_PORT", "DB_PASSWORD",
}

func Load() (Config, error) {
	var config Config

	k := koanf.New(".")
	parser := yaml.Parser()

	// Load YML config.
	if err := k.Load(file.Provider("config.yml"), parser); err != nil {
		return config, errors.Wrapf(err, "error loading config: %v", err)
	}

	if err := k.Load(file.Provider("config_override.yml"), parser); err != nil {
		return config, errors.Wrapf(err, "error loading config: %v", err)
	}

	err := k.Unmarshal("config", &config)
	if err != nil {
		return config, err
	}

	if config.DB.Type == "" {
		config.DB.Type = InMemoryDB
	}

	return config, nil
}
