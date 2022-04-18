package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type Config struct {
	AwsDefaultRegion string `env:"AWS_DEFAULT_REGION"`
	AwsAccessKeyId   string `env:"AWS_ACCESS_KEY_ID"`
}

func ReadConfig() *Config {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		log.Err(err).Msg("")
	}
	return &config
}
