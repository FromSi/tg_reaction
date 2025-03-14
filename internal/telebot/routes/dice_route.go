package routes

import (
	"github.com/fromsi/tg_reaction/internal/telebot/types"
	"github.com/fromsi/tg_reaction/internal/use_case"
	telebot "gopkg.in/telebot.v3"
)

type DiceRoute struct {
	diceReactionUseCase use_case.DiceReactionUseCase
}

func NewDiceRoute(diceReactionUseCase use_case.DiceReactionUseCase) *DiceRoute {
	return &DiceRoute{
		diceReactionUseCase: diceReactionUseCase,
	}
}

func (receiver *DiceRoute) Endpoint() string {
	return telebot.OnDice
}

// Handle processes dice events and delegates processing to the use case
func (receiver *DiceRoute) Handle(context types.RouteContext) error {
	return receiver.diceReactionUseCase.Execute(context.Chat().ID, context.Message().ID, context.Message().Dice.Value)
}
