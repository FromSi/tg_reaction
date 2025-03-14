package services

import (
	"github.com/fromsi/tg_reaction/internal/adapters"
	"github.com/fromsi/tg_reaction/pkg/json"
)

type BaseTelebotMethodService struct {
	bot adapters.BotAdapter
}

func NewBaseTelebotMethodService(bot adapters.BotAdapter) *BaseTelebotMethodService {
	return &BaseTelebotMethodService{bot: bot}
}

func (receiver *BaseTelebotMethodService) SetMessageReaction(chatId int64, messageId int, reaction json.Reaction) error {
	return receiver.bot.SetMessageReaction(chatId, messageId, reaction)
}
