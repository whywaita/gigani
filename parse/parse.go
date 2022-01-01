package parse

import (
	"time"
)

var (
	// ListBroadcaster is list of Broadcaster (priority order)
	ListBroadcaster = []string{"TOKYO MX", "テレビ東京", "フジテレビ", "日本テレビ", "tvk", "TBS"}
)

const (
	BroadCasterNotFound = "指定されていない放送局です"
)

type Anime struct {
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	StartDate   time.Time `json:"start_date"`
	BloadCaster string    `json:"broadcaster"`
}
