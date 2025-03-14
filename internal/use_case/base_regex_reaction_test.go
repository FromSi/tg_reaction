package use_case

import (
	"testing"

	services_mocks "github.com/fromsi/tg_reaction/mocks/services"
	"github.com/fromsi/tg_reaction/pkg/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBaseRegexReactionUseCase_Execute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)
	mockRegexService := services_mocks.NewMockRegexService(mockController)

	useCase := NewBaseRegexReactionUseCase(mockTelebotMethodService, mockRegexService)

	chatId := int64(12345)
	messageId := 67890
	text := "Привет, как дела?"
	reaction := json.Heart

	mockRegexService.EXPECT().FindReaction(text).Return(reaction)
	mockTelebotMethodService.EXPECT().SetMessageReaction(chatId, messageId, reaction).Return(nil)

	err := useCase.Execute(chatId, messageId, text)

	assert.NoError(t, err)
}

func TestBaseRegexReactionUseCase_Execute_NoReactionFound(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)
	mockRegexService := services_mocks.NewMockRegexService(mockController)

	useCase := NewBaseRegexReactionUseCase(mockTelebotMethodService, mockRegexService)

	chatId := int64(12345)
	messageId := 67890
	text := "Какой-то случайный текст"

	mockRegexService.EXPECT().FindReaction(text).Return(json.Reaction(""))

	err := useCase.Execute(chatId, messageId, text)

	assert.NoError(t, err)
}

func TestBaseRegexReactionUseCase_Execute_Error(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)
	mockRegexService := services_mocks.NewMockRegexService(mockController)

	useCase := NewBaseRegexReactionUseCase(mockTelebotMethodService, mockRegexService)

	chatId := int64(12345)
	messageId := 67890
	text := "Привет, как дела?"
	reaction := json.Heart
	expectedError := assert.AnError

	mockRegexService.EXPECT().FindReaction(text).Return(reaction)
	mockTelebotMethodService.EXPECT().SetMessageReaction(chatId, messageId, reaction).Return(expectedError)

	err := useCase.Execute(chatId, messageId, text)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
