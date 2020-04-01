package annict

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/machinebox/graphql"
)

const (
	AnnictGraphQLURL    = "https://api.annict.com/graphql"
	ErrEpisodeIsNotNull = "graphql: Cannot return null for non-nullable field Program.episode"
)

type AnnictResp struct {
	SearchWorks struct {
		Edges []struct {
			Node struct {
				Title           string `json:"title"`
				OfficialSiteURL string `json:"officialSiteUrl"`
				Media           string `json:"media"`
				EpisodesCount   int    `json:"episodesCount"`
				Programs        struct {
					Nodes []AnnictProgram `json:"nodes"`
				} `json:"programs"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"searchWorks"`
}

type AnnictWork struct {
	Title           string
	OfficialSiteUrl url.URL
	Media           string
	Programs        []AnnictProgram
}

type AnnictProgram struct {
	StartedAt time.Time
	Channel   AnnictChannel `json:"channel"`
}

type AnnictChannel struct {
	Name string `json:"name"`
}

func searchWorks(ctx context.Context, personalToken, season string) (*AnnictResp, error) {
	client := graphql.NewClient(AnnictGraphQLURL)
	q := fmt.Sprintf(`
query{
  searchWorks(seasons: ["%s"]) {
    edges {
      node {
        title
        officialSiteUrl
        media
        episodesCount
        programs {
          nodes {
            startedAt
            channel {
              name
            }
            episode {
              number
            }
          }
        }
      }
    }
  }
}
`, season)

	req := graphql.NewRequest(q)
	req.Header.Set("Authorization", "Bearer "+personalToken)

	var resp AnnictResp
	err := client.Run(ctx, req, &resp)
	if err.Error() == ErrEpisodeIsNotNull {
		// allow null in episode
		return &resp, nil
	}

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func decodeBody(resp *AnnictResp) ([]AnnictWork, error) {
	var works []AnnictWork

	for _, v := range resp.SearchWorks.Edges {
		if v.Node.Media != "TV" {
			continue
		}
		title := v.Node.Title
		media := v.Node.Media

		u, err := url.Parse(v.Node.OfficialSiteURL)
		if err != nil {
			return nil, err
		}

		var ps []AnnictProgram

		for _, p := range v.Node.Programs.Nodes {
			if p.Channel.Name == "" {
				// program is null
				continue
			}
			ps = append(ps, p)
		}

		work := AnnictWork{
			Title:           title,
			OfficialSiteUrl: *u,
			Media:           media,
			Programs:        ps,
		}

		works = append(works, work)
	}

	return works, nil
}
