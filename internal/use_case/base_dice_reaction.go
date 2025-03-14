package use_case

import (
	"github.com/fromsi/tg_reaction/internal/services"
	"github.com/fromsi/tg_reaction/pkg/json"
)

type BaseDiceReactionUseCase struct {
	telebotMethodService services.TelebotMethodService
	diceService          services.DiceService
}

func NewBaseDiceReactionUseCase(
	telebotMethodService services.TelebotMethodService,
	diceService services.DiceService,
) *BaseDiceReactionUseCase {
	return &BaseDiceReactionUseCase{
		telebotMethodService: telebotMethodService,
		diceService:          diceService,
	}
}

// Execute checks if dice value is winning and sends a trophy reaction if needed
func (receiver *BaseDiceReactionUseCase) Execute(chatId int64, messageId int, diceValue int) error {
	if receiver.diceService.IsWinningValue(diceValue) {
		return receiver.telebotMethodService.SetMessageReaction(chatId, messageId, json.Trophy)
	}

	return nil
}
