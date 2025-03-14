package use_case

import (
	"testing"

	services_mocks "github.com/fromsi/tg_reaction/mocks/services"
	"github.com/fromsi/tg_reaction/pkg/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBaseDiceReactionUseCase_Execute_WinningValue(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)
	mockDiceService := services_mocks.NewMockDiceService(mockController)

	useCase := NewBaseDiceReactionUseCase(mockTelebotMethodService, mockDiceService)

	chatId := int64(12345)
	messageId := 67890
	diceValue := 6

	mockDiceService.EXPECT().
		IsWinningValue(diceValue).
		Return(true)

	mockTelebotMethodService.EXPECT().
		SetMessageReaction(chatId, messageId, json.Trophy).
		Return(nil)

	err := useCase.Execute(chatId, messageId, diceValue)

	assert.NoError(t, err)
}

func TestBaseDiceReactionUseCase_Execute_NonWinningValue(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)
	mockDiceService := services_mocks.NewMockDiceService(mockController)

	useCase := NewBaseDiceReactionUseCase(mockTelebotMethodService, mockDiceService)

	chatId := int64(12345)
	messageId := 67890
	diceValue := 3

	mockDiceService.EXPECT().
		IsWinningValue(diceValue).
		Return(false)

	err := useCase.Execute(chatId, messageId, diceValue)

	assert.NoError(t, err)
}

func TestBaseDiceReactionUseCase_Execute_Error(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)
	mockDiceService := services_mocks.NewMockDiceService(mockController)

	useCase := NewBaseDiceReactionUseCase(mockTelebotMethodService, mockDiceService)

	chatId := int64(12345)
	messageId := 67890
	diceValue := 6
	expectedError := assert.AnError

	mockDiceService.EXPECT().
		IsWinningValue(diceValue).
		Return(true)

	mockTelebotMethodService.EXPECT().
		SetMessageReaction(chatId, messageId, json.Trophy).
		Return(expectedError)

	err := useCase.Execute(chatId, messageId, diceValue)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
