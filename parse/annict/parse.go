package annict

import (
	"context"
	"sort"
	"time"

	"github.com/whywaita/gigani/parse"
)

func GetAnimes(ctx context.Context, token, season string) ([]parse.Anime, error) {
	var animes []parse.Anime

	resp, err := searchWorks(ctx, token, season)
	if err != nil {
		return nil, err
	}

	works, err := decodeBody(resp)
	if err != nil {
		return nil, err
	}

	for _, work := range works {
		animes = append(animes, workToAnime(work))
	}

	return animes, nil
}

func workToAnime(work AnnictWork) parse.Anime {
	anime := parse.Anime{}

	anime.Name = work.Title
	anime.URL = work.OfficialSiteUrl.String()

	// parse.Anime has first broadcaster
	programs := work.Programs
	sort.Slice(programs, func(i, j int) bool {
		return programs[i].StartedAt.Before(programs[j].StartedAt)
	})

	if len(programs) == 0 {
		anime.StartDate = time.Time{}
		anime.BloadCaster = parse.BroadCasterNotFound
		return anime
	}

	for _, p := range programs {
		for _, broadcaster := range listBroadCaster() {
			if p.Channel.Name == broadcaster {
				// found!
				anime.StartDate = p.StartedAt
				anime.BloadCaster = p.Channel.Name

				return anime
			}
		}
	}

	// if not found, return first media
	anime.StartDate = programs[0].StartedAt
	anime.BloadCaster = programs[0].Channel.Name

	return anime
}

func listBroadCaster() [6]string {
	return parse.ListBroadcaster
}
