package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type MediaRoute struct {
	regexReactionUseCase use_case.RegexReactionUseCase
}

func NewMediaRoute(regexReactionUseCase use_case.RegexReactionUseCase) *MediaRoute {
	return &MediaRoute{regexReactionUseCase: regexReactionUseCase}
}

func (receiver *MediaRoute) Endpoint() string {
	return telebot.OnMedia
}

// Handle processes media and sends a reaction.
func (receiver *MediaRoute) Handle(context types.RouteContext) error {
	var text string

	if context.Message().Document != nil {
		text = context.Message().Document.UniqueID
	}

	if context.Message().Caption != "" {
		if text != "" {
			text += " "
		}

		text += context.Message().Caption
	}

	return receiver.regexReactionUseCase.Execute(context.Chat().ID, context.Message().ID, text)
}
