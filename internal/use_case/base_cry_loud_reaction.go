package use_case

import (
	"github.com/fromsi/tg_reaction/internal/services"
	"github.com/fromsi/tg_reaction/pkg/json"
)

type BaseCryLoudReactionUseCase struct {
	telebotMethodService services.TelebotMethodService
}

func NewBaseCryLoudReactionUseCase(
	telebotMethodService services.TelebotMethodService,
) *BaseCryLoudReactionUseCase {
	return &BaseCryLoudReactionUseCase{
		telebotMethodService: telebotMethodService,
	}
}

// Execute sends a cry loud reaction in Telegram.
func (receiver *BaseCryLoudReactionUseCase) Execute(chatId int64, messageId int) error {
	return receiver.telebotMethodService.SetMessageReaction(chatId, messageId, json.CryLoud)
}
