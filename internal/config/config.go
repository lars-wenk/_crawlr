package config

import (
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = ""

type Config struct {
	AppEnv                 string `envconfig:"APP_ENV" default:"develop"`
	LogLevel               string `envconfig:"LOG_LEVEL" default:"info"`
	ComdirectClientID      string `envconfig:"COMDIRECT_CLIENT_ID"`
	ComdirectSecretID      string `envconfig:"COMDIRECT_SECRET_ID"`
	ComdirectZugangsnummer string `envconfig:"COMDIRECT_ZUGANGSNUMMER"`
	ComdirectTAN           string `envconfig:"COMDIRECT_TAN"`
}

func NewConfig() (Config, error) {
	var conf Config
	err := envconfig.Process(envPrefix, &conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
