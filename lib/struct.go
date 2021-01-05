package lib

import (
	"fmt"
	"time"

	"github.com/whywaita/gigani/format"
	"github.com/whywaita/gigani/parse"
)

func RemoveDuplicateTitle(animes []parse.Anime) []parse.Anime {
	registeredTime := map[string]time.Time{} // key: Title, value: start time in unixtime
	var result []parse.Anime

	for _, a := range animes {
		t, ok := registeredTime[a.Name]
		if ok && a.StartDate.After(t) {
			// if already registered than faster program in same Title, be skip
			continue
		} else if ok && a.StartDate.Before(t) {
			// found faster program, delete already record
			result = removeAnimes(result, a.Name)
		}

		result = append(result, a)
		registeredTime[a.Name] = a.StartDate
	}

	return result
}

func removeAnimes(animes []parse.Anime, title string) []parse.Anime {
	var result []parse.Anime

	for _, v := range animes {
		if v.Name != title {
			result = append(result, v)
		}
	}

	return result
}

func PrintAnime(animes []parse.Anime, sortType, outputFormat, sourceURL string) {
	if sortType == "time" {
		animes = SortAnimeByTime(animes)
	}

	var r string
	switch outputFormat {
	case "markdown":
		r = format.Markdown(animes, sourceURL)
	case "json":
		r = format.JSON(animes)
	}

	fmt.Println(r)
}
