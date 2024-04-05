package apify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jj-attaq/party-bot/models"
)

func prettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}

func Scraper(default_kvs string, token string, keywords []string) {
    f, err := os.Open("input.json")
    if err != nil {
        // handle err
        log.Panicln(err)
    }
    defer f.Close()

    // move to a .env variable
    url := fmt.Sprintf("https://api.apify.com/v2/acts/apify~instagram-post-scraper/run-sync-get-dataset-items?token=%s", token)

    req, err := http.NewRequest("POST", os.ExpandEnv(url), f)
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
    defer resp.Body.Close()

    body, err :=  io.ReadAll(resp.Body)

    var result models.InstagramPosts
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Cannot unmarshal JSON")
    }

    for _, posts := range result {
        fmt.Println(prettyPrint(posts.Timestamp))
        fmt.Println(prettyPrint(posts.URL))
        fmt.Println("")
    }
    // fmt.Println(prettyPrint(result))
}
