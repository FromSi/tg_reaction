package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type TextRoute struct {
	regexReactionUseCase use_case.RegexReactionUseCase
}

func NewTextRoute(regexReactionUseCase use_case.RegexReactionUseCase) *TextRoute {
	return &TextRoute{regexReactionUseCase: regexReactionUseCase}
}

func (receiver *TextRoute) Endpoint() string {
	return telebot.OnText
}

// Handle processes text messages and sends a reaction.
func (receiver *TextRoute) Handle(context types.RouteContext) error {
	return receiver.regexReactionUseCase.Execute(context.Chat().ID, context.Message().ID, context.Text())
}
