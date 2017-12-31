package parse

import (
	"bufio"
	"encoding/xml"
	"strings"
	"time"
)

type htmlTitle struct {
	H2 Value `xml:"h2"`
	A  Value `xml:"a"`
}

type Value struct {
	Content string `xml:",innerxml"`
}

type Anime struct {
	Name        string    `json:"name"`
	URL         string    `json:"url"`
	StartDate   time.Time `json:"start_date"`
	BloadCaster string    `json:"broadcaster"`
}

func ParseAnime(html string) ([]Anime, error) {
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

	var err error

	for _, sentence := range animeBlock {
		if strings.HasPrefix(sentence, `</p><hr><p class="preface">`) {
			// `</p><hr><p class="preface">` contain title
			anime.Name, err = trimTitle(sentence)
			if err != nil {
				return Anime{}, err
			}
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
					anime.StartDate = time.Time{}
				} else {
					startData, err := parseTime(s[1])
					if err != nil {
						return Anime{}, err
					}

					anime.StartDate = startData
				}
			}
		}
	}

	if anime.BloadCaster == "" {
		anime.BloadCaster = "指定されていない放送局です"
	}

	return *anime, nil
}

func trimTitle(sentence string) (title string, err error) {
	s := strings.TrimPrefix(sentence, `</p><hr><p class="preface"></p>`)
	h := htmlTitle{}

	err = xml.NewDecoder(strings.NewReader(s)).Decode(&h)
	title = h.A.Content

	return title, nil
}

func trimURL(sentence string) (url string) {
	a := strings.TrimPrefix(sentence, `</p><hr><p class="preface"></p><h2><a href="`)
	s := strings.Split(a, `"`)
	url = s[0]

	return url
}
