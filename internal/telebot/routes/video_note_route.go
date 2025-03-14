package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type VideoNoteRoute struct {
	poopReactionUseCase use_case.PoopReactionUseCase
}

func NewVideoNoteRoute(poopReactionUseCase use_case.PoopReactionUseCase) *VideoNoteRoute {
	return &VideoNoteRoute{
		poopReactionUseCase: poopReactionUseCase,
	}
}

func (receiver *VideoNoteRoute) Endpoint() string {
	return telebot.OnVideoNote
}

// Handle processes video note messages and sends a poop reaction
func (receiver *VideoNoteRoute) Handle(context types.RouteContext) error {
	return receiver.poopReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
