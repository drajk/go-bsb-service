package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/drajk/go-bsb-service/pkg/logger"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type envConfig struct {
	ServiceName string `env:"SERVICE_NAME"`
	ServerPort  int    `env:"SERVER_PORT"`
	Version     string `env:"VERSION"`
	LogLevel    string `env:"LOG_LEVEL"`
}

func (c envConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ServiceName, validation.Required, validation.Length(1, 50), is.ASCII),
		validation.Field(&c.ServerPort, validation.Required, validation.Min(80), validation.Max(65535)),
		validation.Field(&c.Version, validation.Required, validation.Length(1, 50), is.Alphanumeric),
		validation.Field(&c.LogLevel, validation.Required, validation.Length(4, 5), is.Alpha),
	)
}

func newEnvironmentConfig() *envConfig {
	cfg := &envConfig{}
	if err := env.Parse(cfg); err != nil {
		logger.WithError(err).Fatal("cannot find configs for server")
	}

	errors := cfg.Validate()
	if errors != nil {
		logger.WithError(errors).Fatal("configuration validation failed with errors")
	}

	return cfg
}
