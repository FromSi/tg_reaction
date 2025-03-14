package adapters

import (
	"github.com/fromsi/tg_reaction/pkg/json"
	telebot "gopkg.in/telebot.v3"
)

//go:generate mockgen -destination=../../mocks/adapters/mock_tele_bot.go -package=adapters_mocks github.com/fromsi/tg_reaction/internal/adapters TeleBot
type TeleBot interface {
	Start()
	Stop()
	Handle(endpoint interface{}, h telebot.HandlerFunc, m ...telebot.MiddlewareFunc)
	Raw(method string, payload interface{}) ([]byte, error)
}

//go:generate mockgen -destination=../../mocks/adapters/mock_bot_adapter.go -package=adapters_mocks github.com/fromsi/tg_reaction/internal/adapters BotAdapter
type BotAdapter interface {
	SetMessageReaction(chatId int64, messageId int, reaction json.Reaction) error
}