package routes

import (
	"testing"

	types_mocks "github.com/fromsi/tg_reaction/mocks/telebot/types"
	use_case_mocks "github.com/fromsi/tg_reaction/mocks/use_case"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	telebot "gopkg.in/telebot.v3"
)

func TestChatPhotoDeletedRoute_Endpoint(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockCryLoudReactionUseCase := use_case_mocks.NewMockCryLoudReactionUseCase(mockController)
	route := NewChatPhotoDeletedRoute(mockCryLoudReactionUseCase)

	assert.NotEmpty(t, route.Endpoint())
	assert.Equal(t, telebot.OnGroupPhotoDeleted, route.Endpoint())
}

func TestChatPhotoDeletedRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockCryLoudReactionUseCase := use_case_mocks.NewMockCryLoudReactionUseCase(mockController)
	route := NewChatPhotoDeletedRoute(mockCryLoudReactionUseCase)

	mockMessage := &telebot.Message{
		ID: 67890,
	}

	mockCryLoudReactionUseCase.EXPECT().Execute(int64(12345), 67890).Return(nil)
	mockRouteContext.EXPECT().Chat().Return(&telebot.Chat{ID: 12345})
	mockRouteContext.EXPECT().Message().Return(mockMessage)

	err := route.Handle(mockRouteContext)
	assert.NoError(t, err)
}
