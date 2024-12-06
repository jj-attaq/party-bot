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

func Scrape(apifyUrl, token string, keywords []string) []models.InstagramPost {
	return filter(getPosts(apifyUrl, token), keywords)
	// return helperFilter(getPosts(apifyUrl, token))
}

func getBody(apifyUrl, token string) *http.Response {
	file, err := os.Open("input.json")
	if err != nil {
		// handle err
		log.Panic(err)
	}
	defer file.Close() // must be closed in the Scraper function!!!
	url := fmt.Sprintf("%s%s", apifyUrl, token)

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

func getPosts(apifyUrl, token string) []byte {
	b := getBody(apifyUrl, token).Body
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
		if containsWords(keywords, post) && !post.IsPinned && !post.DetermineIfTimedOut() {
			result = append(result, post)
		}
	}
	return result
}

// func helperFilter(postsBytes []byte) []models.InstagramPost {
// 	var posts []models.InstagramPost
// 	if err := json.Unmarshal(postsBytes, &posts); err != nil {
// 		fmt.Println("Cannot unmarshal JSON")
// 	}
//
// 	var result []models.InstagramPost
// 	for _, post := range posts {
// 		result = append(result, post)
// 	}
// 	return result
// }

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
