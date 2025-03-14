package use_case

import (
	"github.com/fromsi/tg_reaction/internal/services"
	"github.com/fromsi/tg_reaction/pkg/json"
)

type BaseEyesReactionUseCase struct {
	telebotMethodService services.TelebotMethodService
}

func NewBaseEyesReactionUseCase(
	telebotMethodService services.TelebotMethodService,
) *BaseEyesReactionUseCase {
	return &BaseEyesReactionUseCase{
		telebotMethodService: telebotMethodService,
	}
}

// Execute sends an eyes reaction in Telegram.
func (receiver *BaseEyesReactionUseCase) Execute(chatId int64, messageId int) error {
	return receiver.telebotMethodService.SetMessageReaction(chatId, messageId, json.Eyes)
}
