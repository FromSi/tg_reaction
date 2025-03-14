package use_case

import (
	"testing"

	services_mocks "github.com/fromsi/tg_reaction/mocks/services"
	"github.com/fromsi/tg_reaction/pkg/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBasePartyReactionUseCase_Execute(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)

	useCase := NewBasePartyReactionUseCase(mockTelebotMethodService)

	chatId := int64(12345)
	messageId := 67890

	mockTelebotMethodService.EXPECT().
		SetMessageReaction(chatId, messageId, json.Party).
		Return(nil)

	err := useCase.Execute(chatId, messageId)

	assert.NoError(t, err)
}

func TestBasePartyReactionUseCase_Execute_Error(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebotMethodService := services_mocks.NewMockTelebotMethodService(mockController)

	useCase := NewBasePartyReactionUseCase(mockTelebotMethodService)

	chatId := int64(12345)
	messageId := 67890
	expectedError := assert.AnError

	mockTelebotMethodService.EXPECT().
		SetMessageReaction(chatId, messageId, json.Party).
		Return(expectedError)

	err := useCase.Execute(chatId, messageId)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
