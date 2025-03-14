package services

import (
	"github.com/fromsi/tg_reaction/pkg/json"
)

//go:generate mockgen -destination=../../mocks/services/mock_regex_service.go -package=services_mocks github.com/fromsi/tg_reaction/internal/services RegexService
type RegexService interface {
	FindReaction(text string) json.Reaction
}
