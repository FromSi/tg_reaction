package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type ChatPhotoDeletedRoute struct {
	cryLoudReactionUseCase use_case.CryLoudReactionUseCase
}

func NewChatPhotoDeletedRoute(cryLoudReactionUseCase use_case.CryLoudReactionUseCase) *ChatPhotoDeletedRoute {
	return &ChatPhotoDeletedRoute{cryLoudReactionUseCase: cryLoudReactionUseCase}
}

func (receiver *ChatPhotoDeletedRoute) Endpoint() string {
	return telebot.OnGroupPhotoDeleted
}

// Handle processes chat photo deletion events and sends a cry loud reaction.
func (receiver *ChatPhotoDeletedRoute) Handle(context types.RouteContext) error {
	return receiver.cryLoudReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
