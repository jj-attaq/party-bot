package models

import (
	"fmt"
	"time"
)

func (p InstagramPost) DetermineIfTimedOut() bool {
	now := time.Now().Truncate(time.Second)
	loadLoc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	currentUTC := now.In(loadLoc)
	pTime := p.Timestamp

	ct := currentUTC.Truncate(time.Hour)
	pt := pTime.Truncate(time.Hour)

	fmt.Printf("curr timestamp: %s, type is: %T\n", currentUTC, currentUTC)
	fmt.Printf("post timestamp: %s, type is: %T\n", pTime, pTime)

	// IF CURRENT TIME > PTIME + 3DAYS RETURN FALSE
	if pt.Before(ct.AddDate(0, 0, -3)) {
		fmt.Println("givenTime is older than 3 days")
		return true
	} else {
		fmt.Println("givenTime is not older than 3 days")
		return false
	}
}

type InstagramPost struct {
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
