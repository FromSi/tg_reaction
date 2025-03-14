package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type LeftChatMemberRoute struct {
	cryLoudReactionUseCase use_case.CryLoudReactionUseCase
}

func NewLeftChatMemberRoute(cryLoudReactionUseCase use_case.CryLoudReactionUseCase) *LeftChatMemberRoute {
	return &LeftChatMemberRoute{cryLoudReactionUseCase: cryLoudReactionUseCase}
}

func (receiver *LeftChatMemberRoute) Endpoint() string {
	return telebot.OnUserLeft
}

// Handle processes user left events and sends a cry loud reaction.
func (receiver *LeftChatMemberRoute) Handle(context types.RouteContext) error {
	return receiver.cryLoudReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
