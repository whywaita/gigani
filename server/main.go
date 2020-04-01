package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/whywaita/gigani/parse/gigazine"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	"github.com/whywaita/gigani/format"
	"github.com/whywaita/gigani/lib"
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

	animes, err := gigazine.ParseAnime(html)
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
	port = strings.Join([]string{":", port}, "")

	json, err := getAnimesJsonForServer(url)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(healthCheck)

	router.HandleFunc("/{year}/{season}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if strings.EqualFold(vars["year"], "2018") && strings.EqualFold(vars["season"], "winter") {
			fmt.Fprintf(w, json)
		} else {
			http.NotFound(w, r)
		}
	})

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("[ERROR]", zap.Error(err))
	}
}
