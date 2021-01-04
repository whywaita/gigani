package syobocal

import (
	"log"
	"sync"
	"time"

	"github.com/whywaita/gigani/lib"
	"github.com/whywaita/gigani/parse"
)

var (
	titles = map[uint64]Title{}
	mu     sync.Mutex
)

func ParseAnime(data Response) ([]parse.Anime, error) {
	mu.Lock()
	for _, t := range data.Titles {
		titles[t.TID] = t
	}
	mu.Unlock()

	animes := getAnimes(data.Programs, titles)
	return lib.RemoveDuplicateTitle(animes), nil
}

func foundLink(links [][]string) string {
	if len(links) == 0 {
		return ""
	}

	for _, link := range links {
		if len(link) != 2 {
			return ""
		}

		if link[1] == "公式" {
			return link[0]
		}
	}

	return links[0][0]
}

func getAnimes(programs map[string]Program, titles map[uint64]Title) []parse.Anime {
	var animes []parse.Anime
	for _, p := range programs {
		title, ok := titles[p.TID]
		if !ok {
			log.Printf("found program, but not found title (Program: %+v)\n", p)
			continue
		}

		if !canAppendProgram(p, title) {
			continue
		}

		a := parse.Anime{
			Name:        title.Title,
			URL:         foundLink(title.Links),
			StartDate:   time.Unix(p.StTime, 0),
			BloadCaster: p.ChName,
		}
		animes = append(animes, a)
	}

	return animes
}

func canAppendProgram(p Program, title Title) bool {
	if p.Count != 1 {
		return false
	}

	if p.ChName == "AT-X" {
		return false
	}

	return true
}
