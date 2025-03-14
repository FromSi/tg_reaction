package use_case

import (
	"github.com/fromsi/tg_reaction/internal/services"
	"github.com/fromsi/tg_reaction/pkg/json"
)

type BasePoopReactionUseCase struct {
	telebotMethodService services.TelebotMethodService
}

func NewBasePoopReactionUseCase(
	telebotMethodService services.TelebotMethodService,
) *BasePoopReactionUseCase {
	return &BasePoopReactionUseCase{
		telebotMethodService: telebotMethodService,
	}
}

// Execute sends a poop reaction in Telegram.
func (receiver *BasePoopReactionUseCase) Execute(chatId int64, messageId int) error {
	return receiver.telebotMethodService.SetMessageReaction(chatId, messageId, json.Poop)
}
