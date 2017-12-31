package cli

import (
	"flag"
	"log"

	"github.com/whywaita/gigani/lib"
	"github.com/whywaita/gigani/output"
	"github.com/whywaita/gigani/parse"
)

func Start(args []string) {
	// parse args
	defaultURL := ""
	defaultFormat := ""
	var (
		flagURL          = flag.String("url", defaultURL, "target URL")
		flagOutputFormat = flag.String("output", defaultFormat, "output format, can use markdown/json")
	)
	flag.Parse()

	url := *flagURL
	outputFormat := *flagOutputFormat

	err := validateFlag(url, outputFormat)
	if err != nil {
		log.Fatal(err)
	}

	html, err := lib.GetHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	animes, err := parse.ParseAnime(html)
	if err != nil {
		log.Fatal(err)
	}

	if outputFormat == "markdown" {
		output.Markdown(animes, url)
	} else if outputFormat == "json" {
		output.JSON(animes)
	}

}
