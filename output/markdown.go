package output

import (
	"fmt"

	"github.com/whywaita/gigani/parse"
)

func Markdown(animes []parse.Anime, url string) {
	// output Markdown table
	fmt.Println(`# 録画リスト (auto generate by [whywaita/giagni](https://github.com/whywaita/gigani/))`)
	fmt.Printf(`## target URL is %s`, url)
	fmt.Println(`| 名前 | 局 | 放送時間 | done |`)
	fmt.Println(`|---|---|---|---|`)

	for _, anime := range animes {
		layout := "2006/01/02(Monday) 15:04〜"
		startDate := anime.StartDate.Format(layout)

		fmt.Println(`|[` + anime.Name + `](` + anime.URL + `) |` + anime.BloadCaster + `|` + startDate + `||`)
	}
}
