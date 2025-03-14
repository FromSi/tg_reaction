package configs

//go:generate mockgen -destination=../../mocks/configs/mock_config.go -package=configs_mocks github.com/fromsi/tg_reaction/internal/configs AppConfig
type AppConfig interface {
	GetTelegramToken() string
}

//go:generate mockgen -destination=../../mocks/configs/mock_config.go -package=configs_mocks github.com/fromsi/tg_reaction/internal/configs Config
type Config interface {
	AppConfig
}