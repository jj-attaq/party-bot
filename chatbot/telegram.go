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

	// for update := range updates {
	// 	if update.Message != nil { // If we got a message
	// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	//
	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	// 		msg.ReplyToMessageID = update.Message.MessageID
	//
	// 		bot.Send(msg)
	// 	}
	// }
	for update := range updates {
		if update.Message != nil { // If we got a message
			// log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			if update.Message.Text == "/dance" {
				for _, post := range apifyData {
					danceMsg := tgbotapi.NewMessage(update.Message.Chat.ID, post.URL)
					log.Println("TIMESTAMP INFO: ")
					post.DetermineIfTimedOut()
					bot.Send(danceMsg)
					danceMsg.ReplyToMessageID = update.Message.MessageID
				}
			} else {
				bot.Send(msg)
			}
		}
	}
}
