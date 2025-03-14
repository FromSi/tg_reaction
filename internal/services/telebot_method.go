package services

import "github.com/fromsi/tg_reaction/pkg/json"

//go:generate mockgen -destination=../../mocks/services/mock_query_telebot_method.go -package=services_mocks github.com/fromsi/tg_reaction/internal/services QueryTelebotMethodService
type QueryTelebotMethodService interface {}

//go:generate mockgen -destination=../../mocks/services/mock_mutable_telebot_method.go -package=services_mocks github.com/fromsi/tg_reaction/internal/services MutableTelebotMethodService
type MutableTelebotMethodService interface{
	SetMessageReaction(chatId int64, messageId int, reaction json.Reaction) error
}

//go:generate mockgen -destination=../../mocks/services/mock_telebot_method.go -package=services_mocks github.com/fromsi/tg_reaction/internal/services TelebotMethodService
type TelebotMethodService interface {
	QueryTelebotMethodService
	MutableTelebotMethodService
}