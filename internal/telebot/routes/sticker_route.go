package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type StickerRoute struct {
	regexReactionUseCase use_case.RegexReactionUseCase
}

func NewStickerRoute(regexReactionUseCase use_case.RegexReactionUseCase) *StickerRoute {
	return &StickerRoute{regexReactionUseCase: regexReactionUseCase}
}

func (receiver *StickerRoute) Endpoint() string {
	return telebot.OnSticker
}

// Handle processes stickers and sends a reaction.
func (receiver *StickerRoute) Handle(context types.RouteContext) error {
	return receiver.regexReactionUseCase.Execute(context.Chat().ID, context.Message().ID, context.Message().Sticker.UniqueID)
}
