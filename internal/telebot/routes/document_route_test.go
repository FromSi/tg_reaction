package routes

import (
	"testing"

	types_mocks "github.com/fromsi/tg_reaction/mocks/telebot/types"
	use_case_mocks "github.com/fromsi/tg_reaction/mocks/use_case"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	telebot "gopkg.in/telebot.v3"
)

func TestDocumentRoute_Endpoint(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)
	route := NewDocumentRoute(mockRegexReactionUseCase)

	assert.NotEmpty(t, route.Endpoint())

	assert.Equal(t, route.Endpoint(), telebot.OnDocument)
}

func TestDocumentRoute_Handle(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockRouteContext := types_mocks.NewMockRouteContext(mockController)
	mockRegexReactionUseCase := use_case_mocks.NewMockRegexReactionUseCase(mockController)
	route := NewDocumentRoute(mockRegexReactionUseCase)

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
	mockRouteContext.EXPECT().Message().Return(mockMessage).Times(2)

	err := route.Handle(mockRouteContext)
	assert.NoError(t, err)
}
