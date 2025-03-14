package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type VoiceRoute struct {
	poopReactionUseCase use_case.PoopReactionUseCase
}

func NewVoiceRoute(poopReactionUseCase use_case.PoopReactionUseCase) *VoiceRoute {
	return &VoiceRoute{
		poopReactionUseCase: poopReactionUseCase,
	}
}

func (receiver *VoiceRoute) Endpoint() string {
	return telebot.OnVoice
}

// Handle processes voice messages and sends a poop reaction
func (receiver *VoiceRoute) Handle(context types.RouteContext) error {
	return receiver.poopReactionUseCase.Execute(context.Chat().ID, context.Message().ID)
}
