package main

import (
	"bufio"
	"strings"
)

func parseAnime(html string) ([]Anime, error) {
	reader := strings.NewReader(html)
	scanner := bufio.NewScanner(reader)

	animeBlock := []string{}
	animes := []Anime{}

	for scanner.Scan() {
		sentence := scanner.Text()

		if sentence == "<b>・キャスト</b><br />" {
			// "・キャスト" is last contents of anime block
			anime, err := _parseAnime(animeBlock)
			if err != nil {
				return nil, err
			}
			animes = append(animes, anime)
			animeBlock = []string{}
		}

		animeBlock = append(animeBlock, sentence) // collect sentence
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return animes, nil
}

func _parseAnime(animeBlock []string) (Anime, error) {
	// anime block per anime
	anime := &Anime{}
	// list of Broadcaster (priority order)
	list_broadcaster := [6]string{"TOKYO MX", "テレビ東京", "フジテレビ", "日本テレビ", "tvk", "TBS"}

	for _, sentence := range animeBlock {
		if strings.HasPrefix(sentence, `</p><hr><p class="preface">`) {
			// `</p><hr><p class="preface">` contain title
			anime.Name = trimTitle(sentence)
			anime.URL = trimURL(sentence)
		}

		for _, b := range list_broadcaster {
			if strings.HasPrefix(sentence, b) {
				// want to fast as soon as possible
				sentence = strings.TrimSuffix(sentence, "<br />")
				s := strings.Split(sentence, "：")
				anime.BloadCaster = s[0]
				if len(s) != 2 {
					// if StartDate is undefined, it is blank
					anime.StartDate = ""
				} else {
					anime.StartDate = s[1]
				}
			}
		}
	}

	if anime.BloadCaster == "" {
		anime.BloadCaster = "指定されていない放送局です"
	}

	return *anime, nil
}

func trimTitle(sentence string) (title string) {
	a := strings.Split(sentence, `target="_blank">`)[1]
	title = strings.Split(a, `</a>`)[0]

	return title
}

func trimURL(sentence string) (url string) {
	a := strings.TrimPrefix(sentence, `</p><hr><p class="preface"></p><h2><a href="`)
	s := strings.Split(a, `"`)
	url = s[0]

	return url
}
