package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	tele "gopkg.in/telebot.v3"
)

var (
	emoji      = []string{"👍", "👎", "❤", "🔥", "🥰", "👏", "😁", "🤔", "🤯", "😱", "🤬", "😢", "🎉", "🤩", "🤮", "💩", "🙏", "👌", "🕊", "🤡", "🥱", "🥴", "😍", "🐳", "❤‍🔥", "🌚", "🌭", "💯", "🤣", "⚡", "🍌", "🏆", "💔", "🤨", "😐", "🍓", "🍾", "💋", "🖕", "😈", "😴", "😭", "🤓", "👻", "👨‍💻", "👀", "🎃", "🙈", "😇", "😨", "🤝", "✍", "🤗", "🫡", "🎅", "🎄", "☃", "💅", "🤪", "🗿", "🆒", "💘", "🙉", "🦄", "😘", "💊", "🙊", "😎", "👾", "🤷‍♂", "🤷", "🤷‍♀", "😡"}
	emojiTotal = len(emoji)
	regex      = regexp.MustCompile(os.Getenv("TG_REACTION_REGEX"))
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TG_REACTION_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)

	if err != nil {
		log.Fatal(err)
	}

	bot.Handle(tele.OnText, func(context tele.Context) error {
		text := context.Text()

		sendReaction(bot, context, text)

		return nil
	})

	bot.Handle(tele.OnDocument, func(context tele.Context) error {
		text := context.Message().Document.UniqueID

		sendReaction(bot, context, text)

		return nil
	})

	bot.Handle(tele.OnSticker, func(context tele.Context) error {
		text := context.Message().Sticker.UniqueID

		log.Println(context.Message().Sticker.UniqueID)

		sendReaction(bot, context, text)

		return nil
	})

	bot.Start()
}

func sendReaction(bot *tele.Bot, context tele.Context, text string) {
	userId := context.Sender().Recipient()
	messageId := context.Message().ID

	matches := regex.FindAllString(text, -1)

	if matches != nil {
		params := map[string]interface{}{
			"chat_id":    userId,
			"message_id": messageId,
			"reaction": []map[string]string{
				{
					"type":  "emoji",
					"emoji": emoji[rand.Intn(emojiTotal)],
				},
			},
		}

		_, err := bot.Raw("setMessageReaction", params)

		if err != nil {
			log.Println(err)
		}
	}
}
