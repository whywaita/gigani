package parse

import (
	"sort"
)

func SortAnimeByTime(animes []Anime) []Anime {
	sort.Slice(animes, func(i, j int) bool {
		return animes[i].StartDate.Before(animes[j].StartDate)
	})

	return animes
}
