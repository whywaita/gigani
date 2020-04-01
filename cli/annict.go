package cli

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/whywaita/gigani/parse/annict"

	"github.com/whywaita/gigani/format"

	"github.com/google/subcommands"
	"github.com/whywaita/gigani/lib"
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

	err = validateFlagOutputFormat(a.format)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	err = validateFlagSort(a.sortType)
	if err != nil {
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

	if a.sortType == "time" {
		animes = lib.SortAnimeByTime(animes)
	}

	var r string
	if a.format == "markdown" {
		r = format.Markdown(animes, "Annict")
	} else if a.format == "json" {
		r = format.JSON(animes)
	}

	fmt.Println(r)

	return subcommands.ExitSuccess
}

func getNowAnnictSeason() string {
	return time.Now().Format(AnnictLayout) + "-" + lib.GetNoWSeason().String()
}
