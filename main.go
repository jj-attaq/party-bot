package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
//
// 	"github.com/jj-attaq/party-bot/apify"
// 	"github.com/jj-attaq/party-bot/chatbot"
// 	"github.com/jj-attaq/party-bot/initializers"
// )

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jj-attaq/party-bot/chatbot"
	"github.com/jj-attaq/party-bot/initializers"
	"github.com/jj-attaq/party-bot/models"
)

func init() {
	initializers.LoadEnvVariables()
}

// func main() {
// 	default_kvs, apifyToken, telegramToken, keywords, url := os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN"), os.Getenv("TELEGRAM_TOKEN"), strings.Split(os.Getenv("FILTER_KEYWORDS"), ","), os.Getenv("APIFY_URL")
// 	if default_kvs == "" || apifyToken == "" || url == "" {
// 		log.Fatal("Missing required env vars")
// 		return
// 	}
//
// 	// fmt.Println("Scraping...")
//
// 	posts := apify.Scrape(url, apifyToken, keywords)
// 	if telegramToken != "" {
// 		fmt.Println(posts)
// 	}
//
// 	file, _ := json.MarshalIndent(posts, "", "\t")
// 	_ = os.WriteFile("output.json", file, 0644)
//
// 	chatbot.TelegramBot(telegramToken, posts)
// }

func main() {
	default_kvs, apifyToken, telegramToken := os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN"), os.Getenv("TELEGRAM_TOKEN")
	if default_kvs == "" || apifyToken == "" {
		log.Fatal("Missing required env vars")
		return
	}

	content, err := os.ReadFile("output.json")
	if err != nil {
		log.Fatal(err)
	}

	posts := []models.InstagramPost{}
	err = json.Unmarshal(content, &posts)
	if err != nil {
		panic(err)
	}
	for _, post := range posts {
		log.Println(post.DetermineIfTimedOut())
	}

	chatbot.TelegramBot(telegramToken, posts)
}

// func main() {
// 	default_kvs, apifyToken, telegramToken, keywords, url := os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN"), os.Getenv("TELEGRAM_TOKEN"), strings.Split(os.Getenv("FILTER_KEYWORDS"), ","), os.Getenv("APIFY_URL")
// 	if default_kvs == "" || apifyToken == "" || url == "" {
// 		log.Fatal("Missing required env vars")
// 		return
// 	}
//
// 	// fmt.Println("Scraping...")
//
// 	posts := apify.Scrape(url, apifyToken, keywords)
// 	if telegramToken != "" {
// 		fmt.Println(posts)
// 	}
// 	chatbot.TelegramBot(telegramToken, posts)
// }
