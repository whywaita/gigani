package cli

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/subcommands"

	"github.com/whywaita/gigani/lib"
	"github.com/whywaita/gigani/parse/gigazine"
)

type GigazineCmd struct {
	url      string
	format   string
	sortType string
}

func (*GigazineCmd) Name() string     { return "gigazine" }
func (*GigazineCmd) Synopsis() string { return "Parse gigazine page" }
func (*GigazineCmd) Usage() string {
	return `gigazine -url [-format] [-sort]:
  Parse gigazine page.
`
}

func (g *GigazineCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&g.url, "url", "", "URL of GIGAZINE")
	f.StringVar(&g.format, "format", "markdown", "format of print format")
	f.StringVar(&g.sortType, "sort", "post", "sort type of output")
}

func (g *GigazineCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	err := validateFlagURL(g.url)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if err := validateFlag(g.format, g.sortType); err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	html, err := lib.GetHTML(g.url)
	if err != nil {
		log.Fatal(err)
	}

	animes, err := gigazine.ParseAnime(html)
	if err != nil {
		log.Fatal(err)
	}

	lib.PrintAnime(animes, g.sortType, g.format, g.url)
	return subcommands.ExitSuccess
}
