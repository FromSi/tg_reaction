package use_case

import (
	"github.com/fromsi/tg_reaction/internal/services"
	"github.com/fromsi/tg_reaction/pkg/json"
)

type BasePartyReactionUseCase struct {
	telebotMethodService services.TelebotMethodService
}

func NewBasePartyReactionUseCase(
	telebotMethodService services.TelebotMethodService,
) *BasePartyReactionUseCase {
	return &BasePartyReactionUseCase{
		telebotMethodService: telebotMethodService,
	}
}

// Execute sends a party reaction in Telegram.
func (receiver *BasePartyReactionUseCase) Execute(chatId int64, messageId int) error {
	return receiver.telebotMethodService.SetMessageReaction(chatId, messageId, json.Party)
}
