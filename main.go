package main

import (
	"fmt"
	"log"
)

type Anime struct {
	Name        string
	URL         string
	StartDate   string
	BloadCaster string
}

func main() {
	// anime 2017 autumn
	url := "http://gigazine.net/news/20170917-anime-2017autumn/"

	html, err := getHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	animes, err := parseAnime(html)
	if err != nil {
		log.Fatal(err)
	}

	for _, anime := range animes {
		fmt.Println(anime.Name)
		fmt.Println(anime.URL)
		fmt.Println(anime.BloadCaster)
		fmt.Println(anime.StartDate)
		fmt.Println("")
	}
}
