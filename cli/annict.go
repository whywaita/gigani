package cli

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/google/subcommands"

	"github.com/whywaita/gigani/lib"
	"github.com/whywaita/gigani/parse/annict"
)

type AnnictCmd struct {
	personalToken string
	format        string
	sortType      string
}

const (
	AnnictLayout = "2006"
)

func (*AnnictCmd) Name() string     { return "annict" }
func (*AnnictCmd) Synopsis() string { return "Parse Annict API" }
func (*AnnictCmd) Usage() string {
	return `gigazine -token [-format] [-sort]:
  Parse Annict API.
`
}

func (a *AnnictCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&a.personalToken, "token", "", "personal token of Annict")
	f.StringVar(&a.format, "format", "markdown", "format of print format")
	f.StringVar(&a.sortType, "sort", "post", "sort type of output")
}

func (a *AnnictCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	var err error

	if a.personalToken == "" {
		fmt.Println("token must be set")
		return subcommands.ExitFailure
	}

	if err := validateFlag(a.format, a.sortType); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	err = annict.MustWatchAllWorksInSeason(ctx, a.personalToken, getNowAnnictSeason())
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	animes, err := annict.GetAnimes(ctx, a.personalToken, getNowAnnictSeason())
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	lib.PrintAnime(animes, a.sortType, a.format, annict.AnnictGraphQLURL)
	return subcommands.ExitSuccess
}

func getNowAnnictSeason() string {
	return time.Now().Format(AnnictLayout) + "-" + lib.GetNoWSeason().String()
}
