package routes

import (
	"testing"

	types_mocks "github.com/fromsi/tg_reaction/mocks/telebot/types"
	use_case_mocks "github.com/fromsi/tg_reaction/mocks/use_case"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	telebot "gopkg.in/telebot.v3"
)

func TestMediaRoute_Endpoint(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)
	route := NewMediaRoute(mockRegexReactionUseCase)

	assert.NotEmpty(t, route.Endpoint())

	assert.Equal(t, route.Endpoint(), telebot.OnMedia)
}

func TestMediaRoute_Handle_WithDocument(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)
	route := NewMediaRoute(mockRegexReactionUseCase)

	mockMessage := &telebot.Message{
		ID: 67890,
		Document: &telebot.Document{
			File: telebot.File{
				UniqueID: "test-document-id",
			},
		},
	}

	mockRegexReactionUseCase.EXPECT().Execute(int64(12345), 67890, "test-document-id").Return(nil)
	mockRouteContext.EXPECT().Chat().Return(&telebot.Chat{ID: 12345})
	mockRouteContext.EXPECT().Message().Return(mockMessage).Times(4)

	err := route.Handle(mockRouteContext)
	assert.NoError(t, err)
}

func TestMediaRoute_Handle_WithCaption(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)
	route := NewMediaRoute(mockRegexReactionUseCase)

	mockMessage := &telebot.Message{
		ID:      67890,
		Caption: "test caption",
	}

	mockRegexReactionUseCase.EXPECT().Execute(int64(12345), 67890, "test caption").Return(nil)
	mockRouteContext.EXPECT().Chat().Return(&telebot.Chat{ID: 12345})
	mockRouteContext.EXPECT().Message().Return(mockMessage).Times(4)

	err := route.Handle(mockRouteContext)
	assert.NoError(t, err)
}

func TestMediaRoute_Handle_WithDocumentAndCaption(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)
	route := NewMediaRoute(mockRegexReactionUseCase)

	mockMessage := &telebot.Message{
		ID:      67890,
		Caption: "test caption",
		Document: &telebot.Document{
			File: telebot.File{
				UniqueID: "test-document-id",
			},
		},
	}

	mockRegexReactionUseCase.EXPECT().Execute(int64(12345), 67890, "test-document-id test caption").Return(nil)
	mockRouteContext.EXPECT().Chat().Return(&telebot.Chat{ID: 12345})
	mockRouteContext.EXPECT().Message().Return(mockMessage).Times(5)

	err := route.Handle(mockRouteContext)
	assert.NoError(t, err)
}
