package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type PinnedMessageRoute struct {
	eyesReactionUseCase use_case.EyesReactionUseCase
}

func NewPinnedMessageRoute(eyesReactionUseCase use_case.EyesReactionUseCase) *PinnedMessageRoute {
	return &PinnedMessageRoute{eyesReactionUseCase: eyesReactionUseCase}
}

func (receiver *PinnedMessageRoute) Endpoint() string {
	return telebot.OnPinned
}

// Handle processes pinned message events and sends an eyes reaction.
func (receiver *PinnedMessageRoute) Handle(context types.RouteContext) error {
	return receiver.eyesReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
