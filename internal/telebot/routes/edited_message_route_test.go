package routes

import (
	"testing"

	types_mocks "github.com/fromsi/tg_reaction/mocks/telebot/types"
	use_case_mocks "github.com/fromsi/tg_reaction/mocks/use_case"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	telebot "gopkg.in/telebot.v3"
)

func TestEditedMessageRoute_Endpoint(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockClearReactionUseCase := use_case_mocks.NewMockClearReactionUseCase(mockController)
	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)

	route := NewEditedMessageRoute(mockClearReactionUseCase, mockRegexReactionUseCase)

	assert.NotEmpty(t, route.Endpoint())
	assert.Equal(t, telebot.OnEdited, route.Endpoint())
}

func TestEditedMessageRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockClearReactionUseCase := use_case_mocks.NewMockClearReactionUseCase(mockController)
	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)

	route := NewEditedMessageRoute(mockClearReactionUseCase, mockRegexReactionUseCase)

	chatId := int64(12345)
	messageId := 67890
	messageText := "Edited message text"

	mockChat := &telebot.Chat{ID: chatId}
	mockMessage := &telebot.Message{ID: messageId}

	mockRouteContext.EXPECT().Chat().Return(mockChat)
	mockRouteContext.EXPECT().Message().Return(mockMessage)
	mockRouteContext.EXPECT().Text().Return(messageText)

	mockClearReactionUseCase.EXPECT().
		Execute(chatId, messageId).
		Return(nil)

	mockRegexReactionUseCase.EXPECT().
		Execute(chatId, messageId, messageText).
		Return(nil)

	err := route.Handle(mockRouteContext)
	assert.NoError(t, err)
}

func TestEditedMessageRoute_Handle_ClearError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockClearReactionUseCase := use_case_mocks.NewMockClearReactionUseCase(mockController)
	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)

	route := NewEditedMessageRoute(mockClearReactionUseCase, mockRegexReactionUseCase)

	chatId := int64(12345)
	messageId := 67890
	expectedError := assert.AnError

	mockChat := &telebot.Chat{ID: chatId}
	mockMessage := &telebot.Message{ID: messageId}

	mockRouteContext.EXPECT().Chat().Return(mockChat)
	mockRouteContext.EXPECT().Message().Return(mockMessage)

	mockClearReactionUseCase.EXPECT().
		Execute(chatId, messageId).
		Return(expectedError)

	err := route.Handle(mockRouteContext)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
}
