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

func Scraper(default_kvs string, token string, keywords []string) ([]byte, []models.InstagramPosts) {
	// log.Println("file: ", file)

	b := getBody(token).Body
	defer b.Close()
	body, err := io.ReadAll(b)
	// log.Println("body: ", string(body))
	if err != nil {
		log.Panic(err)
	}

	var posts []models.InstagramPosts
	if err := json.Unmarshal(body, &posts); err != nil {
		fmt.Println("Cannot unmarshal JSON")
	}

	var result []models.InstagramPosts
	for _, post := range posts {

		// if containsWords(keywords, post) && !post.IsPinned {
		// 	result = append(result, post)
		// 	fmt.Println(prettyPrint(post.Timestamp))
		// 	fmt.Println(prettyPrint(post.URL))
		// 	fmt.Println(prettyPrint(post.IsPinned))
		// 	fmt.Println("")
		// }
		result = append(result, post)
	}
	// fmt.Println(prettyPrint(result))
	return body, result
}

func containsWords(keywords []string, post models.InstagramPosts) bool {
	for _, word := range keywords {
		if strings.Contains(post.Caption, word) {
			return true
		}
	}
	return false
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
