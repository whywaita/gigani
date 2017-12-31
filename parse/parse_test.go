package parse

import (
	"reflect"
	"testing"
	"time"
)

func TestParseAnime(t *testing.T) {
	// sampleHTML cited 2017年秋季開始の新作アニメ一覧 - GIGAZINE (http://gigazine.net/news/20170917-anime-2017autumn/)
	sampleHTML := `
</p><hr><p class="preface"></p><h2><a href="http://kekkaisensen.com/" target="_blank">血界戦線＆BEYOND</a></h2><p class="preface"><a href="http://kekkaisensen.com/"></p><img data-src="http://i.gzn.jp/img/2017/09/17/anime-2017autumn/31.png" border="0" class="lazyload"><p class="preface"></a><br />
<br />
<b>・放送情報</b><br />
MBS：10/7(土) 27:08～<br />
TOKYO MX：10/8(日) 24:30～<br />
BS11：10/8(日) 24:30～<br />
AT-X：10/12(木) 20:00～ほか<br />
<br />
<b>・作品情報</b><br />
一晩で消失した紐育の代わりに出現した街「ヘルサレムズ・ロット」で、世界の均衡を守るべく人知れず活動を続ける集団「秘密結社ライブラ」の活躍を描く、内藤泰弘原作のアニメ第2期。第1期は2015年春に放送され、最終回のみ2015年10月に放送された。スタッフは監督が高柳滋仁に、脚本が加茂靖子に変更となっている。これは「だがしかし」のコンビ。<br />
<br />
<b>・スタッフ</b><br />
原作：内藤泰弘<br />
監督：高柳滋仁<br />
脚本：加茂靖子<br />
キャラクターデザイン：川本利浩<br />
クリーチャーデザイン：杉浦幸次<br />
プロップデザイン：神宮司訓之<br />
エフェクト作画監督：橋本敬史<br />
美術監督：東潤一<br />
色彩設定：中山しほ子<br />
撮影監督：古本真由子<br />
CGIディレクター：太田光希<br />
編集：平木大輔<br />
音響監督：明田川仁<br />
音響効果：今野康之<br />
音楽：岩崎太整<br />
アニメーション制作：ボンズ<br />
制作：血界戦線制作委員会<br />
コピーライト表記：©2017 内藤泰弘／集英社・血界戦線 & BEYOND製作委員会<br />
<br />
Twitter：<b><a href="https://twitter.com/kekkaisensen" target="_blank">@kekkaisensen</a></b><br />
ハッシュタグ：#kekkai_anime<br />
<br />
OP：UNISON SQUARE GARDEN「fake town baby」<br />
ED：DAOKO×岡村靖幸「ステップアップLOVE」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=aXznlDNTKp0" target="_blank">『血界戦線 & BEYOND』PV - YouTube</a></b><br />
</p><div class="iframe-content"><iframe class="lazyload" data-src="https://www.youtube.com/embed/aXznlDNTKp0" frameborder="0" allowfullscreen></iframe></div><p class="preface">
<br />
<b>・キャスト</b><br />
クラウス・V・ラインヘルツ：小山力也<br />
レオナルド・ウォッチ：阪口大助<br />
ザップ・レンフロ：中井和哉<br />
スティーブン・A・スターフェイズ：宮本充<br />
チェイン・皇：小林ゆう<br />
ツェッド・オブライエン：緑川光<br />
K・K：折笠愛<br />
ギルベルト・Ｆ・アルトシュタイン：銀河万丈<br />
ソニック：内田雄馬<br />
ミシェーラ：水樹奈々<br />
<br />
</p><hr><p class="preface"></p><h2><a href="http://3lion-anime.com/" target="_blank">3月のライオン 第2シリーズ</a></h2><p class="preface"><a href="http://3lion-anime.com/"></p><img data-src="http://i.gzn.jp/img/2017/09/17/anime-2017autumn/28.png" border="0" class="lazyload"><p class="preface"></a><br />
<br />
<b>・放送情報</b><br />
NHK総合：10/14(土) 23:00～<br />
全22回<br />
<br />
<b>・作品情報</b><br />
幼いころに家族を失った棋士・桐山零の孤独な日々は、あかり・ひなた・モモという3姉妹と出会ったことから少しずつ変わりはじめていく。<br />
<br />
ヤングアニマルでやや不定期に連載中の人気作が原作。原作者の羽海野チカは新房昭之監督のファンで、アニメ化発表時に「『新房監督でシャフトさんで！』。この夢がかえられないのならアニメ化はできなくてもいい…。そう思っていました」とコメントしていた。新房監督は「物語」シリーズや「魔法少女まどか☆マギカ」「ひだまりスケッチ」などで知られている。<br />
<br />
<b>・スタッフ</b><br />
原作：羽海野チカ<br />
監督：新房昭之<br />
アニメーション制作：シャフト<br />
キャラクターデザイン：杉山延寛<br />
美術監督：田村せいき<br />
美術設定：名倉靖博<br />
音響監督：亀山俊樹<br />
音楽：橋本由香利<br />
製作：「３月のライオン」アニメ製作委員会<br />
コピーライト表記：©羽海野チカ・白泉社／「３月のライオン」アニメ製作委員会<br />
<br />
Twitter：<b><a href="https://twitter.com/3lion_anime" target="_blank">@3lion_anime</a></b><br />
ハッシュタグ：<br />
<br />
OP：YUKI「フラッグを立てろ」<br />
ED：Brian the Sun「カフネ」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=HSC5e_pEeQc" target="_blank">TVアニメ「３月のライオン」第2シリーズティザーPV | ＮＨＫ総合テレビにて10/14(土)より毎週土曜23:00放送 - YouTube</a></b><br />
</p><div class="iframe-content"><iframe class="lazyload" data-src="https://www.youtube.com/embed/HSC5e_pEeQc" frameborder="0" allowfullscreen></iframe></div><p class="preface">
<br />
<b>・キャスト</b><br />
桐山零：河西健吾<br />
川本あかり：茅野愛衣<br />
川本ひなた：花澤香菜<br />
川本モモ：久野美咲<br />
二海堂晴信：岡本信彦<br />
幸田香子：井上麻里奈<br />
高橋勇介：細谷佳正<br />
島田開：三木眞一郎<br />
三角龍雪：杉田智和<br />
松本一砂：木村昴<br />
川本相米二：千葉繁<br />
幸田柾近：大川透<br />
林田高志：櫻井孝宏<br />
花岡：上田燿司<br />
美咲：根谷美智子<br />
宗谷冬司：石田彰<br />
神宮寺崇徳：玄田哲章<br />
横溝億泰：阪口大助<br />
辻井武史：中村悠一<br />
後藤正宗：東地宏樹<br />
藤本雷堂：大塚明夫<br />
安井学：岩田光央<br />
松永正一：岡和男<br />
重田盛夫：津田健次<br />
野口英作：うえだゆうじ<br />
<br />
`
	parsedData, err := ParseAnime(sampleHTML)
	if err != nil {
		t.Fatal(err)
	}

	if len(parsedData) != 2 {
		t.Fatal("animes length must be 2")
	}

	t1 := time.Date(2017, 10, 8, 24, 30, 0, 0, time.UTC)
	anime1 := Anime{"血界戦線＆BEYOND", "http://kekkaisensen.com/", t1, "TOKYO MX"}
	anime2 := Anime{"3月のライオン 第2シリーズ", "http://3lion-anime.com/", time.Time{}, "指定されていない放送局です"}

	correctData := []Anime{}
	correctData = append(parsedData, anime1)
	correctData = append(parsedData, anime2)

	if reflect.DeepEqual(parsedData, correctData) {
		t.Fatal("missing parseAnime")
	}
}
