package output

import (
	"fmt"
	"strings"

	"github.com/whywaita/gigani/parse"
)

func Markdown(animes []parse.Anime, url string) string {
	// output Markdown table
	s := []string{}
	s = append(s, "# 録画リスト (auto generate by [whywaita/giagni](https://github.com/whywaita/gigani/))")

	u := fmt.Sprint("## target URL is %s", url)
	s = append(s, u)

	s = append(s, "| 名前 | 局 | 放送時間 | done |")
	s = append(s, "|---|---|---|---|")

	for _, anime := range animes {
		layout := "2006/01/02(Monday) 15:04〜"
		startDate := anime.StartDate.Format(layout)

		s = append(s, "|["+anime.Name+"]("+anime.URL+") |"+anime.BloadCaster+"|"+startDate+"||")
	}

	result := strings.Join(s, "\n")

	return result
}
