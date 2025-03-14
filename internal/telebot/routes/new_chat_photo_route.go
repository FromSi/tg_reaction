package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type NewChatPhotoRoute struct {
	partyReactionUseCase use_case.PartyReactionUseCase
}

func NewNewChatPhotoRoute(partyReactionUseCase use_case.PartyReactionUseCase) *NewChatPhotoRoute {
	return &NewChatPhotoRoute{partyReactionUseCase: partyReactionUseCase}
}

func (receiver *NewChatPhotoRoute) Endpoint() string {
	return telebot.OnNewGroupPhoto
}

// Handle processes new chat photo updates and sends a party reaction.
func (receiver *NewChatPhotoRoute) Handle(context types.RouteContext) error {
	return receiver.partyReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
