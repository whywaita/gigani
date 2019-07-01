package cli

import (
	"flag"
	"fmt"
	"log"

	"github.com/whywaita/gigani/parse/gigazine"

	"github.com/whywaita/gigani/format"
	"github.com/whywaita/gigani/lib"
)

func Start(args []string) {
	// parse args
	defaultURL := ""
	defaultFormat := ""
	defaultSortType := "post"
	var (
		flagURL          = flag.String("url", defaultURL, "[required] target URL")
		flagOutputFormat = flag.String("format", defaultFormat, "[required] output format, *markdown / json*")
		flagSortType     = flag.String("sort", defaultSortType, "[optional] sort base *post / time*")
	)
	flag.Parse()

	url := *flagURL
	outputFormat := *flagOutputFormat
	sortType := *flagSortType

	err := validateFlag(url, outputFormat, sortType)
	if err != nil {
		log.Fatal(err)
	}

	html, err := lib.GetHTML(url)
	if err != nil {
		log.Fatal(err)
	}

	animes, err := gigazine.ParseAnime(html)
	if err != nil {
		log.Fatal(err)
	}

	if sortType == "post" {
		// do nothing
	} else if sortType == "time" {
		animes = lib.SortAnimeByTime(animes)
	}

	var r string
	if outputFormat == "markdown" {
		r = format.Markdown(animes, url)
	} else if outputFormat == "json" {
		r = format.JSON(animes)
	}

	fmt.Println(r)

}
