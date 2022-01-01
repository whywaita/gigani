package gigazine

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
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
				{"スーパー・クルックス", "https://www.netflix.com/jp/title/81034907", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"ジョジョの奇妙な冒険 ストーンオーシャン", "https://jojo-portal.com/anime/so/", time.Date(2022, 1, 8, 0, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"名探偵コナン 警察学校編 Wild Police Story", "https://www.ytv.co.jp/conan/police/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"鬼滅の刃 遊郭編", "https://kimetsu.com/anime/yukakuhen/", time.Date(2021, 12, 5, 23, 15, 0, 0, time.UTC), "フジテレビ"},
				{"かなしきデブ猫ちゃん", "https://www.nhk.or.jp/matsuyama/debunekochan/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"王子の本命は悪役令嬢", "https://akuyakureijou.af-original.com/", time.Date(2022, 1, 10, 1, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"空色ユーティリティ", "https://yostar-pictures.co.jp/sorairo-utility/", time.Date(2021, 12, 31, 19, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"ロード・エルメロイII世の事件簿 -魔眼蒐集列車 Grace note- 特別編", "https://anime.elmelloi.com/", time.Date(2021, 12, 31, 22, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"魔法科高校の劣等生 追憶編", "https://mahouka.jp/", time.Date(2021, 12, 31, 23, 57, 0, 0, time.UTC), "TOKYO MX"},
				{"イロドリミドリ", "https://chunithm.sega.jp/irodorimidori/anime/", time.Date(2022, 1, 5, 1, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"リアデイルの大地にて", "https://leadale.net/", time.Date(2022, 1, 5, 23, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"ハコヅメ～交番女子の逆襲～", "https://hakozume-anime.com/", time.Date(2022, 1, 6, 1, 5, 0, 0, time.UTC), "TOKYO MX"},
				{"東京24区", "https://tokyo24project.com/", time.Date(2022, 1, 6, 0, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"オリエント", "https://orient-anime.jp/", time.Date(2022, 1, 6, 0, 0, 0, 0, time.UTC), "テレビ東京"},
				{"最遊記RELOAD -ZEROIN-", "https://saiyuki-r-zeroin.jp/", time.Date(2022, 1, 7, 0, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"ありふれた職業で世界最強 2nd season", "https://arifureta.com/", time.Date(2022, 1, 7, 0, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"スローループ", "https://slowlooptv.com/", time.Date(2022, 1, 7, 22, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"ヴァニタスの手記(第2クール・ジェヴォーダン編)", "https://vanitas-anime.com/", time.Date(2022, 1, 8, 0, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"ドールズフロントライン", "https://www.gf-anime.com/", time.Date(2022, 1, 8, 1, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"からかい上手の高木さん3", "https://takagi3.me/index.html", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"終末のハーレム", "https://end-harem-anime.com/", time.Date(2022, 1, 8, 1, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"CUE!", "https://cue-animation.jp/", time.Date(2022, 1, 8, 1, 55, 0, 0, time.UTC), "TBS"},
				{"龍青神ブルーヴ", "https://bluev.tv/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"ニンジャラ", "https://www.tv-tokyo.co.jp/anime/ninjala/", time.Date(2022, 1, 8, 10, 30, 0, 0, time.UTC), "テレビ東京系6局ネット"},
				{"失格紋の最強賢者", "https://shikkakumon.com/", time.Date(2022, 1, 8, 22, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"その着せ替え人形は恋をする", "https://bisquedoll-anime.com/", time.Date(2022, 1, 9, 0, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"明日(あけび)ちゃんのセーラー服", "https://akebi-chan.jp/", time.Date(2022, 1, 9, 0, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"現実主義勇者の王国再建記 第二部", "https://genkoku-anime.com/", time.Date(2022, 1, 9, 1, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"怪人開発部の黒井津さん", "https://kuroitsusan-anime.com/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"時光代理人 -LINK CLICK-", "https://link-click.jp/", time.Date(2022, 1, 9, 21, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"錆色のアーマ-黎明-", "https://rustedarmors-anime.com/", time.Date(2022, 1, 9, 22, 0, 0, 0, time.UTC), "TOKYO MX1"},
				{"薔薇王の葬列", "https://baraou-anime.com/", time.Date(2022, 1, 9, 22, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"フットサルボーイズ!!!!!", "https://futsalboys.com/anime", time.Date(2022, 1, 9, 23, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"佐々木と宮野", "https://sasamiya.com/", time.Date(2022, 1, 10, 0, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"進撃の巨人 The Final Season Part 2", "https://shingeki.tv/final/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"オンエアできない！", "https://www.tv-tokyo.co.jp/anime/oadekinai/", time.Date(2022, 1, 11, 2, 30, 0, 0, time.UTC), "テレビ東京"},
				{"ガル学。Ⅱ～Lucky Stars～", "https://garugaku.com/", time.Date(2022, 1, 10, 7, 5, 0, 0, time.UTC), "テレビ東京系6局ネット"},
				{"トライブナイン", "https://tribenine.tokyo/anime/", time.Date(2022, 1, 10, 22, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"賢者の弟子を名乗る賢者", "https://kendeshi-anime.com/", time.Date(2022, 1, 12, 0, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"プリンセスコネクト！Re:Dive Season 2", "https://anime.priconne-redive.jp/", time.Date(2022, 1, 11, 0, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"錆喰いビスコ", "https://sabikuibisco.jp/", time.Date(2022, 1, 11, 0, 30, 0, 0, time.UTC), "TOKYO MX"},
				{"幻想三國誌 -天元霊心記-", "https://gensou-sangokushi.com/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"お昼のショッカーさん", "https://twitter.com/Shocker_SAN1111", time.Date(2022, 1, 21, 19, 28, 0, 0, time.UTC), "TOKYO MX"},
				{"天才王子の赤字国家再生術", "https://tensaiouji-anime.com/", time.Date(2022, 1, 11, 23, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"異世界美少女受肉おじさんと", "https://fabiniku.com/", time.Date(2022, 1, 12, 0, 0, 0, 0, time.UTC), "テレビ東京"},
				{"殺し愛", "https://love-of-kill.com/", time.Date(2022, 1, 13, 0, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"平家物語", "https://heike-anime.asmik-ace.co.jp/", time.Date(2022, 1, 13, 0, 55, 0, 0, time.UTC), "フジテレビ"},
				{"ぷちセカ", "https://pjsekai.sega.jp/", time.Date(2022, 1, 14, 1, 0, 0, 0, time.UTC), "TOKYO MX"},
				{"リーマンズクラブ", "https://rymansclub.com/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"テイルズ オブ ルミナリア The Fateful Crossroad", "https://luminaria.tales-ch.jp/anime/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"地球外少年少女", "https://chikyugai.com/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
				{"テイコウペンギン", "https://www.tv-tokyo.co.jp/anime/kinder/intro/sakuhin30/", time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), "テレビ東京系6局ネット"},
				{"デリシャスパーティ♡プリキュア", "https://www.toei-anim.co.jp/tv/delicious-party_precure/", time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC), "指定されていない放送局です"},
			},
		},
	}

	for _, test := range tests {
		got, err := ParseAnime(test.input)
		if err != nil {
			t.Fatalf("ParseAnime return err: %+v", err)
		}

		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("ParseAnime() mismatch (-want +got):\n%s", diff)
		}
	}
}
