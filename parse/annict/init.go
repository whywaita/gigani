package annict

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/machinebox/graphql"
)

type WorksResp struct {
	SearchWorks struct {
		Edges []struct {
			Node struct {
				ID                string `json:"id"`
				Title             string `json:"title"`
				ViewerStatusState string `json:"viewerStatusState"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"searchWorks"`
}

func MustWatchAllWorksInSeason(ctx context.Context, personalToken, season string) error {
	q := fmt.Sprintf(`query {
  searchWorks(seasons: ["%s"]) {
    edges {
      node {
        id
        title
        viewerStatusState
      }
    }
  }
}
`, season)

	client := graphql.NewClient(AnnictGraphQLURL)
	allWorksReq := graphql.NewRequest(q)
	allWorksReq.Header.Set("Authorization", "Bearer "+personalToken)
	var wr WorksResp
	err := client.Run(ctx, allWorksReq, &wr)
	if err != nil {
		return errors.Wrap(err, "failed to get all works")
	}

	client.Log = func(s string) { fmt.Println(s) }
	for _, node := range wr.SearchWorks.Edges {
		state := "WANNA_WATCH"
		if node.Node.ViewerStatusState == state {
			continue
		}

		workId := node.Node.ID
		m := fmt.Sprintf(`mutation {
  updateStatus(input: { state: %s, workId: "%s"} ) {
    clientMutationId
  }
}
`, state, workId)

		req := graphql.NewRequest(m)
		req.Header.Set("Authorization", "Bearer "+personalToken)

		var resp interface{}
		err = client.Run(ctx, req, &resp)
		if err != nil {
			return errors.Wrap(err, "failed to update state")
		}
		fmt.Printf("updated to %s (%s)\n", state, node.Node.Title)
	}

	return nil
}
