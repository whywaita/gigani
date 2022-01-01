package lib

import (
	"io"
	"net/http"
)

func GetHTML(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	html := string(body)

	return html, nil
}
