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
	var (
		flagURL          = flag.String("url", defaultURL, "target URL")
		flagOutputFormat = flag.String("output", "markdown", "output format")
	)
	flag.Parse()

	url := *flagURL
	outputFormat := *flagOutputFormat

	err := validateFlag(url, outputFormat)
	if err != nil {
		log.Fatal(err)
	}

	html, err := getHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	animes, err := parseAnime(html)
	if err != nil {
		log.Fatal(err)
	}

	if outputFormat == "markdown" {
		outputMarkdown(animes)
	}

}
