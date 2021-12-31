package gigazine

import (
	"reflect"
	"testing"
	"time"

	"github.com/whywaita/gigani/parse"
)

func TestParseAnime(t *testing.T) {
	tests := []struct {
		input string
		want  []parse.Anime
	}{
		{
			input: sample2017AutumnHTML,
			want: []parse.Anime{
				{"血界戦線＆BEYOND", "http://kekkaisensen.com/", time.Date(2017, 10, 8, 24, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"3月のライオン 第2シリーズ", "http://3lion-anime.com/", time.Time{}, "指定されていない放送局です"},
			},
		},
		{
			input: sample2022WinterHTML,
			want: []parse.Anime{
				{"スローループ", "https://slowlooptv.com/", time.Date(2022, 1, 7, 22, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"ヴァニタスの手記(第2クール・ジェヴォーダン編)", "https://vanitas-anime.com/", time.Date(2022, 1, 8, 0, 0, 0, 0, time.UTC), "TOKYO MX"},
			},
		},
	}

	for _, test := range tests {
		got, err := ParseAnime(test.input)
		if err != nil {
			t.Fatalf("ParseAnime return err: %+v", err)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("ParseAnime want to return %+v, but got %+v", test.want, got)
		}
	}
}
