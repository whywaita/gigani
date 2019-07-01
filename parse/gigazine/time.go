package gigazine

import "strings"

func getYear(sentence string) (year string, err error) {
	r := strings.NewReader(sentence)
	htmlTitle, _ := GetHtmlTitle(r)

	s := strings.Split(htmlTitle, "年")
	year = s[0]

	return year, nil
}
