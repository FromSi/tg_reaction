package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type NewMemberRoute struct {
	partyReactionUseCase use_case.PartyReactionUseCase
}

func NewNewMemberRoute(partyReactionUseCase use_case.PartyReactionUseCase) *NewMemberRoute {
	return &NewMemberRoute{partyReactionUseCase: partyReactionUseCase}
}

func (receiver *NewMemberRoute) Endpoint() string {
	return telebot.OnUserJoined
}

// Handle processes new member events and sends a party reaction.
func (receiver *NewMemberRoute) Handle(context types.RouteContext) error {
	return receiver.partyReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
