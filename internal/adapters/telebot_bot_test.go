package adapters

import (
	"errors"
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

	chatId := int64(12345)
	messageId := 67890
	reaction := json.Fire

	mockTelebot.EXPECT().Raw("setMessageReaction", gomock.Eq(
		map[string]interface{}{
			"chat_id":    chatId,
			"message_id": messageId,
			"is_big":     true,
			"reaction": []map[string]string{
				{
					"type":  "emoji",
					"emoji": string(reaction),
				},
			},
		})).Return(nil, nil)

	err := telebotBotAdapter.SetMessageReaction(chatId, messageId, reaction)
	assert.NoError(t, err, "Expected no error, got %v", err)
}

func TestTelebotBotAdapter_SetMessageReaction_EmptyReaction(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebot := adapters_mocks.NewMockTeleBot(mockController)
	telebotBotAdapter := NewTelebotBotAdapter(mockTelebot)

	chatId := int64(12345)
	messageId := 67890
	reaction := json.Empty

	mockTelebot.EXPECT().Raw("setMessageReaction", gomock.Eq(
		map[string]interface{}{
			"chat_id":    chatId,
			"message_id": messageId,
			"is_big":     true,
			"reaction":   []map[string]string{},
		})).Return(nil, nil)

	err := telebotBotAdapter.SetMessageReaction(chatId, messageId, reaction)

	assert.NoError(t, err, "Expected no error, got %v", err)
}

func TestTelebotBotAdapter_SetMessageReaction_IgnoresReactionEmptyError(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockTelebot := adapters_mocks.NewMockTeleBot(mockController)
	telebotBotAdapter := NewTelebotBotAdapter(mockTelebot)

	chatId := int64(12345)
	messageId := 67890
	reaction := json.Empty
	reactionEmptyError := errors.New("Bad Request: REACTION_EMPTY (400)")

	mockTelebot.EXPECT().Raw("setMessageReaction", gomock.Any()).Return(nil, reactionEmptyError)

	err := telebotBotAdapter.SetMessageReaction(chatId, messageId, reaction)

	assert.NoError(t, err, "Expected no error, got %v", err)
}
