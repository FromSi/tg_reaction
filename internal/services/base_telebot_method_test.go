package services

import (
	"testing"

	adapters_mocks "github.com/fromsi/tg_reaction/mocks/adapters"
	"github.com/fromsi/tg_reaction/pkg/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestBaseTelebotMethodService_SetMessageReaction(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockBotAdapter := adapters_mocks.NewMockBotAdapter(mockController)

	telebotMethodService := NewBaseTelebotMethodService(mockBotAdapter)

	chatId := int64(12345)
	messageId := 67890
	reaction := json.Fire

	mockBotAdapter.EXPECT().
		SetMessageReaction(chatId, messageId, reaction).
		Return(nil)

	err := telebotMethodService.SetMessageReaction(chatId, messageId, reaction)

	assert.NoError(t, err)
}

func TestBaseTelebotMethodService_SetMessageReaction_Error(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockBotAdapter := adapters_mocks.NewMockBotAdapter(mockController)
	telebotMethodService := NewBaseTelebotMethodService(mockBotAdapter)

	chatId := int64(12345)
	messageId := 67890
	reaction := json.Fire
	expectedError := assert.AnError

	mockBotAdapter.EXPECT().
		SetMessageReaction(chatId, messageId, reaction).
		Return(expectedError)

	err := telebotMethodService.SetMessageReaction(chatId, messageId, reaction)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
