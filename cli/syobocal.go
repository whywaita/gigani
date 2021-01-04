package cli

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/whywaita/gigani/lib"

	"github.com/whywaita/gigani/parse/syobocal"

	"github.com/google/subcommands"
)

type SyobocalCmd struct {
	format   string
	sortType string
}

func (*SyobocalCmd) Name() string     { return "syobocal" }
func (*SyobocalCmd) Synopsis() string { return "Parse syobocal page" }
func (*SyobocalCmd) Usage() string {
	return `syobocal -url [-format] [-sort]:
  parse from syobocal JSON API.
`
}

func (s *SyobocalCmd) SetFlags(f *flag.FlagSet) {
	//f.StringVar(&s.url, "url", "", "URL of syobocal")
	f.StringVar(&s.format, "format", "markdown", "format of print format")
	f.StringVar(&s.sortType, "sort", "post", "sort type of output")
}

func (s *SyobocalCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := validateFlagOutputFormat(s.format); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if err := validateFlagSort(s.sortType); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	ctx := context.Background()
	now := time.Now()
	season := lib.GetNoWSeason()

	data, err := syobocal.GetData(ctx, now.Year(), season)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	animes, err := syobocal.ParseAnime(*data)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	lib.PrintAnime(animes, s.sortType, s.format, syobocal.APIEndpoint.String())
	return subcommands.ExitSuccess
}
