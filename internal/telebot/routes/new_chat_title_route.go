package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type NewChatTitleRoute struct {
	partyReactionUseCase use_case.PartyReactionUseCase
}

func NewNewChatTitleRoute(partyReactionUseCase use_case.PartyReactionUseCase) *NewChatTitleRoute {
	return &NewChatTitleRoute{partyReactionUseCase: partyReactionUseCase}
}

func (receiver *NewChatTitleRoute) Endpoint() string {
	return telebot.OnNewGroupTitle
}

// Handle processes new chat title updates and sends a party reaction.
func (receiver *NewChatTitleRoute) Handle(context types.RouteContext) error {
	return receiver.partyReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
