package configs

import (
	"github.com/fromsi/tg_reaction/pkg/env"
)

const (
	BaseConfigDefaultAppTelegramToken = "secret"
)

type BaseConfig struct {
	appTelegramToken string
}

func NewBaseConfig() *BaseConfig {
	var config BaseConfig

	config.appTelegramToken = env.GetEnv("APP_TELEGRAM_TOKEN", BaseConfigDefaultAppTelegramToken)

	return &config
}

func (receiver BaseConfig) GetTelegramToken() string {
	return receiver.appTelegramToken
}
