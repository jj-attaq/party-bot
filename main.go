package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jj-attaq/party-bot/initializers"
	"github.com/jj-attaq/party-bot/chatbot"
	"github.com/jj-attaq/party-bot/apify"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	default_kvs, apifyToken, telegramToken, keywords := os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN"), os.Getenv("TELEGRAM_TOKEN"), strings.Split(os.Getenv("FILTER_KEYWORDS"), ",")
	if default_kvs == "" || apifyToken == "" {
		log.Fatal("Missing required env vars")
		return
    }

    fmt.Println("Scraping...")

    posts := apify.Scraper(default_kvs, apifyToken, keywords)
    chatbot.TelegramBot(telegramToken, posts)
}
