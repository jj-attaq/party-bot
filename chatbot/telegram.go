package chatbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jj-attaq/party-bot/models"
)

func TelegramBot(telegramToken string, apifyData []models.InstagramPost) {
	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			userMessage := update.Message.Text
			switch userMessage {
			case "/dance":
				if len(apifyData) < 1 {
					danceMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Looks like there are no parties coming up. No dancing for you!")
					bot.Send(danceMsg)
					danceMsg.ReplyToMessageID = update.Message.MessageID
				} else {
					for _, post := range apifyData {
						danceMsg := tgbotapi.NewMessage(update.Message.Chat.ID, post.URL)
						// log.Println("TIMESTAMP INFO: ")
						// post.DetermineIfTimedOut()
						log.Println("CAPTION: ")
						log.Println(post.Caption)
						bot.Send(danceMsg)
						danceMsg.ReplyToMessageID = update.Message.MessageID
					}
				}
			default:
				bot.Send(msg)
			}

			if update.Message.Text == "/dance" && len(apifyData) < 1 {
			} else {
			}
		}
	}
}
