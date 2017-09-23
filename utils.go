package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getHTML(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(body)
	html := buf.String()

	return html, nil
}
