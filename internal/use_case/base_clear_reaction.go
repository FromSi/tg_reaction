package use_case

import (
	"github.com/fromsi/tg_reaction/internal/services"
	"github.com/fromsi/tg_reaction/pkg/json"
)

type BaseClearReactionUseCase struct {
	telebotMethodService services.TelebotMethodService
}

func NewBaseClearReactionUseCase(
	telebotMethodService services.TelebotMethodService,
) *BaseClearReactionUseCase {
	return &BaseClearReactionUseCase{
		telebotMethodService: telebotMethodService,
	}
}

// Execute clears all reactions from a message by sending empty reaction
func (receiver *BaseClearReactionUseCase) Execute(chatId int64, messageId int) error {
	return receiver.telebotMethodService.SetMessageReaction(chatId, messageId, json.Empty)
}
