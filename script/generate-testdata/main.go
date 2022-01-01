package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/whywaita/gigani/lib"
	"github.com/whywaita/gigani/parse/gigazine"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	u := os.Args[1]

	html, err := lib.GetHTML(u)
	if err != nil {
		return fmt.Errorf("failed to get HTML: %w", err)
	}
	animes, err := gigazine.ParseAnime(html)
	if err != nil {
		return fmt.Errorf("failed to parse anime from GIGAZINE: %w", err)
	}

	for _, anime := range animes {
		fmt.Printf(`{"%s", "%s", %s, "%s"},`, anime.Name, anime.URL, formatTime(anime.StartDate), anime.BloadCaster)
		fmt.Println()
	}
	return nil
}

func formatTime(in time.Time) string {
	return fmt.Sprintf(`time.Date(%d, %d, %d, %d, %d, %d, %d, time.UTC)`,
		in.Year(),
		in.Month(),
		in.Day(),
		in.Hour(),
		in.Minute(),
		in.Second(),
		in.Nanosecond())
}
