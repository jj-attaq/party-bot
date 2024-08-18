package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jj-attaq/party-bot/apify"
	"github.com/jj-attaq/party-bot/chatbot"
	"github.com/jj-attaq/party-bot/initializers"
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

	// fmt.Println("Scraping...")

	posts := apify.Scrape(apifyToken, keywords)
	if telegramToken != "" {
		fmt.Println(posts)
	}
	chatbot.TelegramBot(telegramToken, posts)
}
