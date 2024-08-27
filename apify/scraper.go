package apify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jj-attaq/party-bot/models"
)

func Scrape(token string, keywords []string) []models.InstagramPost {
	return filter(getPosts(token), keywords)
}

func getBody(token string) *http.Response {
	file, err := os.Open("input.json")
	if err != nil {
		// handle err
		log.Panic(err)
	}
	defer file.Close() // must be closed in the Scraper function!!!
	// move to a .env variable
	url := fmt.Sprintf("https://api.apify.com/v2/acts/apify~instagram-post-scraper/run-sync-get-dataset-items?token=%s", token)

	req, err := http.NewRequest("POST", os.ExpandEnv(url), file)
	if err != nil {
		// handle err
		log.Panicln(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Panicln(err)
	}
	// defer resp.Body.Close() // closed in Scraper function!!!

	return resp
}

func getPosts(token string) []byte {
	b := getBody(token).Body
	defer b.Close()
	postsBytes, err := io.ReadAll(b)
	if err != nil {
		log.Panic(err)
	}

	return postsBytes
}

func filter(postsBytes []byte, keywords []string) []models.InstagramPost {
	var posts []models.InstagramPost
	if err := json.Unmarshal(postsBytes, &posts); err != nil {
		fmt.Println("Cannot unmarshal JSON")
	}

	var result []models.InstagramPost
	for _, post := range posts {
		if containsWords(keywords, post) && !post.IsPinned {
			result = append(result, post)
		}
	}
	return result
}

func containsWords(keywords []string, post models.InstagramPost) bool {
	for _, word := range keywords {
		if strings.Contains(strings.ToLower(post.Caption), strings.ToLower(word)) {
			return true
		}
	}
	return false
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
