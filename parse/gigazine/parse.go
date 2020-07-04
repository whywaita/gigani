package gigazine

import (
	"bufio"
	"encoding/xml"
	"strings"
	"time"

	"github.com/whywaita/gigani/parse"
)

type htmlAnimeTitle struct {
	H2 Value `xml:"h2"`
	A  Value `xml:"a"`
}

type htmlAnimeTitleWithB struct {
	H2 Value  `xml:"h2"`
	B  BValue `xml:"b"`
}

type BValue struct {
	A Value `xml:"a"`
}

type Value struct {
	Content string `xml:",innerxml"`
}

func ParseAnime(html string) ([]parse.Anime, error) {
	var err error
	reader := strings.NewReader(html)
	scanner := bufio.NewScanner(reader)

	animeBlock := []string{}
	animes := []parse.Anime{}

	year := ""

	for scanner.Scan() {
		sentence := scanner.Text()

		if strings.HasPrefix(sentence, `<title>`) {
			year, err = getYear(sentence)
			if err != nil {
				return []parse.Anime{}, err
			}
		}

		if sentence == "<b>・キャスト</b><br />" {
			// "・キャスト" is last contents of anime block
			anime, err := _parseAnime(animeBlock, year)
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

func _parseAnime(animeBlock []string, year string) (parse.Anime, error) {
	// anime block per anime
	anime := &parse.Anime{}
	var err error

L:

	for _, sentence := range animeBlock {
		if strings.HasPrefix(sentence, `</p><hr><p class="preface">`) {
			// `</p><hr><p class="preface">` contain title

			if anime.Name != "" {
				continue
			}

			anime.Name, err = trimTitle(sentence)
			if err != nil {
				return parse.Anime{}, err
			}
			anime.URL = trimURL(sentence)
		}

		for _, b := range parse.ListBroadcaster {
			if strings.HasPrefix(sentence, b) {
				// want to fast as soon as possible
				sentence = strings.TrimSuffix(sentence, "<br />")
				s := strings.Split(sentence, "：")
				anime.BloadCaster = s[0]
				if len(s) != 2 {
					// if StartDate is undefined, it is blank
					anime.StartDate = time.Time{}
				} else {
					startData, err := parseTime(s[1], year)
					if err != nil {
						return parse.Anime{}, err
					}

					anime.StartDate = startData

					break L
				}
			}
		}
	}

	if anime.BloadCaster == "" {
		anime.BloadCaster = parse.BroadCasterNotFound
	}

	return *anime, nil
}

func trimTitle(sentence string) (title string, err error) {
	s := strings.TrimPrefix(sentence, `</p><hr><p class="preface"></p>`)
	s = strings.TrimSuffix(s, `<p class="preface">`)

	if strings.Contains(s, "</b>") {
		h := htmlAnimeTitleWithB{}

		err = xml.NewDecoder(strings.NewReader(s)).Decode(&h)
		if err != nil {
			return "", err
		}

		title = h.B.A.Content
		return title, nil
	} else {
		h := htmlAnimeTitle{}

		err = xml.NewDecoder(strings.NewReader(s)).Decode(&h)
		if err != nil {
			return "", err
		}

		title = h.A.Content
		return title, nil
	}

	return title, nil
}

func trimURL(sentence string) (url string) {
	var prefix string
	if strings.Contains(sentence, "</b>") {
		prefix = `</p><hr><p class="preface"></p><h2><b><a href="`
	} else {
		prefix = `</p><hr><p class="preface"></p><h2><a href="`
	}

	a := strings.TrimPrefix(sentence, prefix)
	s := strings.Split(a, `"`)
	url = s[0]

	return url
}

func listBroadCaster() [6]string {
	// return list of broadcaster for gigazine notation
	return parse.ListBroadcaster
}
