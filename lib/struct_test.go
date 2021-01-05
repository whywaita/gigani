package lib

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/whywaita/gigani/parse"
)

func TestRemoveDuplicateTitle(t *testing.T) {
	tests := []struct {
		input []parse.Anime
		want  []parse.Anime
		err   bool
	}{
		{
			input: []parse.Anime{
				{
					Name:        "転生したらスライムだった件 第2期 第1部",
					URL:         "https://www.ten-sura.com/anime/tensura",
					StartDate:   time.Date(2021, 1, 5, 23, 00, 00, 0, time.Local),
					BloadCaster: "TOKYO MX",
				},
			},
			want: []parse.Anime{
				{
					Name:        "転生したらスライムだった件 第2期 第1部",
					URL:         "https://www.ten-sura.com/anime/tensura",
					StartDate:   time.Date(2021, 1, 5, 23, 00, 00, 0, time.Local),
					BloadCaster: "TOKYO MX",
				},
			},
			err: false,
		},
		{
			input: []parse.Anime{
				{
					Name:        "呪術廻戦",
					URL:         "https://jujutsukaisen.jp/",
					StartDate:   time.Date(2021, 1, 9, 1, 55, 00, 0, time.Local),
					BloadCaster: "TBS",
				},
				{
					Name:        "呪術廻戦",
					URL:         "https://jujutsukaisen.jp/",
					StartDate:   time.Date(2021, 1, 9, 1, 56, 00, 0, time.Local),
					BloadCaster: "MBS毎日放送",
				},
			},
			want: []parse.Anime{
				{
					Name:        "呪術廻戦",
					URL:         "https://jujutsukaisen.jp/",
					StartDate:   time.Date(2021, 1, 9, 1, 55, 00, 0, time.Local),
					BloadCaster: "TBS",
				},
			},
			err: false,
		},
	}

	for _, test := range tests {
		got := RemoveDuplicateTitle(test.input)

		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("mismatch (-want +got):\n%s", diff)
		}
	}
}
