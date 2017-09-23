package main

import (
	"flag"
	"log"
)

type Anime struct {
	Name        string
	URL         string
	StartDate   string
	BloadCaster string
}

func main() {
	// parse args
	defaultURL := "http://gigazine.net/news/20170917-anime-2017autumn/"
	var url = flag.String("url", defaultURL, "target URL")
	flag.Parse()

	html, err := getHTML(*url)
	if err != nil {
		log.Fatal(err)
	}

	animes, err := parseAnime(html)
	if err != nil {
		log.Fatal(err)
	}

	// for _, anime := range animes {
	// 	fmt.Println(anime.Name)
	// 	fmt.Println(anime.URL)
	// 	fmt.Println(anime.BloadCaster)
	// 	fmt.Println(anime.StartDate)
	// 	fmt.Println("")
	// }

	outputMarkdown(animes)
}
