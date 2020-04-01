package gigazine

import (
	"strings"
	"time"

	"github.com/whywaita/gigani/parse"
)

func getYear(sentence string) (year string, err error) {
	r := strings.NewReader(sentence)
	htmlTitle, _ := GetHtmlTitle(r)

	s := strings.Split(htmlTitle, "年")
	year = s[0]

	return year, nil
}

func parseTime(times string, year string) (time.Time, error) {
	return parse.ParseTime(times, year)
}
