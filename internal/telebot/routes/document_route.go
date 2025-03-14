package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type DocumentRoute struct {
	regexReactionUseCase use_case.RegexReactionUseCase
}

func NewDocumentRoute(regexReactionUseCase use_case.RegexReactionUseCase) *DocumentRoute {
	return &DocumentRoute{regexReactionUseCase: regexReactionUseCase}
}

func (receiver *DocumentRoute) Endpoint() string {
	return telebot.OnDocument
}

// Handle processes documents and sends a reaction.
func (receiver *DocumentRoute) Handle(context types.RouteContext) error {
	return receiver.regexReactionUseCase.Execute(context.Chat().ID, context.Message().ID, context.Message().Document.UniqueID)
}
