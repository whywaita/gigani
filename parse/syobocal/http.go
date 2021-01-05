package syobocal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/whywaita/gigani/lib"
)

var (
	APIEndpoint, _ = url.Parse("https://cal.syoboi.jp/json.php")
)

// Response is response of
type Response struct {
	Programs map[string]Program `json:"programs"`
	Titles   map[string]Title   `json:"titles"`
}

// Program is information of broadcast
type Program struct {
	PID         string      `json:"PID"`
	TID         uint64      `json:"TID,string"`
	ChID        string      `json:"ChID"`
	ChName      string      `json:"ChName"`
	ChEPGURL    string      `json:"ChEPGURL"`
	Count       int         `json:"Count,string"`
	StTime      int64       `json:"StTime,string"`
	EdTime      string      `json:"EdTime"`
	SubTitle2   string      `json:"SubTitle2"`
	ProgComment string      `json:"ProgComment"`
	ConfFlag    interface{} `json:"ConfFlag"`
}

// Title is information of animation
type Title struct {
	TID           uint64      `json:"TID,string"`
	Title         string      `json:"Title"`
	ShortTitle    string      `json:"ShortTitle"`
	TitleYomi     string      `json:"TitleYomi"`
	TitleEN       string      `json:"TitleEN"`
	Cat           string      `json:"Cat"`
	FirstCh       string      `json:"FirstCh"`
	FirstYear     string      `json:"FirstYear"`
	FirstMonth    string      `json:"FirstMonth"`
	FirstEndYear  interface{} `json:"FirstEndYear"`
	FirstEndMonth interface{} `json:"FirstEndMonth"`
	TitleFlag     string      `json:"TitleFlag"`
	Links         [][]string  `json:"Links"`
}

func GetData(ctx context.Context, year int, season lib.Season) (*Response, error) {
	var r Response
	months := season.Months()

	for _, month := range months {
		ps, ts, err := getData(ctx, year, month)
		if err != nil {
			return nil, fmt.Errorf("failed to get monthly data: %w", err)
		}

		r.Programs = mergeProgram(r.Programs, ps)
		r.Titles = mergeTitle(r.Titles, ts)
	}

	return &r, nil
}

func getData(ctx context.Context, year, month int) (map[string]Program, map[string]Title, error) {
	// get year/month/01 ~ last day of year/month
	startTime := fmt.Sprintf("%d-%02d-01", year, month)

	q := APIEndpoint.Query()
	q.Set("Req", "ProgramByDate,TitleMedium")
	q.Set("Start", startTime)
	q.Set("days", "32")
	APIEndpoint.RawQuery = q.Encode()

	resp, err := http.Get(APIEndpoint.String())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to GET request: %w", err)
	}
	defer resp.Body.Close()

	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return r.Programs, r.Titles, nil
}

func mergeProgram(m ...map[string]Program) map[string]Program {
	ans := make(map[string]Program, 0)

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}

func mergeTitle(m ...map[string]Title) map[string]Title {
	ans := make(map[string]Title, 0)

	for _, c := range m {
		for k, v := range c {
			ans[k] = v
		}
	}
	return ans
}
