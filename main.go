package main

import (
	"log"
	"os"
	"strings"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TG_REACTION_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	regexWithEmoji, err := NewRegexWithEmoji(os.Getenv("TG_REACTION_DATA"))

	if err != nil {
		log.Fatal(err)
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(tele.OnText, func(context tele.Context) error {
		text := context.Text()

		sendReaction(bot, context, text, regexWithEmoji)

		return nil
	})

	bot.Handle(tele.OnDocument, func(context tele.Context) error {
		text := context.Message().Document.UniqueID

		sendReaction(bot, context, text, regexWithEmoji)

		return nil
	})

	bot.Handle(tele.OnMedia, func(context tele.Context) error {
		var text strings.Builder

		text.WriteString(context.Message().Document.UniqueID)
		text.WriteRune(' ')
		text.WriteString(context.Message().Caption)

		sendReaction(bot, context, text.String(), regexWithEmoji)

		return nil
	})

	bot.Handle(tele.OnSticker, func(context tele.Context) error {
		text := context.Message().Sticker.UniqueID

		sendReaction(bot, context, text, regexWithEmoji)

		return nil
	})

	bot.Start()
}

func sendReaction(bot *tele.Bot, context tele.Context, text string, regexWithEmoji *RegexWithEmoji) {
	chatId := context.Chat().ID
	messageId := context.Message().ID

	emoji := regexWithEmoji.GetEmoji(text)

	if emoji != "" {
		params := map[string]interface{}{
			"chat_id":    chatId,
			"message_id": messageId,
			"reaction": []map[string]string{
				{
					"type":  "emoji",
					"emoji": emoji,
				},
			},
		}

		_, err := bot.Raw("setMessageReaction", params)

		if err != nil {
			log.Println(err)
		}
	}
}
