package adapters

import (
	"strings"

	"github.com/fromsi/tg_reaction/pkg/json"
)

type TelebotBotAdapter struct {
	bot TeleBot
}

func NewTelebotBotAdapter(bot TeleBot) *TelebotBotAdapter {
	return &TelebotBotAdapter{bot: bot}
}

func (receiver *TelebotBotAdapter) SetMessageReaction(chatId int64, messageId int, reaction json.Reaction) error {
	params := map[string]interface{}{
		"chat_id":    chatId,
		"message_id": messageId,
		"is_big":     true,
		"reaction":   []map[string]string{},
	}

	if reaction != json.Empty {
		params["reaction"] = []map[string]string{
			{
				"type":  "emoji",
				"emoji": string(reaction),
			},
		}
	}

	_, err := receiver.bot.Raw("setMessageReaction", params)

	// Ignore the REACTION_EMPTY error, which occurs
	// when trying to remove a reaction that doesn't exist
	if err != nil && strings.Contains(err.Error(), "Bad Request: REACTION_EMPTY") {
		return nil
	}

	return err
}
