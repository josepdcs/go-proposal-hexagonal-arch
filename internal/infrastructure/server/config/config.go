package config

import (
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2/log"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
)

const (
	// ConfigPathEnv is the environment variable that contains the path to the configuration file.
	ConfigPathEnv         = "CONFIG_PATH"
	ConfigOverridePathEnv = "CONFIG_OVERRIDE_PATH"

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
	SSLMode  string `koanf:"ssl-mode"`
}

func Load() (Config, error) {
	var config Config

	k := koanf.New(".")
	parser := yaml.Parser()

	rootDir, _ := os.Getwd()

	configPath := filepath.Join(rootDir, "config.yml")
	if value, exists := os.LookupEnv(ConfigPathEnv); exists {
		configPath = value
	}
	log.Debugf("Config file path: %s", configPath)

	// Load YML config.
	if err := k.Load(file.Provider(configPath), parser); err != nil {
		return config, errors.Wrapf(err, "error loading config: %v", err)
	}

	configOverridePath := filepath.Join(rootDir, "config_override.yml")
	if value, exists := os.LookupEnv(ConfigOverridePathEnv); exists {
		configOverridePath = value
	}
	if configOverridePath != "" {
		log.Debugf("Config override file path: %s", configOverridePath)

		if err := k.Load(file.Provider(configOverridePath), parser); err != nil {
			return config, errors.Wrapf(err, "error loading config: %v", err)
		}
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
