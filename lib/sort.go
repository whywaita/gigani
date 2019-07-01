package lib

import (
	"sort"

	"github.com/whywaita/gigani/parse"
)

func SortAnimeByTime(animes []parse.Anime) []parse.Anime {
	sort.Slice(animes, func(i, j int) bool {
		return animes[i].StartDate.Before(animes[j].StartDate)
	})

	return animes
}
