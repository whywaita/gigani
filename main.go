package main

import (
	"context"
	"flag"
	"os"

	"github.com/whywaita/gigani/cli"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&cli.GigazineCmd{}, "")
	subcommands.Register(&cli.AnnictCmd{}, "")
	subcommands.Register(&cli.SyobocalCmd{}, "")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
