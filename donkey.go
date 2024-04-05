package main

import (
    "fmt"
    "log"
    "os"
    "net/http"

    "github.com/jj-attaq/party-bot/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func donkey() {
	log.Println("Example actor written in Go.")


	default_kvs, token := os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN")
	if default_kvs == "" || token == "" {
		log.Fatal("Missing required env vars")
		return
	}

    apifyUrl := fmt.Sprintf("https://api.apify.com/v2/acts/apify~instagram-post-scraper/runs?token=%v", token)

    resp, err := http.Get(apifyUrl)
    if err != nil {
        log.Fatal(err)
        return
    }

	client := http.Client{}
	url := fmt.Sprintf("https://api.apify.com/v2/key-value-stores/%v/records/OUTPUT?token=%v", default_kvs, token)
	req, _ := http.NewRequest(http.MethodPut, url, resp.Body)
	req.Header.Set("Content-Type", "text/html; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Saved fetched html to OUTPUT in key-value store.")
    fmt.Println(resp)
}
