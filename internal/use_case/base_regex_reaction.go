package use_case

import (
	"github.com/fromsi/tg_reaction/internal/services"
)

type BaseRegexReactionUseCase struct {
	telebotMethodService services.TelebotMethodService
	regexService         services.RegexService
}

func NewBaseRegexReactionUseCase(
	telebotMethodService services.TelebotMethodService,
	regexService services.RegexService,
) *BaseRegexReactionUseCase {
	return &BaseRegexReactionUseCase{
		telebotMethodService: telebotMethodService,
		regexService:         regexService,
	}
}

// Execute sends a reaction in Telegram.
func (receiver *BaseRegexReactionUseCase) Execute(chatId int64, messageId int, text string) error {
	reaction := receiver.regexService.FindReaction(text)

	if reaction == "" {
		return nil
	}

	return receiver.telebotMethodService.SetMessageReaction(chatId, messageId, reaction)
}
