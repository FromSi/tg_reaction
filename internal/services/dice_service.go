package services

//go:generate mockgen -destination=../../mocks/services/mock_dice_service.go -package=services_mocks github.com/fromsi/tg_reaction/internal/services DiceService
type DiceService interface {
	IsWinningValue(value int) bool
}
