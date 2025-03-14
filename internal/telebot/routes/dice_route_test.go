package routes

import (
	"testing"

	types_mocks "github.com/fromsi/tg_reaction/mocks/telebot/types"
	use_case_mocks "github.com/fromsi/tg_reaction/mocks/use_case"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	telebot "gopkg.in/telebot.v3"
)

func TestDiceRoute_Endpoint(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockDiceReactionUseCase := use_case_mocks.NewMockDiceReactionUseCase(mockController)
	route := NewDiceRoute(mockDiceReactionUseCase)

	assert.NotEmpty(t, route.Endpoint())
	assert.Equal(t, telebot.OnDice, route.Endpoint())
}

func TestDiceRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockDiceReactionUseCase := use_case_mocks.NewMockDiceReactionUseCase(mockController)
	route := NewDiceRoute(mockDiceReactionUseCase)

	mockMessage := &telebot.Message{
		ID: 67890,
		Dice: &telebot.Dice{
			Value: 6,
		},
	}

	mockDiceReactionUseCase.EXPECT().Execute(int64(12345), 67890, 6).Return(nil)

	mockRouteContext.EXPECT().Chat().Return(&telebot.Chat{ID: 12345})
	mockRouteContext.EXPECT().Message().Return(mockMessage).Times(2)

	err := route.Handle(mockRouteContext)
	assert.NoError(t, err)
}
