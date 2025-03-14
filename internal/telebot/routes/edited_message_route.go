package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type EditedMessageRoute struct {
	clearReactionUseCase use_case.ClearReactionUseCase
	regexReactionUseCase use_case.RegexReactionUseCase
}

func NewEditedMessageRoute(
	clearReactionUseCase use_case.ClearReactionUseCase,
	regexReactionUseCase use_case.RegexReactionUseCase,
) *EditedMessageRoute {
	return &EditedMessageRoute{
		clearReactionUseCase: clearReactionUseCase,
		regexReactionUseCase: regexReactionUseCase,
	}
}

func (receiver *EditedMessageRoute) Endpoint() string {
	return telebot.OnEdited
}

// Handle processes edited messages by first clearing previous reactions and then setting a new one if needed
func (receiver *EditedMessageRoute) Handle(context types.RouteContext) error {
	chat := context.Chat()
	message := context.Message()

	if err := receiver.clearReactionUseCase.Execute(chat.ID, message.ID); err != nil {
		return err
	}

	return receiver.regexReactionUseCase.Execute(chat.ID, message.ID, context.Text())
}
