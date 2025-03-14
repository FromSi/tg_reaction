package adapters

import (
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

	return err
}
