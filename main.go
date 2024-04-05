package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jj-attaq/party-bot/initializers"
)

type InstagramPosts []struct {
	InputURL       string   `json:"inputUrl"`
	ID             string   `json:"id"`
	Type           string   `json:"type"`
	ShortCode      string   `json:"shortCode"`
	Caption        string   `json:"caption"`
	Hashtags       []any    `json:"hashtags"`
	Mentions       []string `json:"mentions"`
	URL            string   `json:"url"`
	CommentsCount  int      `json:"commentsCount"`
	FirstComment   string   `json:"firstComment"`
	LatestComments []struct {
		ID                 string    `json:"id"`
		Text               string    `json:"text"`
		OwnerUsername      string    `json:"ownerUsername"`
		OwnerProfilePicURL string    `json:"ownerProfilePicUrl"`
		Timestamp          time.Time `json:"timestamp"`
		LikesCount         int       `json:"likesCount"`
		RepliesCount       int       `json:"repliesCount,omitempty"`
		Replies            []struct {
			ID                 string    `json:"id"`
			Text               string    `json:"text"`
			OwnerUsername      string    `json:"ownerUsername"`
			OwnerProfilePicURL string    `json:"ownerProfilePicUrl"`
			Timestamp          time.Time `json:"timestamp"`
			LikesCount         int       `json:"likesCount"`
		} `json:"replies,omitempty"`
	} `json:"latestComments"`
	DimensionsHeight int       `json:"dimensionsHeight"`
	DimensionsWidth  int       `json:"dimensionsWidth"`
	DisplayURL       string    `json:"displayUrl"`
	Images           []any     `json:"images"`
	VideoURL         string    `json:"videoUrl"`
	Alt              any       `json:"alt"`
	LikesCount       int       `json:"likesCount"`
	VideoViewCount   int       `json:"videoViewCount"`
	VideoPlayCount   int       `json:"videoPlayCount"`
	Timestamp        time.Time `json:"timestamp"`
	ChildPosts       []any     `json:"childPosts"`
	OwnerFullName    string    `json:"ownerFullName"`
	OwnerUsername    string    `json:"ownerUsername"`
	OwnerID          string    `json:"ownerId"`
	ProductType      string    `json:"productType"`
	VideoDuration    float64   `json:"videoDuration"`
	IsSponsored      bool      `json:"isSponsored"`
	TaggedUsers      []struct {
		FullName      string `json:"full_name"`
		ID            string `json:"id"`
		IsVerified    bool   `json:"is_verified"`
		ProfilePicURL string `json:"profile_pic_url"`
		Username      string `json:"username"`
	} `json:"taggedUsers"`
	IsPinned  bool `json:"isPinned"`
	MusicInfo struct {
		ArtistName            string `json:"artist_name"`
		SongName              string `json:"song_name"`
		UsesOriginalAudio     bool   `json:"uses_original_audio"`
		ShouldMuteAudio       bool   `json:"should_mute_audio"`
		ShouldMuteAudioReason string `json:"should_mute_audio_reason"`
		AudioID               string `json:"audio_id"`
	} `json:"musicInfo"`
	CoauthorProducers []struct {
		ID            string `json:"id"`
		IsVerified    bool   `json:"is_verified"`
		ProfilePicURL string `json:"profile_pic_url"`
		Username      string `json:"username"`
	} `json:"coauthorProducers"`
}

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	default_kvs, token, keywords:= os.Getenv("APIFY_DEFAULT_KEY_VALUE_STORE_ID"), os.Getenv("APIFY_TOKEN"), strings.Split(os.Getenv("FILTER_KEYWORDS"), ",")
	if default_kvs == "" || token == "" {
		log.Fatal("Missing required env vars")
		return
    }

    fmt.Println(keywords)

    // fmt.Printf("type of keywords is %T\n", keywords)

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

    // respDump, err := httputil.DumpResponse(resp, true)
    // if err != nil {
    //     log.Fatal(err)
    // } 
    // 
    // fmt.Printf("RESPONSE: \n%s\n", string(respDump))

    body, err :=  io.ReadAll(resp.Body)

    var result InstagramPosts
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Cannot unmarshal JSON")
    }

    fmt.Println(PrettyPrint(result))
}

func PrettyPrint(i interface{}) string {
    s, _ := json.MarshalIndent(i, "", "\t")
    return string(s)
}
