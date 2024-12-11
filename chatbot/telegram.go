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
			} else {
				bot.Send(msg)
			}
		}
	}
}

// // RequestFileData represents the data to be used for a file.
// type RequestFileData interface {
// 	// NeedsUpload shows if the file needs to be uploaded.
// 	NeedsUpload() bool
//
// 	// UploadData gets the file name and an `io.Reader` for the file to be uploaded. This
// 	// must only be called when the file needs to be uploaded.
// 	UploadData() (string, io.Reader, error)
// 	// SendData gets the file data to send when a file does not need to be uploaded. This
// 	// must only be called when the file does not need to be uploaded.
// 	SendData() string
// }
//
