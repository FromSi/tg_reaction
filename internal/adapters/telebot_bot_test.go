package adapters

import (
	"testing"

	adapters_mocks "github.com/fromsi/tg_reaction/mocks/adapters"
	"github.com/fromsi/tg_reaction/pkg/json"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTelebotBotAdapter_SetMessageReaction(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	
	mockTelebot := adapters_mocks.NewMockTeleBot(mockController)
	telebotBotAdapter := NewTelebotBotAdapter(mockTelebot)

	chatId := int64(123456789)
	messageId := 123
	reaction := json.ThumbsUp

	mockTelebot.EXPECT().Raw("setMessageReaction", gomock.Eq(
		map[string]interface{}{
			"chat_id":    chatId,
			"message_id": messageId,
			"reaction": []map[string]string{
				{"type": "emoji", "emoji": string(reaction)},
			},
		},
	)).Return(nil, nil)

	err := telebotBotAdapter.SetMessageReaction(chatId, messageId, reaction)
	assert.NoError(t, err, "Expected no error, got %v", err)
}
