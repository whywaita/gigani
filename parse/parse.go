package parse

import (
	"time"
)

type Anime struct {
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	StartDate   time.Time `json:"start_date"`
	BloadCaster string    `json:"broadcaster"`
}
