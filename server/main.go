package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/whywaita/gigani/format"
	"github.com/whywaita/gigani/lib"
	"github.com/whywaita/gigani/parse"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	jsonStr := `{"health": "true"}`

	fmt.Fprintf(w, jsonStr)
}

func getAnimesJsonForServer(url string) (string, error) {
	html, err := lib.GetHTML(url)
	if err != nil {
		return "", err
	}

	animes, err := parse.ParseAnime(html)
	if err != nil {
		return "", err
	}

	r := format.JSON(animes)

	return r, nil
}

func main() {
	url := "https://gigazine.net/news/20171210-anime-2018winter/"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = ":" + port

	json, err := getAnimesJsonForServer(url)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(healthCheck)

	router.HandleFunc("/{year}/{season}", func(w http.ResponseWriter, r *http.Request) {
		// TODO: check year and season
		fmt.Fprintf(w, json)
	})

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("[ERROR]", zap.Error(err))
	}
}
