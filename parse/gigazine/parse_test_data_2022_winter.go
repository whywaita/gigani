package gigazine

// sample2022WinterHTML cited 2022年冬開始の新作アニメ一覧 - GIGAZINE (https://gigazine.net/news/20211215-anime-2022winter/)
const sample2022WinterHTML = `
<!DOCTYPE html>
<link rel="dns-prefetch" href="https://i.gzn.jp">
<link rel="dns-prefetch" href="https://gigazine.be">
<html lang="ja">
<head prefix="og: https://ogp.me/ns# fb: https://ogp.me/ns/fb# article: https://ogp.me/ns/article#">
<meta charset="utf-8">
<title>2022年冬開始の新作アニメ一覧 - GIGAZINE</title>

<h1 class="title">2022年冬開始の新作アニメ一覧</h1>
<!-- google_ad_section_start -->
<p class="preface"></p><a href="https://i.gzn.jp/img/2021/12/15/anime-2022winter/00.png" target="_blank"><img src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/00_m.png" border="0" class="lzsmall img-standard-size"></a><p class="preface">
<br />
年4回の番組改編期が近づき、また多くの新作アニメが始まります。2022年冬(新春)の改編でスタートするアニメの本数は<b>50本強</b>。<br />
<br />
Netflixは改編期に縛られないこともあって、すでに注目作「ジョジョの奇妙な冒険 ストーンオーシャン」は第1話～第12話が地上波放送に先駆けて一括配信されています。このほか、原作モノでは「鬼滅の刃 遊郭編」「進撃の巨人 The Final Season Part 2」などの人気作が放送・配信されます。オリジナルアニメでは「空色ユーティリティ」「東京24区」「地球外少年少女」「永遠の831」「リーマンズクラブ」などの放送・配信が予定されています。<br />
<br />
以下、作品リストは放送・配信時期が早いものから順に並べています。作品名の頭に「◆」をつけているので、「◆」でページ内検索すると1作品ごとにどんどんジャンプしていけます。また、目次からも各作品部分へ移動可能です。<br />
<br />
放送情報欄は上が放送、下が配信。配信は独占配信や最も早いタイミングで配信を行うサービスのみ記載し、複数サービスで同時に配信がスタートするケースは、その中で早いタイミングを「他配信」としています。また、基本的に見放題サービスの配信タイミングで掲載しており、都度課金配信は異なる配信時期が設定されていることがあります。<br />
<br />
なお、2021年末放送予定だった「<b><a href="https://anime.magireco.com/" target="_blank">マギアレコード 魔法少女まどか☆マギカ外伝 Final SEASON-浅き夢の暁-</a></b>」は<b><a href="https://anime.magireco.com/news/?id=59392" target="_blank">制作上の都合のため</a></b>放送時期が2022年春に変更されています。<br />
<br />
<b id="mokuji_title">▼目次</b><span id="mokuji_button" style="color:#1020d0; cursor: pointer; font-size: 0.8em; margin-left: 4px;">表示</span><br />
<div id="mokuji" style="display:none;"><b>・<a href="#01">スーパー・クルックス</a></b><br />
<b>・<a href="#02">ジョジョの奇妙な冒険 ストーンオーシャン</a></b><br />
<b>・<a href="#03">名探偵コナン 警察学校編 Wild Police Story</a></b><br />
<b>・<a href="#04">鬼滅の刃 遊郭編</a></b><br />
<b>・<a href="#05">かなしきデブ猫ちゃん</a></b><br />
<b>・<a href="#27">王子の本命は悪役令嬢</a></b><br />
<b>・<a href="#06">空色ユーティリティ</a></b><br />
<b>・<a href="#40">ロード・エルメロイII世の事件簿 -魔眼蒐集列車 Grace note- 特別編</a></b><br />
<b>・<a href="#07">魔法科高校の劣等生 追憶編</a></b><br />
<b>・<a href="#51">イロドリミドリ</a></b><br />
<b>・<a href="#08">リアデイルの大地にて</a></b><br />
<b>・<a href="#09">ハコヅメ～交番女子の逆襲～</a></b><br />
<b>・<a href="#10">東京24区</a></b><br />
<b>・<a href="#11">オリエント</a></b><br />
<b>・<a href="#12">最遊記RELOAD -ZEROIN-</a></b><br />
<b>・<a href="#13">ありふれた職業で世界最強 2nd season</a></b><br />
<b>・<a href="#14">スローループ</a></b><br />
<b>・<a href="#15">ヴァニタスの手記(第2クール・ジェヴォーダン編)</a></b><br />
<b>・<a href="#43">ドールズフロントライン</a></b><br />
<b>・<a href="#16">からかい上手の高木さん3</a></b><br />
<b>・<a href="#17">終末のハーレム</a></b><br />
<b>・<a href="#18">CUE!</a></b><br />
<b>・<a href="#56">龍青神ブルーヴ</a></b><br />
<b>・<a href="#42">ニンジャラ</a></b><br />
<b>・<a href="#44">失格紋の最強賢者</a></b><br />
<b>・<a href="#19">その着せ替え人形は恋をする</a></b><br />
<b>・<a href="#20">明日(あけび)ちゃんのセーラー服</a></b><br />
<b>・<a href="#21">現実主義勇者の王国再建記 第二部</a></b><br />
<b>・<a href="#22">怪人開発部の黒井津さん</a></b><br />
<b>・<a href="#45">時光代理人 -LINK CLICK-</a></b><br />
<b>・<a href="#23">錆色のアーマ-黎明-</a></b><br />
<b>・<a href="#24">薔薇王の葬列</a></b><br />
<b>・<a href="#41">フットサルボーイズ!!!!!</a></b><br />
<b>・<a href="#25">佐々木と宮野</a></b><br />
<b>・<a href="#26">進撃の巨人 The Final Season Part 2</a></b><br />
<b>・<a href="#48">オンエアできない！</a></b><br />
<b>・<a href="#47">ガル学。Ⅱ～Lucky Stars～</a></b><br />
<b>・<a href="#28">トライブナイン</a></b><br />
<b>・<a href="#29">賢者の弟子を名乗る賢者</a></b><br />
<b>・<a href="#30">プリンセスコネクト！Re:Dive Season 2</a></b><br />
<b>・<a href="#31">錆喰いビスコ</a></b><br />
<b>・<a href="#32">幻想三國誌 -天元霊心記-</a></b><br />
<b>・<a href="#50">お昼のショッカーさん</a></b><br />
<b>・<a href="#33">天才王子の赤字国家再生術</a></b><br />
<b>・<a href="#39">異世界美少女受肉おじさんと</a></b><br />
<b>・<a href="#46">殺し愛</a></b><br />
<b>・<a href="#34">平家物語</a></b><br />
<b>・<a href="#54">ぷちセカ</a></b><br />
<b>・<a href="#38">あたしゃ川尻こだまだよ～デンジャラスライフハッカーのただれた生活～</a></b><br />
<b>・<a href="#37">リーマンズクラブ</a></b><br />
<b>・<a href="#53">テイルズ オブ ルミナリア The Fateful Crossroad</a></b><br />
<b>・<a href="#35">地球外少年少女</a></b><br />
<b>・<a href="#36">永遠の831</a></b><br />
<b>・<a href="#49">テイコウペンギン</a></b><br />
<b>・<a href="#55">闇芝居(第10期)</a></b><br />
<b>・<a href="#52">デリシャスパーティ♡プリキュア</a></b><br />
<b>・<a href="#57">暴太郎戦隊ドンブラザーズ</a></b><br />
</div></p><script>document.getElementById("mokuji_button").addEventListener("click",()=>{const button = document.getElementById("mokuji_button");const mokuji = document.getElementById("mokuji");const title = document.getElementById("mokuji_title");if(button.innerHTML === "表示"){mokuji.style.display = "block"; button.innerHTML = "非表示"; title.innerHTML = "◆目次"; }else {mokuji.style.display = "none"; button.innerHTML = "表示"; title.innerHTML = "▼目次"; }})</script><p class="preface">
</p><hr></p><p class="preface"><h2 id="01">◆<b><a href="https://www.netflix.com/jp/title/81034907" target="_blank">スーパー・クルックス</a></b></h2><p class="preface">
</p><a href="https://www.netflix.com/jp/title/81034907" target="_blank"><img src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/01.png" border="0" class="lzsmall"></a><p class="preface">
<br />
<b>・放送情報</b><br />
Netflix：11/25(木)～<br />
全13話<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>ツキに見放されっぱなしの強盗団に、運命の女神はほほ笑むのか？<br />
ケチな悪党ジョニー・ボルトが、スーパーパワーを持つ犯罪者仲間を集め、イチかバチか最後の強盗に打って出る！ターゲットは、これまたスーパーパワーを持つ裏の世界の大物。これで役者は揃った！</blockquote><p class="preface">

<div id='Google_IL' style='padding-top: 8px; padding-bottom: 8px;'>
<div id='div-gpt-ad-1546928741451-0'>
</p><script>
googletag.cmd.push(function() {googletag.display('div-gpt-ad-1546928741451-0');});
</script><p class="preface">
</div>
</div><p class="preface"><br />
原作はマーク・ミラーによるグラフィックノベル。他のマーク・ミラー作品としては映画化された「<b><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B016LTQVYW/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B016LTQVYW/gigazine-22');">キングスマン</a></b>」「<b><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B004IEAZ18/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B004IEAZ18/gigazine-22');">キック・アス</a></b>」「<b><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B006QJT0BO/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B006QJT0BO/gigazine-22');">ウォンテッド</a></b>」などがある。<br />
<br />
<b>・スタッフ</b><br />
原作：マーク・ミラー、レイニル・ユー<br />
原案：マーク・ミラー<br />
監督：堀元宣<br />
シリーズ構成：佐藤大<br />
キャラクター原案：レイニル・ユー<br />
キャラクターデザイン：三谷高史<br />
コンセプトデザイン：ブリュネ・スタニスラス<br />
音楽：TOWA TEI<br />
アニメーション制作：ボンズ<br />
<br />
<b>・キャスト</b><br />
ジョニー・ボルト：津田健次郎<br />
ケイシー・アン：坂本真綾<br />
クリストファー・マッツ：家中宏<br />
カーマイン/ザ・ヒート：木村靖司<br />
グラディエーター：ピエール瀧<br />
ジョシュ/ザ・ゴースト：諏訪部順一<br />
サラマンダー：江川央生<br />
TK・マッケイブ：竹本英史<br />
ロディ・ディーゼル：稲田徹<br />
サミー・ディーゼル：木村昴<br />
プレトリアン：羽多野渉<br />
フォアキャスト：KENN<br />
<br />
<b><a href="https://www.youtube.com/watch?v=ar8YX63af8o" target="_blank">『スーパー・クルックス』予告編 - Netflix</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/ar8YX63af8o" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="02">◆<b><a href="https://jojo-portal.com/anime/so/" target="_blank">ジョジョの奇妙な冒険 ストーンオーシャン</a></b></h2><p class="preface">
</p><a href="https://jojo-portal.com/anime/so/" target="_blank"><img src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/02.png" border="0" class="lzsmall"></a><p class="preface">
<br />
<b>・放送情報</b><br />
Netflix：12/1(水)～【第1話～第12話全世界独占配信】<br />
<br />
TOKYO MX：1/7(金) 24:30～<br />
BS11：1/7(金) 24:30～<br />
MBS：1/7(金) 26:55～<br />
アニマックス：1/22(土) 20:00～<br />
<br />
<b>・作品情報</b><br />
恋人の裏切りにより罪を着せられ15年の刑を受けた空条徐倫は、州立グリーン・ドルフィン・ストリート重警備刑務所に収容される。父・空条承太郎から差し入れられたペンダントによりスタンド能力に目覚めた徐倫は、刑務所という「石作りの海」から自由になることはできるのか。<br />
<br />
原作は「ジョジョの奇妙な冒険」シリーズの第6部で、1999年～2003年にかけて「週刊少年ジャンプ」に連載された。全17巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4088728661/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4088728661/gigazine-22');"><img src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/02-02.png" border="0" class="lzsmall"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：荒木飛呂彦<br />
総監督：鈴木健一<br />
監督：加藤敏幸<br />
シリーズ構成：小林靖子<br />
キャラクターデザイン：筱雅律<br />
サブキャラクターデザイン：土屋圭<br />
スタンドデザイン：石本峻一<br />
プロップデザイン：新妻大輔、宝谷幸稔<br />
美術設定：滝れーき、長澤順子、渡邊由里子<br />
美術監督：渡辺佳人<br />
色彩設計：佐藤裕子<br />
撮影監督：山田和弘<br />
編集：廣瀬清志<br />
CGプロデューサー：濱中裕<br />
CGディレクター：宍戸光太郎<br />
音響監督：岩浪美和<br />
音楽：菅野祐悟<br />
アニメーション制作：david production<br />
コピーライト表記：©LUCKY LAND COMMUNICATIONS/集英社・ジョジョの奇妙な冒険SO製作委員会<br />
<br />
<b>・キャスト</b><br />
空条徐倫：ファイルーズあい<br />
エルメェス・コステロ：田村睦心<br />
フー・ファイターズ：伊瀬茉莉也<br />
エンポリオ・アルニーニョ：種崎敦美<br />
ウェザー・リポート：梅原裕一郎<br />
ナルシソ・アナスイ：浪川大輔<br />
空条承太郎：小野大輔<br />
<br />
Twitter：<a href="https://twitter.com/anime_jojo" target="_blank">@anime_jojo</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/jojo_anime" target="_blank">#jojo_anime</a><br />
<br />
OP：ichigo from 岸田教団&THE 明星ロケッツ「STONE OCEAN」<br />
ED：Duffy「Distant Dreamer」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=yIk-OUT6ZlA" target="_blank">アニメ「ジョジョの奇妙な冒険 ストーンオーシャン」PV ～空条徐倫～</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/yIk-OUT6ZlA" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
<b><a href="https://www.youtube.com/watch?v=mgxDyrEnnoE" target="_blank">アニメ「ジョジョの奇妙な冒険 ストーンオーシャン」OP映像</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/mgxDyrEnnoE" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="03">◆<b><a href="https://www.ytv.co.jp/conan/police/" target="_blank">名探偵コナン 警察学校編 Wild Police Story</a></b></h2><p class="preface">
</p><a href="https://www.ytv.co.jp/conan/police/" target="_blank"><img src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/03.png" border="0" class="lzsmall"></a><p class="preface">
<br />
<b>・放送情報</b><br />
読売テレビ・日本テレビ系列全国ネット：12/4(土) 18:00～<br />
<br />
<b>・作品情報</b><br />
原作は2019年～2020年に週刊少年サンデーに掲載された「名探偵コナン」スピンオフ。全2巻。放送は「名探偵コナン」枠で不定期に行われる。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4098503670/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4098503670/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/03-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：青山剛昌、新井隆広<br />
コピーライト表記：&#x00A9;青山剛昌・新井隆広／小学館・読売テレビ・TMS2021 &#x00A9;青山剛昌／小学館・読売テレビ・TMS 1996<br />
<br />
<b>・キャスト</b><br />
古谷徹：降谷零/安室透/バーボン<br />
神奈延年：松田陣平<br />
東地宏樹：伊達航<br />
三木眞一郎：萩原研二<br />
緑川光：諸伏景光<br />
<br />
Twitter：<a href="https://twitter.com/conan_anime1000" target="_blank">@conan_anime1000</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/コナン警察学校編" target="_blank">#コナン警察学校編</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="04">◆<b><a href="https://kimetsu.com/anime/yukakuhen/" target="_blank">鬼滅の刃 遊郭編</a></b></h2><p class="preface">
</p><a href="https://kimetsu.com/anime/yukakuhen/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/04.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
フジテレビ：12/5(日) 23:15～<br />
北海道文化放送：12/5(日) 23:15～<br />
岩手めんこいテレビ：12/5(日) 23:15～<br />
仙台放送：12/5(日) 23:15～<br />
秋田テレビ：12/5(日) 23:15～<br />
さくらんぼテレビ：12/5(日) 23:15～<br />
福島テレビ：12/5(日) 23:15～<br />
NST新潟総合テレビ：12/5(日) 23:15～<br />
長野放送：12/5(日) 23:15～<br />
テレビ静岡：12/5(日) 23:15～<br />
富山テレビ：12/5(日) 23:15～<br />
石川テレビ：12/5(日) 23:15～<br />
福井テレビ：12/5(日) 23:15～<br />
東海テレビ：12/5(日) 23:15～<br />
関西テレビ：12/5(日) 23:15～<br />
TSKさんいん中央テレビ：12/5(日) 23:15～<br />
OHK岡山放送：12/5(日) 23:15～<br />
テレビ新広島：12/5(日) 23:15～<br />
テレビ愛媛：12/5(日) 23:15～<br />
高知さんさんテレビ：12/5(日) 23:15～<br />
テレビ西日本：12/5(日) 23:15～<br />
サガテレビ：12/5(日) 23:15～<br />
テレビ長崎：12/5(日) 23:15～<br />
テレビ熊本：12/5(日) 23:15～<br />
鹿児島テレビ：12/5(日) 23:15～<br />
沖縄テレビ：12/5(日) 23:15～<br />
TOKYO MX：12/11(土) 23:30～<br />
BS11：12/11(土) 23:30～<br />
群馬テレビ：12/11(土) 23:30～<br />
とちぎテレビ：12/11(土) 23:30～<br />
AT-X：12/13(月) 24:00～、ほか<br />
アニマックス：12/18(土) 21:00～<br />
<br />
配信：12/5(日) 25:00～(第2話以降24:00～)<br />
(ABEMA、Amazon Prime Video、auスマートパスプレミアム、dアニメストア、dアニメストア ニコニコ支店、dアニメストア for Prime Video、dTV、FOD、GYAO!、Hulu、J:COMオンデマンド、Netflix、Paravi、TELASA、TSUTAYA TV、U-NEXT、アニメ放題、アニメカ、ニコニコ生放送、ニコニコチャンネル、バンダイチャンネル、ひかりTV、DMM.com、GYAO!ストア、milplus、ビデオマーケット)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>無限列車での任務を終えた炭治郎たちは、アオイやなをに代わり、音柱・宇髄天元とともに次の任務地・遊郭へ向かう。</blockquote><p class="preface">
<br />
原作は2016年～2020年に週刊少年ジャンプに連載された吾峠呼世晴による漫画で、全23巻。アニメは2019年4月～9月に「竈門炭治郎 立志編」が放送され、2020年10月に劇場版「無限列車編」が公開された。2021年10月からは映画を全7話構成に再編集したテレビ版「無限列車編」が放送されていた。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4088812832/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4088812832/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/04-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：吾峠呼世晴<br />
監督：外崎春雄<br />
キャラクターデザイン・総作画監督：松島晃<br />
脚本制作：ufotable<br />
サブキャラクターデザイン：佐藤美幸、梶山庸子、菊池美花<br />
プロップデザイン：小山将治<br />
美術監督：衛藤功二<br />
撮影監督：寺尾優一<br />
3D監督：西脇一樹<br />
色彩設計：大前祐子<br />
編集：神野学<br />
音楽：梶浦由記、椎名豪<br />
アニメーション制作：ufotable<br />
コピーライト表記：<span translate="no">&#x00A9;吾峠呼世晴／集英社・アニプレックス・ufotable</span><br />
<br />
<b>・キャスト</b><br />
竈門炭治郎：花江夏樹<br />
竈門襧豆子：鬼頭明里<br />
我妻善逸：下野紘<br />
嘴平伊之助：松岡禎丞<br />
宇髄天元：小西克幸<br />
まきを：石上静香<br />
須磨：東山奈央<br />
雛鶴：種﨑敦美<br />
堕姫(上弦の陸)：沢城みゆき<br />
<br />
Twitter：<a href="https://twitter.com/kimetsu_off" target="_blank">@kimetsu_off</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/鬼滅の刃" target="_blank">#鬼滅の刃</a><br />
<br />
OP：Aimer「残響散歌」<br />
ED：Aimer「朝が来る」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=EedCPn2MF8E" target="_blank">テレビアニメ「鬼滅の刃」遊郭編 第2弾PV 12月5日（日）放送開始</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/EedCPn2MF8E" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="05">◆<b><a href="https://www.nhk.or.jp/matsuyama/debunekochan/" target="_blank">かなしきデブ猫ちゃん</a></b></h2><p class="preface">
</p><a href="https://www.nhk.or.jp/matsuyama/debunekochan/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/05.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
NHK総合(愛媛)：12/6(月) 8:42～<br />
NHK総合(全国)：12/19(日)～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>松山市のネコカフェに保護されていた１匹のチビネコ。その場所だけが自分の生きる世界だと信じていたある日、内気な少女アンナと出会い、その家族に引き取られる。「マル」という名前を与えられ、窓から大きな庭を見渡せる家での新しい生活が始まった。マルが驚いたのが食事のおいしさだった。魚の切り身の入った〝ご飯〟のすばらしさ。毎日食べて、ゴロゴロして。いつの間にか丸々と太ってしまった。幸せだったはずのマルの日常は、メスの子ネコ「スリジエ」が家族に仲間入りして一変する。家族の愛はスリジエに注がれ、そのスリジエからいたずらされてばかり。ある夜、我慢の限界に達し、スリジエを組み伏せようとしたマルの耳に響いたのは…。<br />
「コラ！デブ猫！どうしてスリジエをいじめるの！」〝大親友〟だったアンナから放たれたひと言。何よりも悲しそうなアンナの表情がマルの心を締め付けた。独り窓辺でくるまっていたその夜、突然、窓の向こうにクロネコが現れ、マルにこう告げる。「気高き者よ。その目で広い世界を見るのです」言われるままに立ち上がったマル。体の奥底の震えを感じながら、満月に向かってほえ叫んでいた。「ニャーーーーーン！」<br />
愛媛各地を東へ、西へ。愛と哀しみの大冒険がいよいよ始まる！</blockquote><p class="preface">
<br />
原作は愛媛新聞に連載された創作童話。2019年に絵本化された。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/486087143X/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/486087143X/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/05-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：文・早見和真、絵・かのうかりん<br />
監督・脚本：加藤道哉<br />
美術監督：渡邊洋一<br />
キャラクターデザイン：竹内進二<br />
音楽：長岡亮介<br />
アニメーション制作：サイクロングラフィックス<br />
制作：NHKエンタープライズ<br />
取材協力：愛媛新聞社<br />
制作・著作：NHK松山拠点放送局<br />
<br />
<b>・キャスト</b><br />
マル：滝藤賢一<br />
マドンナ/：水樹奈々<br />
アンナ：水樹奈々<br />
赤シャツ/おじいさん/ミヤギさん/シキさん：石住昭彦<br />
ママ/キヨ/ピース：渡辺美佐<br />
パパ/坊っちゃん/ヘミングウェイ：白石稔<br />
スリジエ/山嵐/カタマ/おばさん：小松里歌<br />
野だいこ/ジャック：福島潤<br />
<br />
</p><hr></p><p class="preface"><h2 id="27">◆<b><a href="https://akuyakureijou.af-original.com/" target="_blank">王子の本命は悪役令嬢</a></b></h2><p class="preface">
</p><a href="https://akuyakureijou.af-original.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/27.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/9(日) 25:00～<br />
BS11：1/9(日) 25:00～<br />
<br />
AnimeFesta：12/10(金) 18:00～(「オトナ向けプレミアム版」もあり)<br />
他配信：1/9(日) 25:00～<br />
(YouTube、ニコニコ動画、dアニメストア、U-NEXT、アニメ放題、ビデオマーケット)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>「もしかして王子…めちゃくちゃ悪役令嬢のことすきなの！？」<br />
ある日、社畜OLの主人公は大好きな乙女ゲームの世界に転生してしまう。<br />
それも、悪役令嬢・ディアナとして…！<br />
その上、ゲームのヒロインと結ばれるはずのシリウス王子が、なぜか自分(悪役令嬢)と急接近！？<br />
「王子に嫌われなきゃ！」バッドエンドを回避する為、王子に色仕掛けをするディアナ。<br />
しかし、王子はそれを嫌悪するどころか、熱を帯びた瞳で見つめてきて…。<br />
結ばれたらバッドエンド！甘くて淫らな禁断の恋物語。</blockquote><p class="preface">
<br />
原作は漫画配信サイト「Comic Festa」で配信されている女性向けコミック「<b><a href="https://comic.iowl.jp/titles/112705/volumes" target="_blank">転生初夜からむさぼりエッチ～王子の本命は悪役令嬢</a></b>」。既刊10巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B08P57N29D/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B08P57N29D/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/27-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：Re:mimu<br />
監督：騎日月<br />
脚本：黒崎エーヨ<br />
キャラクターデザイン・総作画監督：吉川真一<br />
色彩設計：わしみ<br />
美術設定：上垣裕起子<br />
美術監督：前塚太一<br />
撮影監督：桑良人、簡佳瑩<br />
編集：安田ナツキ<br />
音響監督：西山寛基(オンエア版/雨桑みくも(プレミアム版)<br />
音響制作：スタジオマウス<br />
アニメーション制作：studio HōKIBOSHI<br />
製作：彗星社<br />
コピーライト表記：<span translate="no">&#x00A9; Re:mimu／Suiseisha Inc.</span><br />
<br />
<b>・キャスト</b><br />
シリウス・イヴェール：高塚智人(オンエア版)/猿飛総司(プレミアム版)<br />
ディアナ・ディアフェル：関根明良(オンエア版)/春原りんね(プレミアム版)<br />
クリス：渋谷彩乃(オンエア版)/雨木実柑(プレミアム版)<br />
スピカ・ディアフェル：仲村宗悟(オンエア版)/深川緑(プレミアム版)<br />
<br />
Twitter：<a href="https://twitter.com/af_originaltl" target="_blank">@af_originaltl</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/王子の本命は悪役令嬢" target="_blank">#王子の本命は悪役令嬢</a><br />
<br />
OP：シリウス・イヴェール(CV.高塚智人/猿飛総司)「アストロラーベを狂わせて」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=fXVDR61b8aA" target="_blank">【公式】TVアニメ「王子の本命は悪役令嬢」オンエア版Ver.2022年1月放送スタート！【PV】</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/fXVDR61b8aA" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
<b><a href="https://www.youtube.com/watch?v=Vxcz3lxPlVE" target="_blank">【公式】TVアニメ「王子の本命は悪役令嬢」プレミアム版Ver.2022年1月放送スタート！【PV】</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/Vxcz3lxPlVE" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="06">◆<b><a href="https://yostar-pictures.co.jp/sorairo-utility/" target="_blank">空色ユーティリティ</a></b></h2><p class="preface">
</p><a href="https://yostar-pictures.co.jp/sorairo-utility/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/06.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：12/31(金) 19:30～<br />
本編15分＋キャスト＆スタッフトーク15分<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>普段と違う早起きした朝、クラブと一緒にわくわくした気持ちをゴルフバッグに詰めて、今日も一日を始めよう。<br />
主人公の美波はゴルフを始めて間もない初心者。<br />
遥と彩花はそんな美波を見守りつつ、それぞれのスタンスでゴルフを楽しみ、その楽しさを美波に伝えていく。<br />
学校も違えば好みも違う三人。<br />
だけどその三人がゴルフを通じて、それぞれの世界をもっと色鮮やかに変えていく。<br />
これはそんな物語の1ページ。</blockquote><p class="preface">
<br />
「アズールレーン」「雀魂 -じゃんたま-」「ブルーアーカイブ -Blue Archive-」などの運営で知られるYostarが設立したアニメーション制作会社・Yostar Picturesによる初のオリジナルアニメ。<br />
<br />
<b>・スタッフ</b><br />
監督・キャラクターデザイン・演出・作画監督：斉藤健吾<br />
脚本：望公太<br />
絵コンテ：雨宮哲<br />
制作：Yostar Pictures<br />
コピーライト表記：<span translate="no">&#x00A9; 2021 Yostar Pictures, Inc. All rights reserved.</span><br />
<br />
<b>・キャスト</b><br />
美波：高木美佑<br />
遥：天海由梨奈<br />
彩花：後藤彩佐<br />
<br />
Twitter：<a href="https://twitter.com/YostarPictures" target="_blank">@YostarPictures</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/空色ユーティリティ" target="_blank">#空色ユーティリティ</a><br />
<br />
<b><a href="https://www.youtube.com/watch?v=zvuJ-D1W_bU" target="_blank">TVアニメーション『空色ユーティリティ』予告映像</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/zvuJ-D1W_bU" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="40">◆<b><a href="https://anime.elmelloi.com/" target="_blank">ロード・エルメロイII世の事件簿 -魔眼蒐集列車 Grace note- 特別編</a></b></h2><p class="preface">
</p><a href="https://anime.elmelloi.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/40.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：12/31(金) 22:00～<br />
BS11：12/31(金) 22:00～<br />
群馬テレビ：12/31(金) 22:00～<br />
とちぎテレビ：12/31(金) 22:00～<br />
<br />
ニコニコ生放送<br />
ABEMA<br />
<br />
<b>・作品情報</b><br />
大晦日特番“Fate Project 大晦日TVスペシャル2021”内で放送される。<br />
<br />
<b>・スタッフ</b><br />
原作：三田誠／TYPE-MOON <br />
キャラクター原案：坂本みねぢ<br />
監督：加藤誠<br />
キャラクターデザイン：中井準<br />
スーパーバイザー：あおきえい<br />
音楽：梶浦由記<br />
制作：TROYCA<br />
コピーライト表記：<span translate="no">&#x00A9; 三田誠・TYPE-MOON / LEMPC</span><br />
<br />
<b>・キャスト</b><br />
ロード・エルメロイⅡ世：浪川大輔<br />
グレイ：上田麗奈<br />
<br />
Twitter：<a href="https://twitter.com/elmelloi_anime" target="_blank">@elmelloi_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/事件簿アニメ" target="_blank">#事件簿アニメ</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="07">◆<b><a href="https://mahouka.jp/" target="_blank">魔法科高校の劣等生 追憶編</a></b></h2><p class="preface">
</p><a href="https://mahouka.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/07.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：12/31(金) 23:57～<br />
とちぎテレビ：12/31(金) 23:57～<br />
群馬テレビ：12/31(金) 23:57～<br />
BS11：12/31(金) 23:57～<br />
AT-X：1/1(土) 20:30～、ほか<br />
テレビ愛知：1/3(月) 27:10～<br />
MBS：1/4(火) 26:35～<br />
<br />
ABEMA：12/31(金) 24:00～<br />
dアニメストア：12/31(金) 24:00～<br />
他配信：1/6(木) 24:00～<br />
(GYAO!、ひかりTV、Netflix、FOD、バンダイチャンネル、Hulu、TELASA、J:COMオンデマンド、みるプラス見放題プレミアム、U-NEXT、アニメ放題、Amazon Prime Video、アニメカ、TVer、MBS動画イズム)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>魔法が技術として確立され、約一世紀が過ぎた二〇九五年四月。魔法師を育成する国立魔法大学付属第一高校、通称“魔法科高校”に二人の兄妹が入学した。一人は魔法師として致命的な欠陥を抱える劣等生の兄・達也。もう一人は完全無欠の魔法師として讃えられる優等生の妹・深雪。<br />
時に恋人同士と間違われるほど仲睦まじい二人だが、ほんの数年前までは、まるで女主人と使用人のように冷え切った関係だった。<br />
その関係が変わった背景には、ある出来事があった。三年前の沖縄。深雪にとって忘れられない出来事によって、二人の心と、その運命が大きく変わっていく。</blockquote><p class="preface">
<br />
原作は佐島勤によるライトノベルシリーズ。全32巻で続編も出ている作品で、追憶編は8巻にあたる。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4048911589/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4048911589/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/07-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：佐島勤<br />
監督：吉田りさこ<br />
助監督：ジミーストーン<br />
脚本：中本宗応(ライトワークス)<br />
キャラクター原案・総作画監督：石田可奈<br />
色彩設計：小松さくら<br />
撮影監督：川下裕樹(MADBOX)、藤村兵悟(MADBOX)<br />
撮影：MADBOX<br />
編集：木村佳史子<br />
音響監督：本山哲<br />
音楽：岩崎琢<br />
アニメーション制作：エイトビット<br />
コピーライト表記：<span translate="no">&#x00A9;2019 佐島 勤/KADOKAWA/魔法科高校2製作委員会</span><br />
<br />
<b>・キャスト</b><br />
司波達也：中村悠一<br />
司波深雪：早見沙織<br />
司波深夜：井上喜久子<br />
桜井穂波：遠藤綾<br />
<br />
Twitter：<a href="https://twitter.com/mahouka_anime" target="_blank">@mahouka_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/mahouka" target="_blank">#mahouka</a><br />
<br />
OP：八木海莉「Ripe Aster」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=BKkKlE6aB6I" target="_blank">TVアニメ「魔法科高校の劣等生 追憶編」第１弾PV｜2021年12月31(金)23:57より放送決定！</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/BKkKlE6aB6I" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="51">◆<b><a href="https://chunithm.sega.jp/irodorimidori/anime/" target="_blank">イロドリミドリ</a></b></h2><p class="preface">
</p><a href="https://chunithm.sega.jp/irodorimidori/anime/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/51.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/4(火) 25:00～<br />
AT-X：1/6(木) 22:00～、ほか<br />
<br />
配信：1/4(火) 25:05～順次開始<br />
(Amazon Prime Video、dアニメストア、U-NEXT、バンダイチャンネル、ニコニコ動画、FOD、アニメカ、アニメ放題、ABEMA、auスマートパスプレミアム、J:COMオンデマンド メガパック、TELASA、ひかりTV、Hulu、みるプラス)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>舞ヶ原音楽大学付属舞ヶ原高等学校(通称まいまい)。<br />
才能あるミュージシャンの卵が全国から集まるこの学び舎には、ある噂がある。<br />
『ライブですごいパフォーマンスができたら、追加で特別な点数がもらえる(らしい)』<br />
単位の危ない芹菜は噂を信じてバンドを結成！<br />
演奏したり、お泊りしたり、コスプレしたり？な、いろとりどりの毎日。<br />
学園フェスでいい演奏をして無事に単位をゲットできるのか!?<br />
ゆるゆる日常系ガールズバンドストーリー開演！</blockquote><p class="preface">
<br />
セガが展開する音楽ゲーム「<b><a href="https://chunithm.sega.jp/" target="_blank">CHUNITHM(チュウニズム)</a></b>」から生まれ、マンガ動画を軸にメインストーリーを展開し、オリジナル楽曲リリースやライブイベント開催を行っているガールズバンド「<b><a href="https://chunithm.sega.jp/irodorimidori/" target="_blank">イロドリミドリ</a></b>」のショートアニメ。<br />
<br />
<b>・スタッフ</b><br />
原作：七条レタス(IOSYS)、齊藤キャベツ(SEGA)<br />
キャラクター原案：Hisasi<br />
企画：CHUNITHM <br />
監督：田中千駿<br />
シリーズ構成：田中千駿、石倉礼<br />
キャラクターデザイン：佐藤健史、後藤香織<br />
美術監督：魏斯曼(St.PILZ)<br />
色彩設計：池田ひとみ<br />
撮影監督：石山智之<br />
編集：柳圭介<br />
音響監督：西山寛基<br />
音楽：D.watt(IOSYS) <br />
デザインサポート：佐藤啓太(SEGA)<br />
アニメーションプロデューサー：映月<br />
アニメーション制作：暁<br />
プロデューサー：小清水純(SEGA)、跡部泰広(フロントウイング)<br />
アシスタントプロデューサー：植村有紀子(SEGA)、石坂竜二(フロントウイング)、河野麗生(フロントウイング)<br />
プロデュース・宣伝：SEGA、フロントウイング<br />
製作：イロドリミドリ新聞部<br />
コピーライト表記：<span translate="no">&#x00A9;SEGA</span><br />
<br />
<b>・キャスト</b><br />
明坂芹菜：新田恵海<br />
御形アリシアナ：福原綾香<br />
天王洲なずな：山本彩乃<br />
小仏凪：佐倉薫<br />
箱部なる：M・A・O<br />
<br />
Twitter：<a href="https://twitter.com/irodorimidori5" target="_blank">@irodorimidori5</a><br />
<br />
<b><a href="https://www.youtube.com/watch?v=JkC-Y8xuNRo" target="_blank">ショートアニメ『イロドリミドリ』ティザーPV｜2022年1月4日よりTV放送スタート！</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/JkC-Y8xuNRo" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="08">◆<b><a href="https://leadale.net/" target="_blank">リアデイルの大地にて</a></b></h2><p class="preface">
</p><a href="https://leadale.net/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/08.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/5(水) 23:30～<br />
サンテレビ：1/5(水) 24:00～<br />
KBS京都：1/5(水) 24:30～<br />
BS日テレ：1/5(水) 24:30～<br />
AT-X：1/6(木) 21:30～、ほか<br />
<br />
dアニメストア：1/5(火) 23:00～【日本国内独占・地上波先行】<br />
crunchyroll：1月～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>不慮の事故によって生命維持装置がないと生きられない少女、各務桂菜。彼女が唯一自由でいられるのは、VRMMORPG『リアデイル』の中だけだった。そんなある日、生命維持装置が停止し、桂菜は命を落としてしまう。<br />
しかし、目が覚めると桂菜は200年後の『リアデイル』の世界にいた。彼女は、失われたスキルと限界突破したステータスをもつハイエルフの”ケーナ”として、この世界を生きる人々と交流を深めていくことに。しかも、その中にはかつて自らがキャラメイクした"子供"たちもいて……！？<br />
ゲーム世界に転生した少女と個性的な仲間たちによる、笑いあり、涙あり？なのんびり冒険譚、開幕！</blockquote><p class="preface">
<br />
原作はCeezによる小説。「小説家になろう」で2011年・2012年の年間ランキングトップ5に入った作品。2019年からKADOKAWA・ファミ通文庫で書籍版が刊行されている。既刊7巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4047354694/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4047354694/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/08-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
月見だしお作画によるコミック版は最新4巻が2022年1月8日発売予定。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4049129396/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4049129396/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/08-03.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：Ceez<br />
原作イラスト：てんまそ<br />
監督：柳瀬雄之<br />
シリーズ構成：筆安一幸<br />
キャラクターデザイン：舛舘俊秀、小島えり、出口花穂<br />
音響監督：土屋雅紀<br />
音楽：夢見クジラ<br />
アニメーション制作：MAHO FILM<br />
製作：リアデイル製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;2021 Ceez,てんまそ/KADOKAWA/リアデイル製作委員会</span><br />
<br />
<b>・キャスト</b><br />
ケーナ：幸村恵理<br />
スカルゴ：小野大輔<br />
マイマイ：名塚佳織<br />
カータツ：杉田智和<br />
キー：小田切優衣<br />
リット：柳原かなこ<br />
マレール：中村桜<br />
ルイネ：広瀬ゆうき<br />
<br />
Twitter：<a href="https://twitter.com/leadale_anime" target="_blank">@leadale_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/leadale" target="_blank">#leadale</a> <a href="https://twitter.com/hashtag/リアデイル" target="_blank">#リアデイル</a><br />
<br />
OP：TRUE「Happy encount」<br />
ED：田所あずさ「箱庭の幸福」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=J9awZtCDOPQ" target="_blank">TVアニメ「リアデイルの大地にて」PV第2弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/J9awZtCDOPQ" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="09">◆<b><a href="https://hakozume-anime.com/" target="_blank">ハコヅメ～交番女子の逆襲～</a></b></h2><p class="preface">
</p><a href="https://hakozume-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/09.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
AT-X：1/5(水) 23:30～、ほか<br />
TOKYO MX：1/5(水) 25:05～<br />
KBS京都：1/5(水) 25:05～<br />
サンテレビ：1/5(水) 25:30～<br />
テレビ愛知：1/5(水) 26:35～<br />
熊本県民テレビ：1/6(木) 25:29～<br />
BS日テレ：1/6(木) 24:00～<br />
<br />
dアニメストア：1/5(水) 24:00～<br />
Amazon Prime Video：1/5(水) 24:00～<br />
他配信：1/12(水) 24:00～<br />
(ABEMA、dTV、FOD、GYAO！、Hulu、J:COMオンデマンド、Netflix、niconico、TELASA、U-NEXT、アニメ放題、スマートパスプレミアム、ひかりTV、バンダイチャンネル、みるプラス、HAPPY!動画)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>「警察官なんて、もう辞めてやる！」<br />
公務員試験を片っ端から受けて、合格したのは警察官だけ。考えつく限り、最も浅い理由で警察官になった川合麻依は後悔していた。こんなに激務で嫌われ者だって知ってたら、絶対に警察官になんてなってない！<br />
辞表を握りしめて、第二の人生を歩むことを決意する川合のもとに、新しい指導員としてやってきたのは、警察学校を主席で卒業し、"ミス・パーフェクト"の異名を持つ元刑事課のエース・藤聖子。<br />
後輩へのパワハラが過ぎて、刑事課から交番に異動してきたという藤の噂に怯える川合だったが、さっそくペアとしてパトロールに向かうことに……。<br />
新人警察官・川合と元刑事課のエース・藤の凸凹ペアを中心に、個性豊かで魅力的な警察官たちが巻き起こす笑って驚いて、ときどき涙しちゃうお仕事コメディが今、幕を開ける！</blockquote><p class="preface">
<br />
原作は「モーニング」連載の漫画で、既刊19巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4065113229/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4065113229/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/09-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
戸田恵梨香と永野芽郁のW主演で2021年にドラマ化されており、<b><a href="https://twitter.com/hakozume_ntv/status/1467962048281604100" target="_blank">12月28日～31日には年末一挙放送スペシャル</a></b>がある。<br />
<br />
<b>・スタッフ</b><br />
原作：泰三子<br />
監督：佐藤雄三<br />
シリーズ構成：金月龍之介<br />
キャラクターデザイン：土屋圭<br />
副監督：石田暢<br />
美術ボード：橋本和幸、横松紀彦<br />
美術設定：杉山晋史<br />
色彩設計：大野春恵<br />
撮影監督：畑中宏伸<br />
3D監督：田中康隆<br />
編集：塚常真理子<br />
音楽：信澤宣朋<br />
音響監督：小泉紀介<br />
音響効果：山谷尚人<br />
録音スタジオ：スタジオT&T<br />
音響制作：マジックカプセル<br />
制作プロデューサー：芦川真理子、豊田智紀<br />
アニメーション制作：マッドハウス<br />
製作：ハコヅメ製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;泰三子・講談社／ハコヅメ製作委員会</span><br />
<br />
<b>・キャスト</b><br />
川合麻依：若山詩音<br />
藤聖子：石川由依<br />
源誠二：鈴木崚汰<br />
山田武志：土屋神葉<br />
牧高美和：花澤香菜<br />
<br />
Twitter：<a href="https://twitter.com/hakozume_anime" target="_blank">@hakozume_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/ハコヅメ" target="_blank">#ハコヅメ</a><br />
<br />
OP：安月名莉子「知らなきゃ」<br />
ED：nonoc「Change」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=gbYFaLhM2FU" target="_blank">TVアニメ「ハコヅメ～交番女子の逆襲～」PV第1弾【2022年1月放送開始!!】</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/gbYFaLhM2FU" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="10">◆<b><a href="https://tokyo24project.com/" target="_blank">東京24区</a></b></h2><p class="preface">
</p><a href="https://tokyo24project.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/10.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/5(水) 24:00～(第2話以降24:30～)<br />
とちぎテレビ：1/5(水) 24:00～(第2話以降24:30～)<br />
群馬テレビ：1/5(水) 24:00～(第2話以降24:30～)<br />
BS11：1/5(水) 24:00～(第2話以降24:30～)<br />
メ～テレ：1/5(水) 25:56～(第2話以降26:11～)<br />
ABCテレビ：1/5(水) 26:14～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>東京湾に浮かぶ人工島「極東法令外特別地区」――通称“24区”。<br />
そこで生まれ育ったシュウタ、ラン、コウキは、家柄も趣味も性格も違うが、いつもつるんでいる幼馴染だった。<br />
しかし彼らの関係は、とある事件をきっかけに大きく変わってしまう。<br />
事件の一周年追悼ミサで、偶然再会を果たした3人の電話が突如一斉に鳴る。<br />
それは死んだはずの仲間からの着信で、彼らに“未来の選択”を迫るものであった。<br />
3人は、自分の信じるやり方で、愛する24区と人々の未来を守ろうとするが―――</blockquote><p class="preface">
<br />
アダルトゲーム「<b><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00EVLUHUC/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B00EVLUHUC/gigazine-22');">スマガ</a></b>」「<b><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00DVHTTM4/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B00DVHTTM4/gigazine-22');">君と彼女と彼女の恋。</a></b>」などを手がけたニトロプラスのシナリオライター・<b><a href="https://twitter.com/shimokura_vio" target="_blank">下倉バイオ</a></b>が初めてアニメの脚本に挑戦するオリジナルアニメ。<br />
<br />
<b>・スタッフ</b><br />
原作：Team 24<br />
監督：津田尚克<br />
ストーリー構成・脚本：下倉バイオ(ニトロプラス)<br />
キャラクター原案：曽我部修司、ののかなこ(FiFS)<br />
キャラクターデザイン：岸田隆宏<br />
総作画監督：髙田晃、伊藤公規、まじろ<br />
副監督：髙橋英俊<br />
プロップデザイン：髙田晃<br />
グラフィティデザイン：河原秀樹、添野恵<br />
美術設定：塩澤良憲<br />
美術監督：春日美波<br />
色彩設計：中島和子<br />
2Dデザイン：久保田彩<br />
特殊効果：清水彩香<br />
CGディレクター：宮地克明<br />
撮影監督：佐久間悠也<br />
編集：三嶋章紀<br />
音響監督：岩浪美和<br />
音楽：深澤秀行<br />
制作：CloverWorks<br />
コピーライト表記：<span translate="no">&#x00A9;Team24／東京24区プロジェクト</span><br />
<br />
<b>・キャスト</b><br />
蒼生シュウタ：榎木淳弥<br />
朱城ラン：内田雄馬<br />
翠堂コウキ：石川界人<br />
翠堂アスミ：石見舞菜香<br />
櫻木まり：牧野由依<br />
きなこ：兎丸七海<br />
黒葛川早紀子：生天目仁美<br />
翠堂豪理：楠大典<br />
翠堂香苗：大原さやか<br />
クナイ：斉藤壮馬<br />
ヤマモリ：伊丸岡 篤<br />
ラッキー：喜多村英梨<br />
白樺広樹：上田燿司<br />
白樺 梢：日高里菜<br />
筑紫 渉：中村悠一<br />
宍戸花奈：花守ゆみり<br />
進藤薫：江頭宏哉<br />
<br />
Twitter：<a href="https://twitter.com/tokyo24_project" target="_blank">@tokyo24_project</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/東京24区" target="_blank">#東京24区</a><br />
<br />
OP：Survive Said The Prophet「Paper Sky」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=K-VFdafbZ7s" target="_blank">オリジナルTVアニメ「東京24区」第１弾PV＜2022年1月5日(水)放送開始＞</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/K-VFdafbZ7s" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="11">◆<b><a href="https://orient-anime.jp/" target="_blank">オリエント</a></b></h2><p class="preface">
</p><a href="https://orient-anime.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/11.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ東京：1/5(水) 24:00～<br />
BSテレ東：1/5(水) 24:30～<br />
AT-X：1/6(木) 21:00～、ほか<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>鬼神の襲来により、日ノ本の覇権は“人”から“鬼”へと移った。<br />
人々が鬼による支配を受け入れ、“武士”だけが戦い続ける動乱の時代に、武蔵と小次郎は“最強の武士団”結成の夢を誓い、鬼退治に挑む！</blockquote><p class="preface">
<br />
原作は週刊少年マガジン連載で、最新14巻が2022年1月7日発売予定。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4065127254/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4065127254/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/11-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：大高忍<br />
監督：柳沢テツヤ<br />
シリーズ構成：國澤真理子<br />
キャラクターデザイン：岸田隆宏<br />
総作画監督：松本文男、崎本さゆり<br />
プロップデザイン：きむらひでふみ<br />
色彩設計：日比野 仁<br />
美術設定：前田みつき<br />
美術監督：坂上裕文<br />
撮影監督：間中秀典<br />
編集：山岸歩奈美（REAL-T）<br />
音響監督：納谷僚介<br />
音響制作：スタジオマウス<br />
音楽：深澤秀行<br />
アニメーション制作：Ａ・Ｃ・Ｇ・Ｔ<br />
製作：「オリエント」製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;大高忍・講談社／「オリエント」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
武蔵：内田雄馬<br />
鐘巻小次郎：斉藤壮馬<br />
服部つぐみ：高橋李依<br />
武田尚虎：日野聡<br />
犬飼四郎：下野紘<br />
犬坂七緒：和氣あず未<br />
小雨田英雄：羽多野渉<br />
<br />
Twitter：<a href="https://twitter.com/orient_PR" target="_blank">@orient_PR</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/オリエント" target="_blank">#オリエント</a><br />
<br />
OP：Da-iCE「Break out」<br />
ED：羽多野渉「ナニイロ」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=DAV19xp-0K0" target="_blank">ＴＶアニメ「オリエント」本PV【2022年1月5日より放送】</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/DAV19xp-0K0" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="12">◆<b><a href="https://saiyuki-r-zeroin.jp/" target="_blank">最遊記RELOAD -ZEROIN-</a></b></h2><p class="preface">
</p><a href="https://saiyuki-r-zeroin.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/12.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
AT-X：1/6(木) 23:00～、ほか<br />
TOKYO MX：1/6(木) 24:30～<br />
BS11：1/6(木) 24:30～<br />
MBS：1/8(土) 27:08～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>人と妖怪、科学と妖術が共存を果たす無秩序かつ安寧の大陸『桃源郷』。<br />
しかし牛魔王蘇生実験による負の波動の影響で、突然妖怪たちが暴走し、その均衡は崩れてしまった。<br />
天界の観世音菩薩に牛魔王蘇生実験の阻止を命じられた玄奘三蔵は、孫悟空、沙悟浄、猪八戒らとともに、西域・天竺を目指す。<br />
その道中に辿り着いた地で三蔵たちは、死んだ人間を生き返らせることができる蘇生術を持つヘイゼル＝グロースとその従者ガトと邂逅する。<br />
西域からきたという、ヘイゼルの目的とは―――――。</blockquote><p class="preface">
<br />
原作は、中国の小説「西遊記」をモチーフに展開されている峰倉かずやの漫画。さまざまなメディアミックスが行われており、テレビアニメ化だけでも今回が5度目となる。<br />
<br />
なお、今回アニメ化されるのは「ヘイゼル編」に相当。2004年4月～9月に放送された「<b><a href="https://www.tv-tokyo.co.jp/anime/saiyuki_rg/" target="_blank">最遊記RELOAD GUNLOCK</a></b>」の後半でもアニメ化されているが、このときは冒頭部分以外はアニメオリジナルだったため、原作者は当時のアニメしか見ていない人に向けて「<b><a href="https://saiyuki-r-zeroin.jp/special_cat/other/" target="_blank">『ひぐらしのなく頃に 業』を観るようなお気持ちで愉しんで頂けたら</a></b>」とコメントしている。<br />
<br />
<b>・スタッフ</b><br />
原作：峰倉かずや<br />
監督：髙田美里<br />
シリーズ構成：横手美智子、松井亜弥<br />
キャラクターデザイン・プロップデザイン・総作画監督：小倉典子<br />
美術監督：深井絵梨香<br />
美術設定：青木智由紀、イノセユキエ<br />
色彩設計：野地弘納<br />
CGディレクター：山崎嘉雅<br />
撮影監督：後藤晴香<br />
編集：梅津朋美<br />
音楽：白戸佑輔<br />
音楽制作：ランティス<br />
音響監督：髙桑一<br />
アニメーションプロデューサー：柴宏和<br />
アニメーション制作：ライデンフィルム<br />
コピーライト表記：<span translate="no">&#x00A9; 峰倉かずや・一迅社／最遊記RE PROJECT</span><br />
<br />
<b>・キャスト</b><br />
玄奘三蔵：関俊彦<br />
孫悟空：保志総一朗<br />
沙悟浄：平田広明<br />
猪八戒：石田彰<br />
ヘイゼル＝グロース：遠近孝一<br />
ガティ＝ネネホーク：小山力也<br />
白竜：戸田めぐみ<br />
你健一/烏哭三蔵：大塚芳忠<br />
光明三蔵：宮本充<br />
<br />
Twitter：<a href="https://twitter.com/saiyuki_re" target="_blank">@saiyuki_re</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/最遊記" target="_blank">#最遊記</a> <a href="https://twitter.com/hashtag/最遊記ZEROIN" target="_blank">#最遊記ZEROIN</a><br />
<br />
OP：GRANRODEO「カミモホトケモ」<br />
ED：仲村宗悟「流転」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=_yoVUELjMf0" target="_blank">TVアニメ「最遊記RELOAD -ZEROIN-」孫悟空キャラクターPV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/_yoVUELjMf0" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="13">◆<b><a href="https://arifureta.com/" target="_blank">ありふれた職業で世界最強 2nd season</a></b></h2><p class="preface">
</p><a href="https://arifureta.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/13.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
・ありふれた職業で世界最強 2nd season 放送直前 作品を盛り上げようスペシャル！<br />
AT-X：1/6(木) 23:30～、ほか<br />
TOKYO MX：1/6(木) 24:00～<br />
BS11：1/6(木) 24:00～<br />
<br />
・本編<br />
AT-X：1/13(木) 23:30～、ほか<br />
TOKYO MX：1/13(木) 24:00～<br />
BS11：1/13(木) 24:00～<br />
<br />
dアニメストア：1/13(木) 24:00～【地上波同時】<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>かつてのクラスメイトたちと別れた後、主人公ハジメが率いるパーティーは新たな神代魔法を手に入れるため大迷宮へと向かっていた。<br />
様々な試練を突破し、元の世界に帰るために着々と力をつけていく一行。しかし、その裏では魔人族たちも不穏な動きをみせていた。<br />
王都に忍び寄る影と暗躍する謎の人物。<br />
「敵はすべて倒す――たとえそれが神だとしても」<br />
彼らの前に立ちはだかる真の敵とは一体何者なのか――!?<br />
“最強”異世界ファンタジー、セカンドシーズン開幕！</blockquote><p class="preface">
<br />
原作は「<b><a href="https://ncode.syosetu.com/n8611bv/" target="_blank">小説家になろう</a></b>」連載作品。オーバーラップ文庫から刊行されている書籍版が既刊11巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4865546987/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4865546987/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/13-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：白米良<br />
イラスト：たかやKi<br />
監督：岩永彰<br />
シリーズ構成・脚本：佐藤勝一<br />
キャラクターデザイン・総作画監督：小島智加<br />
音楽：高橋諒<br />
アニメーション制作：asread.×studio MOTHER<br />
コピーライト表記：<span translate="no">&#x00A9;Ryo Shirakome, OVERLAP/ARIFURETA Project</span><br />
<br />
<b>・キャスト</b><br />
南雲ハジメ：深町寿成<br />
ユエ：桑原由気<br />
シア・ハウリア：髙橋ミナミ<br />
ティオ・クラルス：日笠陽子<br />
白崎香織：大西沙織<br />
八重樫雫：花守ゆみり<br />
畑山愛子：加隈亜衣<br />
ミュウ：小倉唯<br />
リリアーナ・S・B・ハイリヒ：芝崎典子<br />
ノイント：佐藤利奈<br />
<br />
Twitter：<a href="https://twitter.com/ARIFURETA_info" target="_blank">@ARIFURETA_info</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/ありふれた" target="_blank">#ありふれた</a><br />
<br />
OP：MindaRyn「Daylight」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=a4VkLBHiPjo" target="_blank">TVアニメ「ありふれた職業で世界最強 2nd season」PV第2弾｜2022.Jan ON AIR</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/a4VkLBHiPjo" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="14">◆<b><a href="https://slowlooptv.com/" target="_blank">スローループ</a></b></h2><p class="preface">
</p><a href="https://slowlooptv.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/14.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
AT-X：1/7(金) 22:00～、ほか<br />
TOKYO MX：1/7(金) 22:30～<br />
BS11：1/7(金) 23:00～<br />
サンテレビ：1/7(金) 24:00～<br />
KBS京都：1/7(金) 24:00～<br />
テレビ愛知：1/7(金) 27:05～<br />
<br />
dアニメストア：1/7(金) 22:30～<br />
ABEMA：1/10(月) 22:30～<br />
他配信：1/14(金) 22:30～<br />
(niconico、GYAO!、ひかりTV、FOD、バンダイチャンネル、Hulu、TELASA、J:COMオンデマンド、みるプラス見放題プレミアム、U-NEXT、アニメ放題、Amazon Prime Video、アニメカ)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>海辺でひとり、亡き父に教えてもらったフライフィッシングを嗜む少女・ひより。<br />
いつもどおりに釣りをしていると、いきなり海に入ろうとする天真爛漫な少女・小春と出会います。<br />
一緒に釣りをする事になった2人でしたが、実は親の再婚相手の娘どうしで…？<br />
ひょんな出会いから「姉妹」になったひよりと小春と一緒に、｢釣り」をしながらスローに過ごしてみませんか？</blockquote><p class="preface">
<br />
原作はまんがタイムきららフォワード連載で、既刊5巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4832270761/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4832270761/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/14-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<br />
<b>・スタッフ</b><br />
原作：うちのまいこ<br />
監督：秋田谷典昭<br />
副監督：守田芸成<br />
シリーズ構成：山田由香<br />
キャラクターデザイン・総作画監督：滝本祥子<br />
釣りシーン演出：柴田匠<br />
料理・エフェクト作監：鷲北恭太<br />
美術監督：諸熊倫子(スタジオ天神)<br />
色彩設計：月野えりか<br />
撮影監督：佐藤敦(スタジオシャムロック)<br />
3D監督：濱村敏郎(ワイヤード)<br />
編集：仙土真希(REAL-T)<br />
音響監督：土屋雅紀<br />
音楽：伊賀拓郎<br />
音楽制作：フライングドッグ<br />
アニメーション制作：CONNECT<br />
製作：スローループ製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;うちのまいこ・芳文社／スローループ製作委員会</span><br />
<br />
<b>・キャスト</b><br />
海凪ひより：久住琳<br />
海凪小春：日岡なつみ<br />
吉永恋：嶺内ともみ<br />
福元一花：名塚佳織<br />
福元二葉：村上奈津実<br />
二宮藍子：井上ほの花<br />
<br />
Twitter：<a href="https://twitter.com/slowloop_tv" target="_blank">@slowloop_tv</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/slowloop" target="_blank">#slowloop</a><br />
<br />
OP：ぽかぽかイオン「やじるし→」<br />
ED：Three ∞ Loop「シュワシュワ」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=LyxCk1GUais" target="_blank">TVアニメ「スローループ」PV第2弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/LyxCk1GUais" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="15">◆<b><a href="https://vanitas-anime.com/" target="_blank">ヴァニタスの手記(第2クール・ジェヴォーダン編)</a></b></h2><p class="preface">
</p><a href="https://vanitas-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/15.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
・「ヴァニタスの手記」2クール目放送直前特番 ～En route pour le Gévaudan～<br />
TOKYO MX：1/7(金) 24:00～<br />
とちぎテレビ：1/7(金) 24:00～<br />
群馬テレビ：1/7(金) 24:00～<br />
BS11：1/7(金) 24:00～<br />
MBS：1/8(土) 26:38～<br />
RCC：1/11(火) 25:27～<br />
CBC：1/12(水) 26:35～<br />
北海道放送：1/13(木) 25:58～<br />
RKB毎日放送：1/13(木) 27:00～<br />
<br />
・本編<br />
TOKYO MX：1/14(金) 24:00～<br />
とちぎテレビ：1/14(金) 24:00～<br />
群馬テレビ：1/14(金) 24:00～<br />
BS11：1/14(金) 24:00～<br />
MBS：1/15(土) 26:38～<br />
RCC：1/18(火) 25:27～<br />
CBCテレビ：1/19(水) 26:35～<br />
北海道放送：1/20(木) 25:58～<br />
RKB毎日放送：1/20(木) 27:00～<br />
アニマックス：1/22(土) 22:00～<br />
<br />
Amazon Prime Video：1/14(金) 24:30～<br />
配信：1/18(火) 12:00～<br />
(dアニメストア、dアニメストア for Prime Video、GYAO！、バンダイチャンネル、Hulu、ビデオマーケット、music.jp、ひかりTV、FOD、U-NEXT、アニメ放題、ABEMAビデオ、niconico、Netflix、J:COMオンデマンド、TELASA、みるプラス)<br />
<br />
<b>・作品情報</b><br />
吸血鬼・ノエと「吸血鬼の専門家」ヴァニタスが繰り広げる吸血鬼冒険譚。2人は、18世紀に出現したという「ジェヴォーダンの獣」が再び現れたという話を聞きつけ、呪持ちとの関係を確かめるため、ジェヴォーダンに向かう。<br />
<br />
原作は「ガンガンJOKER」連載で、既刊9巻。アニメ第1期が2021年7月～9月に放送されており、今回のアニメで描かれる「ジェヴォーダン編」は5巻からの内容に相当する。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4757557582/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4757557582/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/15-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：望月淳<br />
監督：板村智幸<br />
シリーズ構成：赤尾でこ<br />
キャラクターデザイン・総作画監督：伊藤嘉之<br />
サブキャラクターデザイン：中山知世<br />
プロップデザイン：石橋慎平<br />
ビジュアルコンセプト・色彩設計：滝沢いづみ<br />
美術デザイン：多田周平<br />
美術監督：金井眞悟<br />
撮影監督：張盈穎<br />
3DCG監督：三宅拓馬<br />
編集：松原理恵<br />
音楽：梶浦由記<br />
音響監督：若林和弘<br />
音響効果：倉橋静男(サウンドボックス)、西佐知子(サウンドボックス)<br />
アニメーション制作：ボンズ<br />
コピーライト表記：<span translate="no">&#x00A9;望月淳／SQUARE ENIX・「ヴァニタスの手記」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
ヴァニタス：花江夏樹<br />
ノエ：石川界人<br />
ジャンヌ：水瀬いのり<br />
ルカ：下地紫野<br />
ドミニク：茅野愛衣<br />
先生：石田彰<br />
ムル：小牧未侑<br />
ダンテ：木内太郎<br />
ヨハン：遊佐浩二<br />
リーチェ：久間梨穂<br />
ローラン：河西健吾<br />
オリヴィエ：前野智昭<br />
アストルフォ：村瀬歩<br />
クロエ：釘宮理恵<br />
ジャン＝ジャック：濱野大輝<br />
<br />
Twitter：<a href="https://twitter.com/vanitas_anime" target="_blank">@vanitas_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/ヴァニタス" target="_blank">#ヴァニタス</a><br />
<br />
<b><a href="https://www.youtube.com/watch?v=IMc5Lzl1bTE" target="_blank">TVアニメ『ヴァニタスの手記』2クール目ショートPV第1弾：ジェヴォーダンの獣 編</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/IMc5Lzl1bTE" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="43">◆<b><a href="https://www.gf-anime.com/" target="_blank">ドールズフロントライン</a></b></h2><p class="preface">
</p><a href="https://www.gf-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/43.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/7(金) 25:00～<br />
BS11：1/7(金) 25:30～<br />
AT-X：1/10(月) 23:30～、ほか<br />
<br />
ABEMA：1/7(金) 25:00～<br />
他配信：1/8(土) 12:00～<br />
(Amazon Prime Video、dアニメストア、Hulu、U-NEXT、アニメ放題、ほか)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>2045年、コーラップス液によって世界中が汚染された地球。<br />
人類は最も原始的な欲求である住居や食料を巡って、全世界を巻き込んだ第三次世界大戦が勃発。<br />
終戦後の世界は、もはや荒廃しきっていた。<br />
文明が滅びる寸前まで追い込まれる一方で、戦争とそれに伴う労働力の不足は機械技術の進歩を促し、その結果人間を模した“人形”が創り出された。<br />
人間への各種サービス提供を目的にしたその見た目は人間の少女と何も変わらないが、人間を遙かに超越した能力を持つ機械生命体。それが“人形”である。<br />
そして技術が更に発展するにつれ、人形にマッチする銃を持たせた「戦術人形」があらゆる軍事組織で採用された。<br />
そうした中、軍需企業「鉄血工造」のAIが突如人類へ反旗を翻し、それを調査・阻止しようとする民間軍事会社「グリフィン」の戦術人形たちとの戦いが幕を開ける。<br />
「グリフィン」に所属する最高級の戦術人形で構成された「AR小隊」のメンバーは、人間の指揮官とともに各地の秩序を取り戻し、人類の安全を守ると同時に、その戦争の背後に潜む巨大な謎を解き明かしていくこととなる。<br />
戦術人形たちが紡ぐ、壮絶な戦いの末に待っているものとは、いったい……。</blockquote><p class="preface">
<br />
原作は<b><a href="https://gf-jp.sunborngame.com/" target="_blank">スマホ向けゲームアプリ</a></b>。<br />
<br />
<b>・スタッフ</b><br />
原作：SUNBORN Network Technology, Mica Team<br />
監督：上田繁<br />
シリーズ構成・脚本：倉田英之<br />
キャラクターデザイン：山田正樹<br />
プロップデザイン：神宮司訓之<br />
背景：グーフィー<br />
美術監督：権瓶岳斗<br />
色彩設計：斎藤麻記<br />
撮影監督：小寺翔太<br />
3DCGディレクター：原一晃<br />
編集：本田優規<br />
音楽：渡邊崇<br />
音響監督：田中亮<br />
アニメーション制作：旭プロダクション<br />
プロデュース：ワーナー ブラザース ジャパン<br />
コピーライト表記：<span translate="no">&#x00A9;SUNBORN Network Technology, Mica Team / GRIFFIN & KRYUGER</span><br />
<br />
<b>・キャスト</b><br />
M4A1：戸松遥<br />
M16A1：山根希美<br />
ST AR-15：加藤英美里<br />
M4 SOPMOD II：田村ゆかり<br />
ジャンシアーヌ：小松未可子<br />
カリーナ：東山奈央<br />
エージェント：生天目仁美<br />
スケアクロウ：奥野香耶<br />
エクスキューショナー：伊藤静<br />
クルーガー：大塚明夫<br />
<br />
Twitter：<a href="https://twitter.com/gf_animePR" target="_blank">@gf_animePR</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/ドールズフロントライン" target="_blank">#ドールズフロントライン</a><br />
<br />
OP：yukaDD(;´∀｀)「BAD CANDY」<br />
ED：TEAM SHACHI「HORIZON」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=uTnmApB8h0Q" target="_blank">アニメ『ドールズフロントライン』PV 第2弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/uTnmApB8h0Q" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="16">◆<b><a href="https://takagi3.me/index.html" target="_blank">からかい上手の高木さん3</a></b></h2><p class="preface">
</p><a href="https://takagi3.me/index.html" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/16.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
MBS・TBS系28局ネット：1/7(金) 25:25～<br />
AT-X：1/8(土) 21:00～、ほか<br />
J:COMテレビ：1/10(月) 25:00～<br />
「<b><a href="https://www.mbs.jp/anime/" target="_blank">スーパーアニメイズム</a></b>」枠<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>とある中学校、隣の席になった女の子・高木さんに、何かとからかわれる男の子・西片。高木さんをからかい返そうと策を練るも、いつも高木さんに見透かされてしまう。<br />
どうにかして、高木さんにひと泡吹かそうと奮闘する日々…。移ろうものは季節だけではなく、西片にも…？一方、優勢とみられた高木さんにも、ゆらぐ事態が―？<br />
スキを見るか、スキを見せるか―ニヤキュン指数100％の「からかいバトル」、いよいよ最終ラウンドに…！？</blockquote><p class="preface">
<br />
原作は「ゲッサン」連載で、既刊16巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4091250157/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4091250157/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/16-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
なお、このアニメ第3期の発表とともに2022年の映画化も発表されている。<br />
<br />
<b>・スタッフ</b><br />
原作：山本崇一朗<br />
監督：赤城博昭<br />
脚本：福田裕子、伊丹あき、加藤還一<br />
キャラクターデザイン：髙野綾<br />
総作画監督：茂木琢次、近藤奈都子、岡昭彦、小野田将人、高野綾<br />
音楽：堤博明<br />
アニメーション制作：シンエイ動画<br />
コピーライト表記：<span translate="no">&#x00A9;2022 山本崇一朗・小学館／からかい上手の高木さん３製作委員会</span><br />
<br />
<b>・キャスト</b><br />
高木さん：高橋李依<br />
西片：梶裕貴<br />
ミナ：小原好美<br />
ユカリ：M・A・O<br />
サナエ：小倉唯<br />
中井：内田雄馬<br />
真野：小岩井ことり<br />
高尾：岡本信彦<br />
木村：落合福嗣<br />
浜口：内山昂輝<br />
北条：悠木碧<br />
田辺先生：田所陽向<br />
<br />
Twitter：<a href="https://twitter.com/takagi3_anime" target="_blank">@takagi3_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/高木さんめ" target="_blank">#高木さんめ</a><br />
<br />
OP：大原ゆい子「まっすぐ」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=N238ads-bdw" target="_blank">TVアニメ『からかい上手の高木さん３』PV第2弾（2022年1月7日（金）放送開始！）</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/N238ads-bdw" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="17">◆<b><a href="https://end-harem-anime.com/" target="_blank">終末のハーレム</a></b></h2><p class="preface">
</p><a href="https://end-harem-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/17.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/7(金) 25:30～<br />
BSフジ：1/7(金) 28:00～<br />
AT-X：1/7(金) 21:30～、ほか<br />
<br />
dアニメストア：1/7(金) 25:30～【地上波同時】<br />
ABEMA：1/7(金) 25:30～【地上波同時】<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>時は近未来――2040年の日本・東京。<br />
ある難病に侵された青年・怜人は幼馴染の絵理沙と再会を誓い、病を治すため“コールドスリープ”することに。<br />
5年後に目を醒ますと、世界は大きな変貌を遂げていた。<br />
MK(Male Killer)ウイルスによって地球上の99.9%の男性が死滅。<br />
地上は5人の男に対して50億の女性が存在する、超ハーレムとなっていた。<br />
MKウイルスへの抵抗力を持つ男性“ナンバーズ”は、わずか5人。<br />
その1人である怜人は、残された女性たちと人類の存続のため“メイティング”(子作り)することを求められる。<br />
パンデミック後の世界に待っていたハーレム生活。<br />
同時に、怜人はナンバーズを巡る世界的な陰謀に巻き込まれていく。<br />
押し寄せる誘惑を乗り越え、世界を救うことはできるのか。</blockquote><p class="preface">
<br />
原作は「少年ジャンプ＋」連載の“近未来エロティックサスペンス”漫画。既刊13巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4088808193/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4088808193/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/17-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
2021年10月から放送された作品だが、「表現の精査が必要」ということで<b><a href="https://end-harem-anime.com/news/index00320000.html" target="_blank">第2話以降の放送・配信が中止</a></b>され、改めて2022年1月開始となった。なお、AT-Xでは「無修正ver.」のほか、Blu-rayに収録される完全新作のセクシーシーン(モザイクあり)を追加した「“超ハーレム”ちょい見せVer.」も放送される。<br />
<br />
<b>・スタッフ</b><br />
原作：LINK・宵野コタロー<br />
監督：信田ユウ<br />
シリーズ構成：高橋龍也<br />
キャラクターデザイン：小関雅<br />
プロップ設定：コレサワシゲユキ(デジタルノイズ)、灯夢(デジタルノイズ)<br />
美術監督：葛琳(スタジオちゅーりっぷ)<br />
美術設定：藤瀬智康(チップチューン)<br />
色彩設計：松山愛子<br />
撮影監督：三上颯太(チップチューン)<br />
編集：萩原うたこ<br />
音響監督：郷文裕貴<br />
音楽：大川茂伸<br />
アニメーション制作スタジオ：Studio五組×AXsiZ<br />
コピーライト表記：<span translate="no">&#x00A9;LINK・宵野コタロー／集英社・終末のハーレム製作委員会</span><br />
<br />
<b>・キャスト</b><br />
水原怜人：市川太一<br />
周防美来：白石晴香<br />
龍造寺朱音：大地葉<br />
山田翠：山根綺<br />
片桐麗亜：渡部恵子<br />
黒田マリア：小坂井祐莉絵<br />
水原まひる：首藤志奈<br />
火野恭司：江口拓也<br />
石動寧々子：石上静香<br />
北山玲奈：天野聡美<br />
土井翔太：浦和希<br />
神谷花蓮：竹達彩奈<br />
羽生柚希：早瀬莉花<br />
柊春歌：道井悠<br />
一条奈都：愛原ありさ<br />
東堂晶：飯田里穂<br />
黒田・レイン・ちふゆ：高森奈津美<br />
<br />
Twitter：<a href="https://twitter.com/harem_official_" target="_blank">@harem_official_</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/終末のハーレム" target="_blank">#終末のハーレム</a><br />
<br />
OP：H-el-ical//「JUST DO IT」<br />
ED：EXiNA「ENDiNG MiRAGE」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=Jr3UJoIKGJM" target="_blank">TVアニメ『終末のハーレム』本PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/Jr3UJoIKGJM" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="18">◆<b><a href="https://cue-animation.jp/" target="_blank">CUE!</a></b></h2><p class="preface">
</p><a href="https://cue-animation.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/18.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
MBS：1/7(金) 25:55～<br />
TBS：1/7(金) 25:55～<br />
BS-TBS：1/7(金) 26:30～<br />
AT-X：1/11(火) 21:00～、ほか<br />
「<b><a href="https://www.mbs.jp/anime/" target="_blank">アニメイズム</a></b>」枠<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>実績も経験もない、できたてホヤホヤの声優事務所『AiRBLUE』。そこに所属するのは個性豊かな声優の卵たち。彼女たちはいまの自分に何ができるのかを考え、それぞれが信じる形で夢へと駆け出していく！<br />
しかし、夢を抱きつつも直面する厳しい現実。苦悩、挫折、葛藤――。いくら練習を重ねても、全員がオーディションに合格することは叶わない。声優を目指す者なら誰もがぶつかる壁を前に、彼女たちはどう立ち向かっていくのか？<br />
新人声優たちによる物語が、いま始まる。</blockquote><p class="preface">
<br />
原作は次世代声優育成ゲームアプリ「<b><a href="https://www.cue-liber.jp/" target="_blank">CUE! -See You Everyday-</a></b>」。<br />
<br />
<b>・スタッフ</b><br />
原作：リベル・エンタテインメント<br />
キャラクター原案：シソ<br />
監督：片貝慎<br />
シリーズ構成：浦畑達彦<br />
キャラクターデザイン：谷口元浩<br />
サブキャラクターデザイン：西村理恵<br />
美術監督：松本浩樹<br />
色彩設計：松森より子<br />
撮影監督：加納篤<br />
編集：近藤勇二<br />
音響監督：ハマノカズゾウ<br />
音楽：中西亮輔<br />
音楽制作：ポニーキャニオン、アップドリーム<br />
アニメーション制作：ゆめ太カンパニー×グラフィニカ<br />
製作：CUE! Animation Project<br />
コピーライト表記：<span translate="no">&#x00A9;CUE! Animation Project</span><br />
<br />
<b>・キャスト</b><br />
六石陽菜：内山悠里菜<br />
鷹取舞花：稗田寧々<br />
鹿野志穂：守屋亨香<br />
月居ほのか：緒方佑奈<br />
恵庭あいり：飯塚麻結<br />
九条柚葉：村上まなつ<br />
夜峰美晴：安齋由香里<br />
神室絢：松田彩希<br />
宮路まほろ：山口愛<br />
日名倉莉子：鶴野有紗<br />
丸山利恵：立花日菜<br />
宇津木聡里：小峯愛未<br />
明神凛音：佐藤舞<br />
遠見鳴：土屋李央<br />
<br />
Twitter：<a href="https://twitter.com/cue_anime" target="_blank">@cue_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/キュー" target="_blank">#キュー</a><br />
<br />
OP：AiRBLUE「スタートライン」<br />
ED：AiRBLUE「はじまりの鐘の音が鳴り響く空」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=yi6zNeKaDbU" target="_blank">TVアニメ『CUE!』PV第1弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/yi6zNeKaDbU" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
<b><a href="https://www.youtube.com/watch?v=A54kpjawEWM" target="_blank">TVアニメ『CUE!』PV第2弾 チームPV Ver.Moon</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/A54kpjawEWM" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
<b><a href="https://www.youtube.com/watch?v=XkdRrYGeznw" target="_blank">TVアニメ『CUE!』PV第3弾 チームPV Ver.Wind</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/XkdRrYGeznw" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
<b><a href="https://www.youtube.com/watch?v=uNj9iGIciBc" target="_blank">TVアニメ『CUE!』PV第4弾 チームPV Ver.Bird</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/uNj9iGIciBc" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="56">◆<b><a href="https://bluev.tv/" target="_blank">龍青神ブルーヴ</a></b></h2><p class="preface">
</p><a href="https://bluev.tv/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/56.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
サンテレビ：1/8(土) 10:00～<br />
全13話<br />
<br />
<b>・作品情報</b><br />
大学生の髙松大地は、父が開発した自然エネルギーから無限の動力を得るシステムを導入したAI搭載型パワードスーツ「ブルーヴスーツ」を身にまとい、「龍青神ブルーヴ」となって、人や自然環境を破壊するロボット達が運営する悪の会社組織と対峙しながら、自然を守るために調和を目指し奮闘する。<br />
<br />
太陽光発電・蓄電池などを手がける総合エネルギー会社・ブルーコンシャスグループが、創業10周年に合わせて推進している環境活動の一環として制作されるオリジナル特撮ドラマ。<br />
<br />
<b>・キャスト</b><br />
髙松大地：直林主浩<br />
日向美空：浜川結琳<br />
アクア：愛子<br />
髙松博士：当山彰一<br />
ゴールドクレー社CEO：Reiji<br />
平社員ロボ シルバ：永尾宗大<br />
平社員ロボ コッパ：錦織聡<br />
高田夏帆<br />
<br />
Twitter：<a href="https://twitter.com/ryuseijin_bluev" target="_blank">@ryuseijin_bluev</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/龍青神ブルーヴ" target="_blank">#龍青神ブルーヴ</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="42">◆<b><a href="https://www.tv-tokyo.co.jp/anime/ninjala/" target="_blank">ニンジャラ</a></b></h2><p class="preface">
</p><a href="https://www.tv-tokyo.co.jp/anime/ninjala/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/42.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ東京系6局ネット：1/8(土) 10:30～<br />
<br />
<b>・作品情報</b><br />
Nintendo Switch向けの<b><a href="https://ninjalathegame.com/jp/" target="_blank">対戦ニンジャガムアクションゲーム</a></b>をアニメ化。<br />
<br />
<b>・スタッフ</b><br />
原案：森下一喜<br />
監督：神戸守<br />
シリーズ構成：藤田伸三<br />
キャラクターデザイン：鈴木敦<br />
音響監督：浦上靖之<br />
音楽：小畑貴裕<br />
アニメーション制作：オー・エル・エム<br />
コピーライト表記：<span translate="no">&#x00A9; GungHo Online Entertainment・TV Tokyo</span><br />
<br />
<b>・キャスト</b><br />
バートン：櫻井孝宏<br />
ベレッカ：鬼頭明里<br />
ロン：安元洋貴<br />
ジェーン：皆川純子<br />
バーン：小林由美子<br />
ルーシー：悠木碧<br />
カッペイ：村瀬歩<br />
エマ：潘めぐみ<br />
ガムッチ：井上ほの花<br />
<br />
Twitter：<a href="https://twitter.com/Ninjala_JP" target="_blank">@Ninjala_JP</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/ニンジャラ" target="_blank">#ニンジャラ</a><br />
<br />
OP：きゃりーぱみゅぱみゅ「メイビーベイビー」<br />
ED：ウォルピスカーター「ニンジャライクニンジャ」<br />
<br />
</p><hr></p><p class="preface"><h2 id="44">◆<b><a href="https://shikkakumon.com/" target="_blank">失格紋の最強賢者</a></b></h2><p class="preface">
</p><a href="https://shikkakumon.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/44.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/8(土) 22:00～<br />
BS11：1/8(土) 22:00～<br />
サンテレビ：1/8(土) 23:00～(第2話以降は22:30～)<br />
AT-X：1/11(火) 23:00～、ほか<br />
<br />
ABEMA：1/8(土) 22:00～<br />
他配信：1/12(水) 24:00～<br />
(dアニメストア、Netflix、Amazon Prime Video、アニメカ、U-NEXT、アニメ放題、GYAO!、FOD、バンダイチャンネル、Hulu、TELASA、J:COMオンデマンド、みるプラス、auスマートパスプレミアム、Google Play、ひかりTV、ニコニコチャンネル、ビデオマーケット、マンガUP!)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>世界最強の魔法使いと謳われながらも、生まれ持った紋章の性能に限界を感じていた【賢者】ガイアス。<br />
その彼が己の紋章を変えるために取った手段――それは転生によって新たな体を得ること！<br />
彼は遥か未来の世界に転生し、求めていた「魔法戦闘に最適な紋章」と、マティアスという名を手に入れた。<br />
しかし、その紋章はこの時代ではなぜか「失格紋」と呼ばれていた……！<br />
時を経た今世では、魔法が衰退し低レベルな魔法理論が跋扈してしまっていたのだ。<br />
魔法戦闘最強の「失格紋」と、賢者の知恵を併せ持つ少年マティアスは、世界の常識を次々と打ち壊していく！</blockquote><p class="preface">
<br />
原作は「小説家になろう」に連載されていた小説で、2016年から書籍版がGAノベルで刊行されている。最新14巻が2022年1月14日発売予定。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4797392363/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4797392363/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/44-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：進行諸島<br />
キャラクター原案：風花風花<br />
監督：秋田谷典昭<br />
シリーズ構成：内田裕基<br />
キャラクターデザイン：大沢美奈<br />
アニメーション制作：J.C.STAFF<br />
<br />
<b>・キャスト</b><br />
マティアス＝ヒルデスハイマー：玉城仁菜<br />
ルリイ＝アーベントロート：鈴代紗弓<br />
アルマ＝レプシウス：白石晴香<br />
イリス：井澤詩織<br />
<br />
Twitter：<a href="https://twitter.com/shikkakumon_PR" target="_blank">@shikkakumon_PR</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/失格紋" target="_blank">#失格紋</a><br />
<br />
OP：fripSide「Leap of faith」<br />
ED：中島由貴「Day of Bright Sunshine」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=o34GOPt0-78" target="_blank">TVアニメ『失格紋の最強賢者』第1弾PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/o34GOPt0-78" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="19">◆<b><a href="https://bisquedoll-anime.com/" target="_blank">その着せ替え人形は恋をする</a></b></h2><p class="preface">
</p><a href="https://bisquedoll-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/19.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/8(土) 24:00～<br />
とちぎテレビ：1/8(土) 24:00～<br />
群馬テレビ：1/8(土) 24:00～<br />
BS11：1/8(土) 24:00～<br />
AT-X：1/9(日) 22:00～、ほか<br />
読売テレビ：1/10(月) 26:29～<br />
メ～テレ：1/13(木) 25:56～<br />
<br />
Amazon Prime Video：1/8(土) 24:30～<br />
dアニメストア：1/8(土) 24:30～<br />
他配信：1/11(火)～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>雛人形の顔を作る、「頭師」を目指す男子高校生・五条新菜。<br />
真面目で雛人形作りに一途な反面、同世代の流行には疎く、中々クラスに馴染めずにいる。<br />
そんな新菜にとって、いつもクラスの輪の中心にいる人気者・喜多川海夢はまるで別世界の住人。<br />
けれどある日、思わぬことをきっかけに、海夢と秘密を共有することになって……！？<br />
決して交わるはずのなかった２人の世界が、動き出す――！</blockquote><p class="preface">
<br />
原作は「ヤングガンガン」連載中で、既刊8巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4757559208/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4757559208/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/19-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：福田晋一<br />
監督：篠原啓輔<br />
シリーズ構成・脚本：冨田頼子<br />
副監督：平峯義大<br />
キャラクターデザイン・総作画監督：石田一将<br />
総作画監督：小林真平、川妻智美、山崎淳<br />
メインアニメーター：髙橋尚矢<br />
衣装デザイン：西原恵利香<br />
色彩設計：山口舞<br />
美術設定：根本洋行<br />
特殊効果：入佐芽詠美<br />
撮影監督：金森つばさ<br />
テクニカルディレクター：佐久間悠也<br />
ＣＧディレクター：宮地克明<br />
編集：平木大輔<br />
音楽：中塚武<br />
音響監督：藤田亜紀子<br />
音響効果：野崎博樹、小林亜依里<br />
制作：CloverWorks<br />
製作：「着せ恋」製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;福田晋一/SQUARE ENIX・「着せ恋」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
喜多川海夢：直田姫奈<br />
五条新菜：石毛翔弥<br />
乾紗寿叶：種﨑敦美<br />
乾心寿：羊宮妃那<br />
五条薫：斧アツシ<br />
<br />
Twitter：<a href="https://twitter.com/kisekoi_anime" target="_blank">@kisekoi_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/着せ恋" target="_blank">#着せ恋</a><br />
<br />
OP：スピラ・スピカ「燦々デイズ」<br />
ED：あかせあかり「恋ノ行方」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=tFKDKd8z-NU" target="_blank">TVアニメ「その着せ替え人形は恋をする」第2弾PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/tFKDKd8z-NU" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="20">◆<b><a href="https://akebi-chan.jp/" target="_blank">明日(あけび)ちゃんのセーラー服</a></b></h2><p class="preface">
</p><a href="https://akebi-chan.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/20.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/8(土) 24:30～<br />
群馬テレビ：1/8(土) 24:30～<br />
とちぎテレビ：1/8(土) 24:30～<br />
BS11：1/8(土) 24:30～<br />
MBS：1/8(土) 26:08～<br />
BS朝日：1/9(日) 23:00～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>舞台は、田舎の名門女子中学・私立蠟梅学園。<br />
あるきっかけから、この学園のセーラー服を着ることが「夢」だった少女、明日小路。<br />
念願叶い、ドキドキで入学式当日を迎えるが―<br />
「私はセーラー服に決めました」<br />
決意を胸に、夢の中学生ライフが始まる♪<br />
クラスメート、給食、部活動…<br />
“初めて”だらけの毎日を、小路は全力で駆け抜ける！<br />
少女たちの、キラキラ輝く青春日記。<br />
「友達いっぱいできるかな？」</blockquote><p class="preface">
<br />
原作は「となりのヤングジャンプ」連載で、既刊9巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4088906713/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4088906713/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/20-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：博<br />
監督：黒木美幸<br />
キャラクターデザイン：河野恵美<br />
シリーズ構成、脚本：山崎莉乃<br />
キャラクターデザイン：河野恵美<br />
サブキャラクターデザイン：川上大志、安野将人<br />
総作画監督：河野恵美、川上大志、安野将人<br />
美術設定：塩澤良憲<br />
美術監督：薄井久代、守安靖尚<br />
色彩設計：横田明日香<br />
撮影監督：川下裕樹(MADBOX)<br />
3Dディレクター：宮原洋平<br />
キャラクターレタッチ：カプセル<br />
編集：齋藤朱里(三嶋編集室)<br />
音楽：うたたね歌菜<br />
音響監督：濱野高年<br />
制作：CloverWorks<br />
コピーライト表記：<span translate="no">&#x00A9;博／集英社・「明日ちゃんのセーラー服」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
明日小路：村上まなつ<br />
木崎江利花：雨宮天<br />
兎原透子：鬼頭明里<br />
古城智乃：若山詩音<br />
谷川景：関根明良<br />
鷲尾瞳：石上静香<br />
水上りり：石川由依<br />
平岩蛍：麻倉もも<br />
四条璃生奈：田所あずさ<br />
神黙根子：伊藤美来<br />
龍守逢：伊瀬茉莉也<br />
峠口鮎美：三上枝織<br />
蛇森生静：神戸光歩<br />
苗代靖子：本渡楓<br />
戸鹿野舞衣：白石晴香<br />
大熊実：小原好美<br />
<br />
Twitter：<a href="https://twitter.com/akebi_chan" target="_blank">@akebi_chan</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/明日ちゃん" target="_blank">#明日ちゃん</a><br />
<br />
<b><a href="https://www.youtube.com/watch?v=q2kTPtVWORs" target="_blank">TVアニメ「明日ちゃんのセーラー服」PV第1弾 | 2022年1月放送開始</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/q2kTPtVWORs" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="21">◆<b><a href="https://genkoku-anime.com/" target="_blank">現実主義勇者の王国再建記 第二部</a></b></h2><p class="preface">
</p><a href="https://genkoku-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/21.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/8(土) 25:30～<br />
BS11：1/8(土) 25:30～<br />
HTB北海道テレビ：1/12(水) 25:50～<br />
アニマックス：1/22(土) 19:30～<br />
<br />
FOD：1/8(土) 25:30～【地上波同時】<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>異世界に召喚された相馬一也は、エルフリーデン王国の王位を譲られて以来、辣腕を発揮し、王国の政治・財政を立て直していった。<br />
そして、三公の反乱とアミドニア公国の侵略を鎮圧したソーマは今、新たな問題に直面していた。<br />
アミドニアの首都ヴァンを失い敗走した公太子ユリウスが復権を狙い、人類宣言の盟主たるグラン・ケイオス帝国に助けを求めたのである。<br />
始まる女皇マリアの妹ジャンヌ･ユーフォリアを代表とする帝国との交渉。<br />
それはエルフリーデン、アミドニアだけではなく、大陸全体の運命をも左右するものとなろうとしていた――。</blockquote><p class="preface">
<br />
原作は「小説家になろう」「pixiv」連載のウェブ小説で、オーバーラップ文庫から書籍版が刊行されている。既刊15巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/486554111X/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/486554111X/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/21-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：どぜう丸<br />
キャラクター原案：冬ゆき<br />
監督：渡部高志<br />
脚本：雑破業、大野木寛<br />
キャラクターデザイン：大塚舞<br />
音楽：立山秋航<br />
音響監督：明田川仁<br />
音響効果：小山恭正<br />
音響制作：マジックカプセル<br />
音楽制作：キングレコード<br />
プロデュース：WOWMAX<br />
アニメーション制作：J.C.STAFF<br />
コピーライト表記：<span translate="no">&#x00A9; どぜう丸・オーバーラップ／現国製作委員会</span><br />
<br />
<b>・キャスト</b><br />
ソーマ・カズヤ：小林裕介<br />
リーシア・エルフリーデン：水瀬いのり<br />
アイーシャ・ウドガルド：長谷川育美<br />
ジュナ・ドーマ：上田麗奈<br />
ハクヤ・クオンミン：興津和幸<br />
トモエ・イヌイ：佳原萌枝<br />
ポンチョ・パナコッタ：水中雅章<br />
カルラ・バルガス：愛美<br />
ロロア・アミドニア：M・A・O<br />
マリア・ユーフォリア：金元寿子<br />
ジャンヌ・ユーフォリア：石川由依<br />
<br />
Twitter：<a href="https://twitter.com/genkoku_info" target="_blank">@genkoku_info</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/現国アニメ" target="_blank">#現国アニメ</a><br />
<br />
OP：水瀬いのり「REAL-EYES」<br />
ED：愛美「LIGHTS」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=onTxCpESx0Y" target="_blank">【第二部-本PV】TVアニメ「現実主義勇者の王国再建記」</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/onTxCpESx0Y" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="22">◆<b><a href="https://kuroitsusan-anime.com/" target="_blank">怪人開発部の黒井津さん</a></b></h2><p class="preface">
</p><a href="https://kuroitsusan-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/22.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
ABC・テレビ朝日系列全国24局ネット：1/8(土) 26:00～<br />
BS11：1/9(日) 24:00～<br />
AT-X：1/12(水) 23:00～、ほか<br />
「<b><a href="https://www.asahi.co.jp/anime/animazing/" target="_blank">ANiMAZiNG!!!</a></b>」枠<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>正義と悪がぶつかり合う世界、人知れず戦う人々がいた。<br />
それは幾度となく正義のヒーローに倒されていく定めの悪の怪人を作り出す開発部で働く人々である。<br />
地下深くに存在する秘密結社アガスティアの研究室で、彼らは実験や研究だけではなく、開発の予算取りからスケジュール管理、そして幹部の決裁をもらうためのプレゼンという名の戦いの日々を過ごしているのだった。<br />
アガスティアでヒラ研究員として働く黒井津燈香は、佐田巻博士と共に数々の難関をかいくぐりヒーローを打倒する怪人を見事作り出すことができるのか!?<br />
その宿敵として黒井津たちの前に立ちはだかるは無名で無敗の変身ヒーロー・剣神ブレイダー！<br />
この物語はヒーローを倒す怪人を開発すべく奮戦する、もうひとつのプロフェッショナル達の日々の記録である。</blockquote><p class="preface">
<br />
原作は「COMICメテオ」連載中で、最新4巻が2022年2月12日発売予定。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4866751339/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4866751339/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/22-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：水崎弘明<br />
監督：斎藤久<br />
シリーズ構成：高山カツヒコ<br />
キャラクターデザイン：森前和也<br />
サブキャラクターデザイン：新谷真昼<br />
ヒーローデザイン：山根理宏、ことぶきつかさ<br />
怪人デザイン：森木靖泰<br />
衣装デザイン：監物ケビン雄太<br />
レジェンド＆ローカルヒーロープロデュース：鈴村展弘<br />
ヒーローキャプテン：まさひろ山根<br />
ローカルヒーローキャプテン：鷲北恭太<br />
美術監督：甲斐政俊<br />
美術設定：赤樹壱磨<br />
色彩設計：ながさか暁<br />
撮影監督：今泉秀樹<br />
3DCG：渡辺哲也<br />
編集：木村祥明<br />
音響監督：飯田里樹<br />
音楽：manzo<br />
音楽制作：CREST<br />
アニメーション制作：Quad<br />
コピーライト表記：<span translate="no">&#x00A9;水崎弘明・COMICメテオ/「怪人開発部の黒井津さん」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
黒井津燈香：前田佳織里<br />
ウルフ・ベート：天野聡美<br />
佐田巻博士：梅原裕一郎<br />
メギストス：稲田徹<br />
アカシック：M・A・O<br />
佐田巻健司(剣神ブレイダー)：寺島拓篤<br />
<br />
Twitter：<a href="https://twitter.com/kuroitsusan" target="_blank">@kuroitsusan</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/黒井津さん" target="_blank">#黒井津さん</a><br />
<br />
OP：AXXX1S「Special Force」<br />
ED：メイビーME「曖昧あいでんてぃてぃ」<br />
<br />
メイビーME「Destiny」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=rbZ2R8qYXbU" target="_blank">TVアニメ『怪人開発部の黒井津さん』ティザーPV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/rbZ2R8qYXbU" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="45">◆<b><a href="https://link-click.jp/" target="_blank">時光代理人 -LINK CLICK-</a></b></h2><p class="preface">
</p><a href="https://link-click.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/45.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/9(日) 21:30～<br />
BS11：1/9(日) 22:30～<br />
ホームドラマチャンネル：2/26(土) 21:30～(2話連続)<br />
<br />
Amazon Prime Video：1/9(日)～【オリジナル版最速配信】<br />
ABEMA：1/9(日)～【放送版最速配信】<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>繁華街の一角に佇む「時光写真館」。<br />
そのさびれたドアの奥には、特殊な能力を持った２人の男がいた——。<br />
写真館を経営するのはトキ(程小時)とヒカル(陸光)。<br />
トキの幼馴染・リン(喬苓)を通じて顧客から舞い込む依頼を遂行すべく、「撮影者の意識にリンクし、写真の世界に入ることができる能力」を持つトキと、「その写真の撮影後12時間の出来事を把握できる能力」を持つヒカルはコンビを組み、過去を引きずるクライアントからの依頼を解決していく。<br />
『絶対に過去の改変をしてはならない』ルールのもと依頼を遂行していた二人だが、正義感の強いトキはつい過去に干渉してしまい、その行動はやがて少しずつ未来を変えていく——。</blockquote><p class="preface">
<br />
映画「<b><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B07DV95WR7/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B07DV95WR7/gigazine-22');">詩季織々</a></b>」や2021年夏に放送された「<b><a href="https://gigazine.net/news/20210625-anime-2021summer/#35" target="_blank">天官賜福</a></b>」などを手がけた李豪凌監督によるオリジナルアニメ。2021年4月からbilibili動画で配信され、配信開始から4カ月で1億6000万回再生された。<br />
<br />
<b>・スタッフ</b><br />
監督・脚本：李豪凌(リー・ハオリン)<br />
キャラクターデザイン原案：INPLICK<br />
キャラクターデザイン：LAN、黄思萌(ホワン・スーモン)、熊丹(ション・ダン)<br />
総演出・総作画監督：LAN<br />
美術監督：丹治匠、朝見知弥、朱立朴(ジュー・リープー)<br />
撮影監督：山条裕香<br />
音楽：天門、yuma yamaguchi、av4ln<br />
アニメーション制作：瀾映画<br />
日本版製作：株式会社ソニー・ミュージックソリューションズ、株式会社アニプレックス<br />
コピーライト表記：<span translate="no">&#x00A9;bilibili/BeDream</span><br />
<br />
<b>・キャスト</b><br />
トキ／程小時(チョン・シャオシー)：豊永利行<br />
ヒカル／陸光(ルー・グアン)：櫻井孝宏<br />
リン／喬苓(チャオ・リン)：古賀葵<br />
<br />
Twitter：<a href="https://twitter.com/linkclick_anime" target="_blank">@linkclick_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/時光代理人" target="_blank">#時光代理人</a><br />
<br />
OP：Gen Kakon「Dive Back In Time」<br />
ED：EAERAN「OverThink」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=0NhDPvc08n0" target="_blank">TVアニメ「時光代理人 -LINK CLICK-」本PV | 2022年1月9日(日)よりTOKYO MX・BS11にて放送開始！</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/0NhDPvc08n0" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="23">◆<b><a href="https://rustedarmors-anime.com/" target="_blank">錆色のアーマ-黎明-</a></b></h2><p class="preface">
</p><a href="https://rustedarmors-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/23.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX1：1/9(日) 22:00～<br />
AT-X：1/9(日) 23:30～、ほか<br />
サンテレビ：1/9(日) 24:30～<br />
テレビ和歌山：1/9(日) 24:00～<br />
BSフジ：1/11(火) 24:00～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>長きに渡る戦乱により、室町幕府が弱体化する中世は血で血を洗う混迷の戦国時代へと突入した。<br />
山海の秘境、紀ノ國。<br />
この深い森の奥に、八咫烏を旗印にした鉄砲集団がいた。<br />
傭兵稼業を生業とし、戦場を駆ける、その名は『雑賀衆』。<br />
雑賀衆の頭として名跡を受け継いだのは、異国から流れ着いた『孫一』だった。<br />
一方、欧州列強からの侵略の気配をいち早く察知し、日ノ本を守るために奔走する男、『三郎』。相まみえるはずのない二人の運命は、異国から襲来した侵略者の刃によって交錯する。<br />
戦国の世に集いし男たちの信念と正義を貫いた物語が、今、開幕する──！</blockquote><p class="preface">
<br />
原作は2017年から展開されている舞台を中心にしたメディアミックスプロジェクト。アニメや漫画をもとにした「2.5次元舞台」があるように、本作は当初から舞台をもとにアニメなどに展開していく「逆2.5次元」を想定していた。<br />
<br />
孫市役の佐藤大樹や織田三郎信長役の増田俊樹など、キャストは多くが舞台版から引き続きの担当。<br />
<br />
<b>・スタッフ</b><br />
原作：「錆色のアーマ」プロジェクト<br />
監督：河原真明<br />
シリーズ構成・脚本：江嵜大兄<br />
キャラクターデザイン：小川さくら<br />
美術監督・美術デザイン：小林友理<br />
編集：上野勇輔<br />
CGディレクター：柱谷佳寿、梅野大輔<br />
音響監督：村松久進(Ally Studio)<br />
音楽：井手コウジ、MAYDENFIELD<br />
音楽制作：Spectrum<br />
アニメーション制作：Kigumi<br />
アニメーション制作協力：JCTV/Studio51<br />
コピーライト表記：<span translate="no">&#x00A9;「錆色のアーマ」プロジェクト / &#x00A9; アニメ「錆色のアーマ」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
孫一：佐藤大樹(EXILE / FANTASTICS from EXILE TRIBE)<br />
織田三郎信長：増田俊樹<br />
鶴首：荒木健太朗<br />
蛍火：永田崇人<br />
黒氷：平田裕一郎<br />
木偶：spi<br />
アゲハ：神里優希<br />
不如帰：崎山つばさ<br />
ルシオ・コルテス：佐藤流司<br />
<br />
Twitter：<a href="https://twitter.com/armors_anime" target="_blank">@armors_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/錆色のアーマ" target="_blank">#錆色のアーマ</a><br />
<br />
OP：孫一(CV.佐藤大樹)、織田三郎信長(CV.増田俊樹)、ルシオ(CV.佐藤流司)「Faith」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=pf6hEVSNeL8" target="_blank">アニメ「錆色のアーマ-黎明-」PV第2弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/pf6hEVSNeL8" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="24">◆<b><a href="https://baraou-anime.com/" target="_blank">薔薇王の葬列</a></b></h2><p class="preface">
</p><a href="https://baraou-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/24.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/9(日) 22:30～<br />
サンテレビ：1/9(日) 23:30～<br />
KBS京都：1/9(日) 23:30～<br />
BS11：1/11(火) 24:00～<br />
2クール連続放送<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>中世イングランド。<br />
ヨーク家とランカスター家が王位争奪を繰り返す薔薇戦争時代。<br />
ヨーク家の三男として生まれたリチャードは、母からは「悪魔の子」と疎まれる一方、同じ名を持つ父からは真っ直ぐな愛情を受けて育っていた。<br />
リチャードの願いは、この世の光である父・ヨーク公爵が王位に就くこと。<br />
だがリチャードの純粋な願いは、イングランドに戦乱の嵐を招くことになる。<br />
さらにリチャードは、男女２つの性を持って生まれたという秘密があった。<br />
誰にも明かせぬ秘密を胸に秘めたまま、リチャードもまた戦いの渦中に巻き込まれていく。<br />
そこで待つのは、彼を愛する人たちとのかけがえのない時間と決定的な別離。<br />
痛ましくも美しい邂逅と別離が、「悪」の道へとリチャードを誘っていく――。</blockquote><p class="preface">
<br />
原作は「月刊プリンセス」連載で、既刊16巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4253271812/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4253271812/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/24-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：菅野文<br />
原案：ウィリアム・シェイクスピア『ヘンリー六世』『リチャード三世』<br />
監督：鈴木健太郎<br />
シリーズ構成・脚本：内田裕基<br />
キャラクターデザイン：橋詰力<br />
美術監督：泉健太郎<br />
色彩設計：店橋真弓<br />
撮影監督：高橋昭裕<br />
音響監督：岩浪美和<br />
音楽：大谷幸<br />
音楽制作：ランティス<br />
アニメーション制作：J.C.STAFF<br />
コピーライト表記：<span translate="no">&#x00A9;菅野文（秋田書店）／薔薇王の葬列製作委員会</span><br />
<br />
<b>・キャスト</b><br />
リチャード：斎賀みつき<br />
ヘンリー：緑川光<br />
ヨーク公爵リチャード：速水奨<br />
エドワード：鳥海浩輔<br />
ジョージ：内匠靖明<br />
ウォリック伯爵：三上哲<br />
ケイツビー：日野聡<br />
マーガレット王妃：大原さやか<br />
エドワード王太子：天﨑滉平<br />
ナレーション：大塚芳忠<br />
<br />
Twitter：<a href="https://twitter.com/baraou_anime" target="_blank">@baraou_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/薔薇王の葬列" target="_blank">#薔薇王の葬列</a><br />
<br />
OP：古川慎「我、薔薇に淫す」<br />
ED：ZAQ「悪夢」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=iiU4qanavsI" target="_blank">TVアニメ「薔薇王の葬列」番宣PV【2022年1月9日(日)～連続2クール放送】</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/iiU4qanavsI" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="41">◆<b><a href="https://futsalboys.com/anime" target="_blank">フットサルボーイズ!!!!!</a></b></h2><p class="preface">
</p><a href="https://futsalboys.com/anime" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/41.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/9(日) 23:00～<br />
BS11：1/11(火) 24:30～<br />
MBS：1/11(火) 26:30～<br />
<br />
ABEMA：1/9(土) 23:00～【地上波同時】<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>フットサルが世界的ブームとなって十数年――。<br />
U-18ワールドカップ決勝戦を見た大和晴は、日本代表選手・天王寺刻成に強い憧れを抱く。<br />
そして、意気揚々と入部した恒陽学園高校フットサル部で、仲間にパスを出さない孤高プレイヤーの榊星一郎や、フットサルを諦めかけた過去を持つ、月丘柊依たちと出会う。<br />
それぞれの過去を抱えて、それでも今、フットサルがしたい。<br />
もろく揺れながらも、熱をまとって駆け抜ける――<br />
男子高校生たちの全力の“瞬間”が、はじまる。</blockquote><p class="preface">
<br />
バンダイナムコアーツ、バンダイナムコエンターテインメント、ディオメディアによるメディアミックスプロジェクトで、アニメのほか、アプリゲームとイベント展開が予定されている。<br />
<br />
<b>・スタッフ</b><br />
原作：鞠田マオ<br />
監督：ひいろゆきな<br />
シリーズ構成：米村正二<br />
潤色：東環妃<br />
キャラクターデザイン原案：河下水希、雪広うたこ、たなかマルメロ、潤宮るか、沙汰、シラノ、星野リリィ<br />
キャラクターデザイン：石川智美<br />
サブキャラクターデザイン：鶴田眸<br />
美術監督：魏斯曼<br />
美術設定：高橋麻穂、海津利子<br />
色彩設計：野地弘納<br />
撮影監督：伊藤康行<br />
3Dディレクター：山本祐希江<br />
編集：齋藤朱里<br />
音響監督：納谷僚介<br />
音楽：R・O・N<br />
製作：バンダイナムコアーツ、バンダイナムコエンターテインメント、ディオメディア<br />
コピーライト表記：<span translate="no">&#x00A9;FUTSALBOYS!!!!! PROJECT</span><br />
<br />
<b>・キャスト</b><br />
＜恒陽学園高校＞<br />
大和晴：高良崚太<br />
榊星一郎：石森周斗<br />
月丘柊依：吉原康平<br />
幸永椿貴：山口諒太郎<br />
南雲竜：古田一紀<br />
天門泰雅：坂井易直<br />
宇佐美智：山下誠一郎<br />
<br />
＜アーダルベルト学院高等部＞<br />
樫良木ルイ：峯田大夢<br />
結城心：新井雄也<br />
久我巧生：村上聡<br />
花山院快斗：馬場惇平<br />
京極聖：宮瀬尚也<br />
二葉ともえ：山本智哉<br />
<br />
＜桃実高校＞<br />
昂守希：佐久間貴生<br />
十河夏輝：千葉瑞己<br />
相庭京介：浦和希<br />
花村理央：多田啓太<br />
鞍馬雪丸：上村源<br />
<br />
＜天ノ川学園高等学校＞<br />
桐生蓮：TAKA(CUBERS)<br />
白河瞬：岡延明<br />
橘藤吾：大海将一郎<br />
今園彩人：森永彩斗<br />
水無瀬亜佐：菊池勇成<br />
秋月奏夜：津田拓也<br />
<br />
＜皇花山学園高等部＞<br />
朝比奈碧：奥山敬人<br />
天王寺刻成：川島慶嗣<br />
世良龍太朗：下川草介<br />
水無瀬涼佑：石井孝英<br />
ガルシア・エミリオ：小塚亮輔<br />
蓮美朔：木津つばさ<br />
<br />
ハリケーン児玉：福山潤<br />
<br />
Twitter：<a href="https://twitter.com/futsal_boys" target="_blank">@futsal_boys</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/サルボ" target="_blank">#サルボ</a><br />
<br />
OP：佐久間貴生「BRAVE MAKER」<br />
ED：STEREO DIVE FOUNDATION<br />
<br />
<b><a href="https://www.youtube.com/watch?v=QUshV0Q2XQ0" target="_blank">【TVアニメ】「フットサルボーイズ!!!!!」第1弾PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/QUshV0Q2XQ0" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="25">◆<b><a href="https://sasamiya.com/" target="_blank">佐々木と宮野</a></b></h2><p class="preface">
</p><a href="https://sasamiya.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/25.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/9(日) 24:30～<br />
KBS京都：1/9(日) 24:45～<br />
サンテレビ：1/9(日) 25:00～<br />
テレビ愛知：1/9(日) 25:35～<br />
BS日テレ：1/10(月) 24:00～<br />
AT-X：1/11(火) 22:00～、ほか<br />
<br />
U-NEXT：1/9(日) 24:00～【地上波先行・最速】<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>女顔がコンプレックスな腐男子の宮野由美は、ある夏の日、校内で喧嘩の場に遭遇してしまう。<br />
勇気を出して止めに入ろうとしたとき、宮野の肩を押し留め、代わりに向かってくれたのはちょっとだけ不良な先輩・佐々木秀鳴だった。<br />
それ以来、なぜか佐々木に気に入られてしまう宮野。<br />
あろうことか「好きなマンガを貸してくれ」と言われ――!?<br />
そして佐々木は、瞳を輝かせてBLを語る宮野に、少しずつ惹かれていく…。</blockquote><p class="preface">
<br />
原作はpixivで「<b><a href="https://www.pixiv.net/artworks/48584088" target="_blank">【創作BL】佐々木と宮野のちょっとした話。</a></b>」としてスタートした連作漫画。「<b><a href="https://comic.pixiv.net/magazines/88" target="_blank">ジーンピクシブ</a></b>」連載中で、既刊7巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4040684710/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4040684710/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/25-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：春園ショウ<br />
監督：石平信司<br />
助監督：上野壮大<br />
シリーズ構成：中村能子<br />
キャラクターデザイン：藤井まき<br />
美術監督：黛昌樹<br />
色彩設計：桂木今里<br />
撮影監督：近藤慎与<br />
編集：白石あかね<br />
音響監督：はたしょう二<br />
音響効果：出雲範子<br />
音響制作：スタジオ・ドンファン<br />
音楽：澁江夏奈<br />
音楽制作：日本コロムビア<br />
アニメーション制作：スタジオディーン<br />
製作：「佐々木と宮野」製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;2022 春園ショウ／ＫＡＤＯＫＡＷＡ／「佐々木と宮野」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
佐々木秀鳴：白井悠介<br />
宮野由美：斉藤壮馬<br />
平野大河：松岡禎丞<br />
小笠原次郎：小野友樹<br />
半澤雅人：内田雄馬<br />
暮沢丞：新井良平<br />
田代権三郎：市来光弘<br />
<br />
Twitter：<a href="https://twitter.com/sasamiya_anime" target="_blank">@sasamiya_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/佐々木と宮野" target="_blank">#佐々木と宮野</a><br />
<br />
OP：ミラクルチンパンジー「瞬き」<br />
ED：佐々木秀鳴(CV:白井悠介)、宮野由美(CV:斉藤壮馬)「いちごサンセット」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=cyy5Hx_rFTE" target="_blank">2022年1月放送開始！TVアニメ「佐々木と宮野」本PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/cyy5Hx_rFTE" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="26">◆<b><a href="https://shingeki.tv/final/" target="_blank">進撃の巨人 The Final Season Part 2</a></b></h2><p class="preface">
</p><a href="https://shingeki.tv/final/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/26.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
NHK総合：1/9(日) 24:05～<br />
<br />
配信：1/10(月) 12:00～<br />
(dTV、dアニメストア、GYAO!、Netflix、TELASA、ひかりTV、U-NEXT、Amazon Prime Video)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>「その巨人はいついかなる時代においても、自由を求めて進み続けた。自由のために戦った。名は――進撃の巨人」<br />
ついに明かされた壁の外の真実と、巨人の正体。<br />
ここに至るまで、人類はあまりにも大きすぎる犠牲を払っていた。<br />
それでもなお、彼らは進み続けなければならない。<br />
壁の外にある海を、自由の象徴を、まだその目で見ていないのだから。<br />
―やがて時は流れ、一度目の「超大型巨人」襲来から６年。<br />
調査兵団はウォール・マリア外への壁外調査を敢行する。<br />
「壁の向こうには海があって、海の向こうには自由がある。ずっとそう信じてた……」<br />
壁の中の人類が、初めて辿り着いた海。<br />
果てしなく広がる水平線の先にあるのは自由か、それとも……？<br />
エレン・イェーガーの物語は、新たな局面を迎える。</blockquote><p class="preface">
<br />
原作は「別冊少年マガジン」に2009年から2021年にかけて連載された人気漫画。全34巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4065234174/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4065234174/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/26-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
2013年春以降、4期にわたってテレビアニメ化されていて、この第4期(The Final Season)のPart 2で完結となる。<br />
<br />
<b>・スタッフ</b><br />
原作：諫山創<br />
監督：林祐一郎<br />
シリーズ構成：瀬古浩司<br />
キャラクターデザイン：岸友洋<br />
総作画監督：新沼大祐、秋田学<br />
演出チーフ：宍戸淳<br />
エフェクト作画監督：酒井智史、古俣太一<br />
色彩設計：大西慈<br />
美術監督：小倉一男<br />
画面設計：淡輪雄介<br />
3DCG監督：奥納基、池田昴<br />
撮影監督：浅川茂輝<br />
編集：吉武将人<br />
音響監督：三間雅文<br />
音楽：KOHTA YAMAMOTO、澤野弘之<br />
音響効果：山谷尚人(サウンドボックス)<br />
音響制作：テクノサウンド<br />
制作：MAPPA<br />
コピーライト表記：<span translate="no">&#x00A9;諫山創・講談社／「進撃の巨人」The Final Season製作委員会</span><br />
<br />
<b>・キャスト</b><br />
エレン・イェーガー：梶裕貴<br />
ミカサ・アッカーマン：石川由依<br />
アルミン・アルレルト：井上麻里奈<br />
コニー・スプリンガー：下野紘<br />
ヒストリア・レイス：三上枝織<br />
ジャン・キルシュタイン：谷山紀章<br />
ライナー・ブラウン：細谷佳正<br />
ハンジ・ゾエ：朴璐美<br />
リヴァイ・アッカーマン：神谷浩史<br />
ジーク・イェーガー：子安武人<br />
ファルコ・グライス：花江夏樹<br />
ガビ・ブラウン：佐倉綾音<br />
ピーク・フィンガー：沼倉愛美<br />
ポルコ・ガリアード：増田俊樹<br />
コルト・グライス：松風雅也<br />
<br />
Twitter：<a href="https://twitter.com/anime_shingeki" target="_blank">@anime_shingeki</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/shingeki" target="_blank">#shingeki</a><br />
<br />
<b><a href="https://www.youtube.com/watch?v=jRVKBrJ7X-o" target="_blank">TVアニメ「進撃の巨人」The Final Season Part 2 PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/jRVKBrJ7X-o" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="48">◆<b><a href="https://www.tv-tokyo.co.jp/anime/oadekinai/" target="_blank">オンエアできない！</a></b></h2><p class="preface">
</p><a href="https://www.tv-tokyo.co.jp/anime/oadekinai/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/48.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
BSテレ東：1/9(日) 25:05～<br />
テレビ東京：1/10(月) 26:30～<br />
AT-X：1/13(木) 22:16～、ほか<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>時は2014年、東京の片隅に位置する「東京はじっこテレビジョン」に入社した新人ADまふねこ(23)。<br />
華やかでキラキラした世界に胸を躍らせて配属された制作局のお仕事は、「どんぐりを600個拾う」、「カメラに映り込んだ(秘)にモザイクをかけ続ける」など地味で謎すぎる業務の連続だった！<br />
毒舌敏腕ディレクターや世界各国どこにでも飛ぶ海外ロケディレクターなど、世の中のフツーとはちょっと違う、けどなんかすごいギョーカイ人に囲まれながら、一人前のテレビマンを目指し仕事にまい進していく。<br />
テレビの裏側で悪戦苦闘するポンコツADまふねこの奮闘記は始まったばかり…!!</blockquote><p class="preface">
<br />
原作はテレビ東京のAD経験がある作者・真船佳奈によるお仕事エッセイマンガ。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4022142375/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4022142375/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/48-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：真船佳奈<br />
監督・シリーズ構成・脚本・絵コンテ：青木純<br />
キャラクターデザイン：青木純、劉洋<br />
美術：方琦星<br />
色彩設計：吉兆祺<br />
音楽：山田淳平<br />
音楽制作：ワーロック<br />
音楽協力：テレビ東京ミュージック<br />
音響効果：稲田祐介<br />
アニメーション制作：神南スタジオ、スぺ―スネコカンパニー<br />
コピーライト表記：<span translate="no">&#x00A9;真船佳奈・テレビ東京／オンエアできない！製作委員会</span><br />
<br />
<b>・キャスト</b><br />
まふねこ：大地葉<br />
鬼河原ディレクター：間宮康弘<br />
よりちゃん：杉山里穂<br />
横山チーフAD：広瀬裕也<br />
ナレーション：川島明(麒麟)<br />
<br />
Twitter：<a href="https://twitter.com/oadekinai" target="_blank">@oadekinai</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/オンエアできない" target="_blank">#オンエアできない</a><br />
<br />
ED：水湊いづき「Work out！」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=Qew3hVNpf0g" target="_blank">【公式】「オンエアできない！」番宣</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/Qew3hVNpf0g" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="47">◆<b><a href="https://garugaku.com/" target="_blank">ガル学。Ⅱ～Lucky Stars～</a></b></h2><p class="preface">
</p><a href="https://garugaku.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/47.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ東京系6局ネット：1/10(月) 7:05～(月曜～金曜放送)<br />
「おはスタ」内<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>ガールズアリーナで開催されたGirls&#x00B2;の「ツナグツナグ」ライブが伝説となった数年後――。<br />
再びその頂を目指すガールズたちの熱い挑戦が始まる。<br />
聖ガールズスクエア学院に入るため、多くのガールズたちが通う養成スタジオ「ガルスタ」。<br />
ガルスタのひとつ、原宿竹下通りスタジオに集まる6人の女の子たち。<br />
その名は「Lucky&#x00B2;(ラッキーラッキー)」！<br />
互いにぶつかり合いながら磨かれ輝きを魅せる成長の物語――。<br />
目指すは聖ガールズスクエア学院への転入をかけたオーディション。<br />
6つの小さなキラキラが新たなキセキを巻き起こす！</blockquote><p class="preface">
<br />
「おはスタ」内で放送されたアニメ「ガル学。～聖ガールズスクエア学院～」、および続編のドラマ「ガル学。～ガールズガーデン～」に続く作品。おはガールとしても活躍中のガールズ・パフォーマンスグループLucky&#x00B2;のメンバーやおはスタMC・木村昴らが出演する。<br />
<br />
<b>・スタッフ</b><br />
アニメーション監督：京極尚彦<br />
監督：中田誠<br />
シリーズ構成：藤咲淳一<br />
キャラクターデザイン：広岡トシヒト<br />
アニメーション制作：OLM<br />
製作：2202「ガル学。Ⅱ」製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;2022「ガル学。Ⅱ」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
山口リナ：山口莉愛<br />
杉浦ユウラ：杉浦優來<br />
永山ツバキ：永山椿<br />
深澤ヒイロ：深澤日彩<br />
比嘉ユウワ：比嘉優和<br />
佐藤カンナ：佐藤栞奈<br />
森村タケオ：木村昴<br />
羽村チノ：杉山里穂<br />
<br />
Twitter：<a href="https://twitter.com/garugaku" target="_blank">@garugaku</a><br />
<br />
OP：Lucky&#x00B2;「ichigo～一期一会～」<br />
<br />
</p><hr></p><p class="preface"><h2 id="28">◆<b><a href="https://tribenine.tokyo/anime/" target="_blank">トライブナイン</a></b></h2><p class="preface">
</p><a href="https://tribenine.tokyo/anime/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/28.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/10(月) 22:30～<br />
BS11：1/10(月) 23:00～<br />
テレビ愛知：1/11(火) 25:30～<br />
サンテレビ：1/12(水) 24:30～<br />
<br />
dアニメストア：1/10(月) 22:30～【地上波同時】<br />
他配信：1/15(土) 22:30～<br />
(アニメカ、アニメ放題、ABEMA、Amazon Prime Video、GYAO！、J:COMオンデマンド、スマートパスプレミアム、dTV、TVer、TELASA、niconico、Paravi、バンダイチャンネル、ひかりTV、Hulu、FOD、みるプラス、U-NEXT、GYAO！ストア、クランクイン！ビデオ、DMM動画、HAPPY!動画、ビデオマーケット、music.jp、MOVIEFULL、Rakuten TV)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>気が弱く、いつもいじめられてばかりの"白金ハル"<br />
最強の男を目指し、海の向こうからやってきた"タイガ"<br />
ふたりの少年はとある夕暮れ、最強のXBプレーヤーであり、ミナトトライブのリーダーである"神谷瞬"と出会う。<br />
それと時を同じくして、ネオトーキョーに散らばる各トライブは、大きな脅威にさらされようとしていた。<br />
ネオトーキョー国王"鳳 天心"の命令によって、謎の男"鳳 王次郎"率いるチヨダトライブが、国中のトライブの制圧をはじめたのだ。<br />
その魔の手は、ミナトトライブにも及ぼうとしており…。<br />
<br />
「自分を変えたい」<br />
「とにかく最強の男を目指す」<br />
「ただ、XBを楽しみたい」<br />
<br />
己の譲れぬ信念を貫くため、ハル、タイガ、神谷はXBの打席に立つ――！</blockquote><p class="preface">
<br />
「ダンガンロンパ」シリーズを手がけた小高和剛らが2017年に設立した「トゥーキョーゲームス」が、「ドラゴンボールZ ドッカンバトル」「八月のシンデレラナイン」などのスマホゲームを手がける「アカツキ」と組んで送り出すコンテンツ。スマートフォン向け3DアクションRPGとアニメ、WEBTOON(Web漫画)などが展開される。<br />
<br />
<b>・スタッフ</b><br />
原作：アカツキ×トゥーキョーゲームス<br />
原案：小高和剛<br />
キャラクター原案：小松崎類、しまどりる<br />
音楽：高田雅史<br />
総合プロデューサー：山口修平<br />
監督：青木悠<br />
シリーズ構成：横手美智子<br />
アニメーションキャラクターデザイン：薮本陽輔<br />
音楽制作：ランティス<br />
アニメーション制作：ライデンフィルム<br />
製作：トライブナイン製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9; Akatsuki Inc./トライブナイン製作委員会</span><br />
<br />
<b>・キャスト</b><br />
神谷瞬：石田彰<br />
白金ハル：堀江瞬<br />
タイガ：沢城千春<br />
有栖川さおり：渕上舞<br />
三田三太郎：田村睦心<br />
大門愛海：落合福嗣<br />
青山カズキ：千葉翔也<br />
鳳天心：中博史<br />
鳳王次郎：諏訪部順一<br />
神木結衣：小松未可子<br />
<br />
Twitter：<a href="https://twitter.com/tribenine_tokyo" target="_blank">@tribenine_tokyo</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/トライブナイン" target="_blank">#トライブナイン</a><br />
<br />
OP：MIYAVI「Strike It Out」<br />
ED：Void_Chords feat. LIO「Infocus」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=O6v0Y4_UNOs" target="_blank">TVアニメ『トライブナイン』アニメーションPV第1弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/O6v0Y4_UNOs" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="29">◆<b><a href="https://kendeshi-anime.com/" target="_blank">賢者の弟子を名乗る賢者</a></b></h2><p class="preface">
</p><a href="https://kendeshi-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/29.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/11(火) 24:30～<br />
BS日テレ：1/11(火) 24:30～<br />
MBS：1/11(火) 27:00～<br />
AT-X：1/15(土) 23:30～、ほか<br />
<br />
ABEMA：1/11(火) 24:00～【最速・地上波先行】<br />
dアニメストア：1/11(火) 24:00～【最速・地上波先行】<br />
他配信：1/18(火) 24:00～<br />
(niconico、GYAO!、ひかりTV、FOD、バンダイチャンネル、Hulu、TELASA、J:COMオンデマンド、milplus見放題プレミアム、アニメ放題、Amazon Prime Video、アニメカ、U-NEXT、Netflix)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>無限の可能性が広がるVRMMO-RPG『アーク・アースオンライン』。<br />
プレイヤーによって建国されたアルカイト王国の九賢者が一人、威厳あふれる老齢の召喚術士ダンブルフもまたプレイヤーの一人だった。<br />
ある日、彼は世界の異変に気づく。<br />
ゲームでは無かった味覚や臭覚が生まれ、ログアウトもできない。さらに、 NPCが実に人間くさい反応を見せる。<br />
――それはゲームが紛れもない現実となった証であった。<br />
しかもこの世界では、30年もの月日が経っているというのだ。<br />
そして何ということか、ダンブルフは諸事情により幼くも美しい少女の姿になっていた！<br />
急変した世界の謎を解き明かすため、ダンブルフは賢者の弟子ミラを名乗り旅立つのであった。<br />
冒険の果てに待ち受けているものとは――。</blockquote><p class="preface">
<br />
原作は「<b><a href="https://ncode.syosetu.com/n6829bd/" target="_blank">小説家になろう</a></b>」連載作品で、書籍版が2014年からGCノベルズで刊行中。既刊15巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4896374657/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4896374657/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/29-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：りゅうせんひろつぐ<br />
キャラクター原案：藤ちょこ<br />
監督：元永慶太郎<br />
シリーズ構成：鴻野貴光<br />
キャラクターデザイン：堀井久美<br />
美術設定：大平司<br />
美術監督：Scott MacDonald<br />
色彩設計：山本未有<br />
撮影監督：横山翼<br />
3D監督：後藤優一<br />
編集：木村祥明<br />
音響監督：えびなやすのり<br />
音響効果：川田清貴<br />
音響制作：ダックスプロダクション<br />
音楽：坂部剛<br />
音楽制作：MAGES.<br />
アニメーション制作：studio A-CAT<br />
製作：わしかわいい製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;2021 りゅうせんひろつぐ・藤ちょこ／ マイクロマガジン社／わしかわいい製作委員会</span><br />
<br />
<b>・キャスト</b><br />
ミラ：大森日雅<br />
ソロモン：村瀬歩<br />
ルミナリア日向未南<br />
ゼフ：斉藤隼一<br />
アスバル：安元洋貴<br />
フリッカ：櫻弥恵<br />
エメラ：夏吉ゆうこ<br />
タクト：井澤佳の実<br />
<br />
Twitter：<a href="https://twitter.com/kendeshi_anime" target="_blank">@kendeshi_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/賢でし" target="_blank">#賢でし</a><br />
<br />
OP：亜咲花「Ready Set Go!!」<br />
ED：エラバレシ「Ambitious」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=ZO5tYbvqGrc" target="_blank">TVアニメ『賢者の弟子を名乗る賢者』PV第3弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/ZO5tYbvqGrc" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="30">◆<b><a href="https://anime.priconne-redive.jp/" target="_blank">プリンセスコネクト！Re:Dive Season 2</a></b></h2><p class="preface">
</p><a href="https://anime.priconne-redive.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/30.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/10(月) 24:00～<br />
BS11：1/10(月) 24:00～<br />
KBS京都：1/10(月) 24:00～<br />
サンテレビ：1/10(月) 24:30～<br />
テレビ北海道：1/10(月) 25:30～<br />
TVQ九州放送：1/10(月) 26:00～、ほか<br />
テレビ愛知：1/10(月) 26:05～<br />
AT-X：1/11(火) 21:30～、ほか<br />
WOWOW：1/12(水) 24:00～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>紡いだ絆に思いを乗せて、ユウキたちの冒険が再び始まる。<br />
その出会いは突然だった。<br />
意気投合した彼らはあるギルドを結成する――その名は【美食殿】。<br />
美食の探求を目的とした彼らは愉快な仲間たちと友情を深め、<br />
美味しいごはんを食べ、そしてときにはちょっぴり危険な冒険に身を投じ、せわしなくも穏やかな日々を送っていた。<br />
ところが、胸に秘めた思いはやがて交錯し、彼らはかつてない困難に巻き込まれていく。</blockquote><p class="preface">
<br />
原作は<b><a href="https://priconne-redive.jp/" target="_blank">同名ソーシャルゲーム</a></b>。アニメ第1期が2020年4月～6月に放送・配信されている。<br />
<br />
<b>・スタッフ</b><br />
原作：Cygames<br />
総監督・シリーズ構成：金崎貴臣<br />
監督：いわもとやすお<br />
キャラクターデザイン：渡辺舞、野田康行、栗田聡美、楊烈駿<br />
色彩設計：合田沙織<br />
3DCGディレクター：中野祥典<br />
美術監督：小木曽宣久<br />
撮影監督：米澤寿<br />
編集：木村佳史子<br />
音楽：イマジン<br />
音響監督：金崎貴臣<br />
録音：山口貴之<br />
音響効果：小山恭正<br />
音響制作：東北新社<br />
アニメーション制作：CygamesPictures<br />
コピーライト表記：<span translate="no">&#x00A9; アニメ「プリンセスコネクト！Re:Dive」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
ペコリーヌ：M・A・O<br />
コッコロ：伊藤美来<br />
キャル：立花理香<br />
ユウキ：阿部敦<br />
<br />
Twitter：<a href="https://twitter.com/priconne_anime" target="_blank">@priconne_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/アニメプリコネR" target="_blank">#アニメプリコネR</a> <a href="https://twitter.com/hashtag/アニメプリコネ" target="_blank">#アニメプリコネ</a> <a href="https://twitter.com/hashtag/プリコネR" target="_blank">#プリコネR</a><br />
<br />
<b><a href="https://www.youtube.com/watch?v=va2xUHy_xeQ" target="_blank">アニメ「プリンセスコネクト！Re:Dive Season 2」第2弾PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/va2xUHy_xeQ" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="31">◆<b><a href="https://sabikuibisco.jp/" target="_blank">錆喰いビスコ</a></b></h2><p class="preface">
</p><a href="https://sabikuibisco.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/31.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/10(月) 24:30～<br />
BS11：1/10(月) 24:30～<br />
読売テレビ：1/10(月) 25:59～<br />
AT-X：1/12(水) 21:00～、ほか<br />
<br />
ABEMA：1/10(月) 24:30～【地上波同時】<br />
dアニメストア：1/13(木) 24:30～<br />
他配信：1/17(月) 24:30～<br />
(niconico、GYAO!、ひかりTV、Netflix、FOD、バンダイチャンネル、Hulu、TELASA、J:COMオンデマンド、みるプラス見放題プレミアム、U-NEXT、アニメ放題、Amazon Prime Video、アニメカ)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>全てを錆びつかせる≪錆び風≫が吹き荒れる日本。<br />
人々は街や生命を蝕む錆に怯えながら暮らしていた。<br />
忌み嫌われる≪キノコ守り≫の一族の少年・赤星ビスコは、瀕死の師匠を救うため、全ての錆を浄化する霊薬キノコ≪錆喰い≫を求めて旅をしていた。<br />
旅の途中、忌浜（いみはま）で出会った美貌の少年医師・猫柳ミロもまた、大切な姉を蝕む錆の対処法を探していた。<br />
愛する者を救うべく、ふたりの少年が手を取る時、新たな冒険が始まる。<br />
人の心までは錆びつかない。</blockquote><p class="preface">
<br />
原作は電撃文庫から刊行されている“世紀末キノコアクション”小説で、既刊8巻。宝島社「<b><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4800290449/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4800290449/gigazine-22');">このライトノベルがすごい！2019</a></b>」で、初めて総合・新作でダブル1位を獲得した作品。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4048936166/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4048936166/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/31-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：瘤久保慎司<br />
原作イラスト：赤岸K<br />
原作世界観イラスト：mocha<br />
監督：碇谷敦<br />
副監督：又賀大介<br />
シリーズ構成・脚本：村井さだゆき<br />
キャラクターデザイン：浅利歩惟、碇谷敦<br />
総作画監督：浅利歩惟、井川典恵<br />
メインアニメーター：松原豊、河合桃子<br />
動画監督：張逸暉<br />
美術監督：三宅昌和<br />
美術・クリーチャー設定：曽野由大<br />
色彩設計：千葉絵美<br />
撮影監督：高木翼<br />
編集：木村祥明<br />
音楽：上田剛士(AA=)、椿山日南子<br />
音楽制作：フライングドッグ<br />
音響監督：小泉紀介<br />
アニメーション制作：OZ<br />
コピーライト表記：<span translate="no">&#x00A9;2021 瘤久保慎司/KADOKAWA/錆喰いビスコ製作委員会</span><br />
<br />
<b>・キャスト</b><br />
赤星ビスコ：鈴木崚汰<br />
猫柳ミロ：花江夏樹<br />
猫柳パウー：近藤玲奈<br />
大茶釜チロル：富田美憂<br />
ジャビ：斎藤志郎<br />
黒革：津田健次郎<br />
<br />
Twitter：<a href="https://twitter.com/SABIKUI_BISCO" target="_blank">@SABIKUI_BISCO</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/錆喰いビスコ" target="_blank">#錆喰いビスコ</a><br />
<br />
OP：JUNNA「風の音さえ聞こえない」<br />
ED：ビスコ＆ミロ(CV：鈴木崚汰＆花江夏樹)「咆哮」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=ZAj5_2L03PE" target="_blank">【2022年1月10日より放送開始!!】TVアニメ『錆喰いビスコ』本PV第2弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/ZAj5_2L03PE" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="32">◆<b><a href="https://gensou-sangokushi.com/" target="_blank">幻想三國誌 -天元霊心記-</a></b></h2><p class="preface">
</p><a href="https://gensou-sangokushi.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/32.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
BS12：1/10(月) 26:00～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>王朝末期。神州中原は群雄割拠の時代を迎えていた。<br />
そんな中、さらなる混乱を引き起こしているのが〝魍魎〟たちであった。<br />
死んだ人々の無念や恐怖=瘴気から生まれ、さらに生きている人間の欲望や怒りといった負の感情に憑依して、暴れ狂う怪物に変えてしまう魍魎。その魍魎によって戦乱は加速し、多くの死がさらなる魍魎を生み出していた。<br />
この世の秩序を守る神や仙人の組織〝天元〟はこうした事態を憂慮し、特別な力を持つ人間を選び出し、魍魎たちを狩る〝魎狩隊〟を結成することにした。</blockquote><p class="preface">
<br />
原作はファンタジースペクタクルRPG「<b><a href="https://www.falcom.co.jp/fs/" target="_blank">幻想三国誌</a></b>」。当初は2020年10月放送開始予定だったが、「<b><a href="https://gensou-sangokushi.com/512" target="_blank">諸般の事情</a></b>」のため、放送延期になっていた。<br />
<br />
<b>・スタッフ</b><br />
原作：「幻想三國誌」<br />
監督：町谷俊輔<br />
副監督：永居慎平<br />
シリーズ構成：鈴木雅詞<br />
キャラクターデザイン：CSPG、油井徹太郎<br />
色彩設計：金丸祐子<br />
美術設定：大原盛仁<br />
美術監督：片野坂恵美(インスパイアード)<br />
美術アドバイザー：増山修<br />
CG監督：宮下るりか(しいたけデジタル)<br />
撮影監督：浅川茂輝(レアトリック)<br />
編集：宮崎直樹(森田編集室)<br />
音楽：田頭勉<br />
音響監督：丹下雄二<br />
音響効果：宅間麻姫<br />
録音：佐々木彰<br />
音響制作：東北新社<br />
音楽制作：GEEKTOYS<br />
アニメーション制作：GEEKTOYS<br />
製作：「幻想三國誌」製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;2021 USERJOY Technology CO.,LTD./「幻想三國誌」製作委員会</span><br />
<br />
<b>・キャスト</b><br />
應幾：小野賢章<br />
小霊：金元寿子<br />
丁研：古川慎<br />
洵喬：中原麻衣<br />
愛心茶桶：福島潤<br />
羅織女：大原さやか<br />
<br />
Twitter：<a href="https://twitter.com/gensou_3gokushi" target="_blank">@gensou_3gokushi</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/幻想三國誌" target="_blank">#幻想三國誌</a> <a href="https://twitter.com/hashtag/天元霊心記" target="_blank">#天元霊心記</a><br />
<br />
OP：Machico「ENISHI」<br />
ED：LACCO TOWER「嘘」<br />
<br />
</p><hr></p><p class="preface"><h2 id="50">◆<b><a href="https://twitter.com/Shocker_SAN1111" target="_blank">お昼のショッカーさん</a></b></h2><p class="preface">
</p><a href="https://twitter.com/Shocker_SAN1111" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/50.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
ABC：1/20(木) 26:34～<br />
TOKYO MX：1/21(金) 19:28～<br />
テレビ愛知：1/22(土) 26:25～(「<b><a href="https://tv-aichi.co.jp/koebu/" target="_blank">キャラ＠声部</a></b>」内)<br />
<br />
<b><a href="https://tokusatsu-fc.jp/" target="_blank">東映特撮ファンクラブ</a></b>：1/11(火) 11:00～(第2話は1/27(水)12:00～)<br />
<br />
<b>・作品情報</b><br />
原作は<b><a href="https://manga.line.me/product/periodic?id=Z0001739" target="_blank">LINEマンガ</a></b>。<br />
<br />
<b>・スタッフ</b><br />
マンガクリエイター：STUDY優作<br />
アニメーション制作：勝鬨スタジオ<br />
監督：ナミキヒロシ<br />
音響制作：チャンスイン<br />
音響監督：サイトウユウ<br />
楽曲制作：佐々木宏人<br />
コピーライト表記：<span translate="no">&#x00A9;STUDY YU-SAKU/MMDGP　&#x00A9;石森プロ・東映</span><br />
<span translate="no">&#x00A9;アニメ「お昼のショッカーさん」製作委員会　&#x00A9;石森プロ・東映</span><br />
<span translate="no">&#x00A9;石森プロ・東映　Illustrated by STUDY YU-SAKU</span><br />
<br />
<b>・キャスト</b><br />
蒼井翔太<br />
榎木淳弥<br />
八代拓<br />
効果音：関智一<br />
<br />
Twitter：<a href="https://twitter.com/Shocker_SAN1111" target="_blank">@Shocker_SAN1111</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/お昼のショッカーさん" target="_blank">#お昼のショッカーさん</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="33">◆<b><a href="https://tensaiouji-anime.com/" target="_blank">天才王子の赤字国家再生術</a></b></h2><p class="preface">
</p><a href="https://tensaiouji-anime.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/33.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
AT-X：1/11(火) 22:30～、ほか<br />
TOKYO MX：1/11(火) 23:00～<br />
BS日テレ：1/11(火) 24:00～<br />
<br />
ABEMA：1/11(火) 23:00～【地上波同時】<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>覇権国家の脅威に晒される弱小国家・ナトラ王国。<br />
若くして国を背負うことになった王子・ウェインは、補佐官のニニムに支えられながら、才能を活かした見事な手腕を発揮し始める。<br />
でも、この国……めちゃくちゃ詰んでる！<br />
内政に手を入れようにも金がない。<br />
よそから奪おうにも軍事力がない。<br />
まともで優秀な人材は他国に流出してしまう。<br />
「早く国売ってトンズラしてえ」<br />
ウェインの願いは、とっとと隠居して悠々自適の生活を送ること。<br />
大国に媚びを売り、国を売れば夢が叶うはず。<br />
しかし、外交も軍事も予想外の方向へ転がってしまい……!?<br />
知恵と機転で世界を揺るがす天才王子の弱小国家マネジメント、ここに開幕！</blockquote><p class="preface">
<br />
原作はGA文庫から刊行されているライトノベル「そうだ、売国しよう～天才王子の赤字国家再生術～」。既刊9巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4797397039/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4797397039/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/33-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：鳥羽徹<br />
キャラクター原案：ファルまろ<br />
監督：玉川真人<br />
副監督：蔡欣亞<br />
シリーズ構成：赤尾でこ<br />
キャラクターデザイン：應地隆之介<br />
3Dディレクター：墳下芳弘<br />
美術監督：栫ヒロツグ(鹿児島ラメカヒリム)<br />
色彩設計：中尾総子(Wish)<br />
編集：長谷川舞(editz)<br />
撮影監督：内田奈津美(アニモキャラメル)<br />
音響監督：納谷僚介<br />
音響制作：スタジオマウス<br />
音楽：佐橋俊彦<br />
アニメーションプロデューサー：大上裕真<br />
アニメーション制作：横浜アニメーションラボ<br />
コピーライト表記：<span translate="no">&#x00A9;鳥羽徹・SBクリエイティブ／天才王子製作委員会</span><br />
<br />
<b>・キャスト</b><br />
ウェイン：斉藤壮馬<br />
ニニム：高橋李依<br />
フラーニャ：千本木彩花<br />
ロウェルミナ：東山奈央<br />
フィシュ・ブランデル：日笠陽子<br />
ナナキ・ラーレイ：榊原優希<br />
ゼノ：中島由貴<br />
<br />
Twitter：<a href="https://twitter.com/tensaiouji_PR" target="_blank">@tensaiouji_PR</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/天才王子" target="_blank">#天才王子</a><br />
<br />
OP：やなぎなぎ×THE SIXTH LIE「LEVEL」<br />
ED：南條愛乃「ヒトリとキミと」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=unmviapbywM" target="_blank">TVアニメ「天才王子の赤字国家再生術」PV</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/unmviapbywM" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="39">◆<b><a href="https://fabiniku.com/" target="_blank">異世界美少女受肉おじさんと</a></b></h2><p class="preface">
</p><a href="https://fabiniku.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/39.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ東京：1/11(火) 24:00～<br />
北海道テレビ：1/11(火) 25:55～<br />
AT-X：1/12(水) 21:30～、ほか<br />
BSテレ東：1/14(金) 24:30～<br />
<br />
配信：1月～<br />
(ABEMA、Google Play、ひかりTV、ツイキャス、バンダイチャンネル、Hulu、アニメ放題、U-NEXT、Amazon Prime Video、スマートパスプレミアム、J:COMオンデマンド、TELASA、dアニメストア、みるプラス、アニメフェスタ、FOD、dTV、ニコニコ動画、GYAO！ストア、クランクイン！ビデオ、DMM.com、ビデオマーケット、music.jp、Rakuten TV、Happy!動画)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>「行くぞ神宮寺」<br />
「あぁ、俺たちで魔王を倒す」<br />
「「俺たちが、互いに惚れる前に！！」」<br />
平凡なサラリーマン生活を送っていた幼馴染の橘日向(32)と神宮寺司(32)。二人はとある合コンの帰り道に、女神を名乗る謎の存在によって異世界に飛ばされてしまう。そこで神宮寺が目にしたのは、金髪碧眼美少女の姿に変わり果てた親友の姿だった…!?<br />
美少女の姿になった橘を見て、あまりの可愛さに戸惑う神宮寺。女性の身体になったことで、神宮寺のかっこよさにグッと来てしまう橘。しかし元は親友同士。この関係を壊さないためにも、一刻も早く魔王を打倒して元の姿に戻らなければならない。<br />
そう、お互いがお互いを好きになる前に――。<br />
これは、おっさんと元おっさん美少女の、絶対に惚れてはいけない異世界ラブコメディー</blockquote><p class="preface">
<br />
原作は「<b><a href="https://cycomi.com/fw/cycomibrowser/chapter/title/138" target="_blank">サイコミ</a></b>」連載中作品で、最新6巻が2022年1月19日発売予定。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4098501023/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4098501023/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/39-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：池澤真・津留崎優／Cygames<br />
監督：山井紗也香<br />
シリーズ構成：竹内利光<br />
キャラクターデザイン：大和葵<br />
音響監督：亀山俊樹<br />
音楽：渡辺剛<br />
音響制作：ビットグルーヴプロモーション<br />
アニメーション制作：OLM Team Yoshioka<br />
製作：ファ美肉製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;池澤真・津留崎優・Cygames / ファ美肉製作委員会</span><br />
<br />
<b>・キャスト</b><br />
橘日向(女)：M・A・O<br />
神宮寺司：日野聡<br />
橘日向(男)：伊東健人<br />
愛と美の女神：釘宮理恵<br />
ティロリロ・リリリ・ルー：藤井ゆきよ<br />
シュバルツ・フォン・リヒテンシュタイン・ローエングラム：石川界人<br />
ルシウス：喜多村英梨<br />
シェン：諏訪部順一<br />
ムリア：芹澤優<br />
ユグレイン：牧野由依<br />
ナレーション：江原正士<br />
<br />
<b><a href="https://www.youtube.com/watch?v=Awp_B8E2Ktg" target="_blank">TVアニメ『異世界美少女受肉おじさんと』PV第1弾</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/Awp_B8E2Ktg" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="46">◆<b><a href="https://love-of-kill.com/" target="_blank">殺し愛</a></b></h2><p class="preface">
</p><a href="https://love-of-kill.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/46.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/12(水) 24:00～<br />
BS日テレ：1/12(水) 24:00～<br />
AT-X：1/13(木) 22:30～、ほか<br />
サンテレビ：1/13(木) 24:00～<br />
KBS京都：1/13(水) 25:00～<br />
<br />
dアニメストア：1/12(水) 23:30～【地上波先行】<br />
他配信：1/19(水) 23:30～<br />
(ABEMA、ニコニコ生放送、ニコニコチャンネル、GYAO!、ひかりTV、Netflix、FOD、バンダイチャンネル、Hulu、TELASA、J:COMオンデマンド、みるプラス見放題プレミアム、アニメ放題、Amazon Prime Video、アニメカ、U-NEXT)<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>とある「仕事場」で対峙する2人の殺し屋。<br />
クールな賞金稼ぎの女・シャトーと謎多き最強の男・リャンハ。<br />
シャトーはこの交戦をきっかけにリャンハと敵対――するはずが、なぜか彼に気に入られ、つきまとわれることに。<br />
彼女はなし崩し的にリャンハと協力関係を結んでしまうが、彼を狙う組織との抗争に巻き込まれていく。<br />
さらにその戦いは、彼女の過去とも関係しているのだった。<br />
リャンハはなぜシャトーに接近するのか。シャトーに秘められた過去とは。<br />
相性最悪の2人が織りなす、「殺し屋×殺し屋」の歪なサスペンス。<br />
奇妙な運命の歯車がいま動き出す。</blockquote><p class="preface">
<br />
原作は「コミックジーン」連載中で、既刊11巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4040682386/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4040682386/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/46-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：Fe<br />
監督：大庭秀昭<br />
シリーズ構成・脚本：久尾歩<br />
キャラクターデザイン：佐藤陽子<br />
サブキャラクターデザイン：小林利充<br />
総作画監督：佐藤陽子、小林利充<br />
アクション作画監督：才木康寛<br />
3D・プロップデザイン：杉村友和<br />
美術監督：黛昌樹<br />
色彩設計：山上愛子<br />
撮影監督：伏原あかね<br />
編集：木村佳史子<br />
音響監督：髙桑一<br />
音響効果：和田俊也<br />
音響制作：ビットグルーヴプロモーション<br />
音楽：吉川慶<br />
音楽制作：TOY'S FACTORY<br />
音楽制作協力：ミラクル・バス<br />
アニメーション制作：プラチナビジョン<br />
製作：殺し愛製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;2022 Fe/KADOKAWA/殺し愛製作委員会</span><br />
<br />
<b>・キャスト</b><br />
シャトー・ダンクワース：大西沙織<br />
ソン・リャンハ：下野紘<br />
エウリペデス・リッツラン：堀内賢雄<br />
ジム：天﨑滉平<br />
ホー：前野智昭<br />
ジノン：村瀬歩<br />
ニッカ：森田成一<br />
ミファ：日笠陽子<br />
ドニー：大塚芳忠<br />
<br />
Twitter：<a href="https://twitter.com/LoveofKill_info" target="_blank">@LoveofKill_info</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/殺し愛" target="_blank">#殺し愛</a><br />
<br />
OP：増田俊樹「Midnight Dancer」<br />
ED：小林愛香「マコトピリオド」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=QlMuqsQacD8" target="_blank">TVアニメ『殺し愛』PV第2弾｜2022年1月12日より放送開始‼</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/QlMuqsQacD8" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="34">◆<b><a href="https://heike-anime.asmik-ace.co.jp/" target="_blank">平家物語</a></b></h2><p class="preface">
</p><a href="https://heike-anime.asmik-ace.co.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/34.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
フジテレビ：1/12(水) 24:55～<br />
「<b><a href="https://plus-ultra.tv/" target="_blank">＋Ultra</a></b>」枠<br />
<br />
FOD：先行配信済み<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>《祇園精舎の鐘の声、諸行無常の響きあり 沙羅双樹の花の色、盛者必衰の理をあらはす》<br />
平安末期。平家一門は、権力・武力・財力あらゆる面で栄華を極めようとしていた。亡者が見える目を持つ男・平重盛は、未来が見える目を持つ琵琶法師の少女・びわに出会い、「お前たちはじき滅びる」と予言される。<br />
貴族社会から武家社会へ――日本が歴史的転換を果たす、激動の15年が幕を開ける。</blockquote><p class="preface">
<br />
原作は古川日出男による現代語訳「平家物語」。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4309728790/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4309728790/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/34-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：古川日出男訳 『池澤夏樹=個人編集 日本文学全集09 平家物語』　河出書房新社刊<br />
監督：山田尚子<br />
脚本：吉田玲子<br />
キャラクター原案：高野文子<br />
音楽：牛尾憲輔<br />
アニメーション制作：サイエンスSARU<br />
キャラクターデザイン：小島崇史<br />
美術監督：久保友孝(でほぎゃらりー)<br />
動画監督：今井翔太郎<br />
色彩設計：橋本賢<br />
撮影監督：出水田和人<br />
編集：廣瀬清志<br />
音響監督：木村絵理子<br />
音響効果：倉橋裕宗（Otonarium）<br />
歴史監修：佐多芳彦<br />
琵琶監修：後藤幸浩<br />
<br />
<b>・キャスト</b><br />
びわ：悠木碧<br />
平重盛：櫻井孝宏<br />
平徳子：早見沙織<br />
平清盛：玄田哲章<br />
千葉繁<br />
井上喜久子<br />
入野自由<br />
小林由美子<br />
岡本信彦<br />
花江夏樹<br />
村瀬歩<br />
西山宏太朗<br />
檜山修之<br />
木村昴<br />
宮崎遊<br />
水瀬いのり<br />
杉田智和<br />
梶裕貴<br />
<br />
Twitter：<a href="https://twitter.com/heike_anime" target="_blank">@heike_anime</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/平家物語" target="_blank">#平家物語</a><br />
<br />
OP：羊文学「光るとき」<br />
ED：agraph feat. ANI（スチャダラパー）「unified perspective」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=n27irsU7x6c" target="_blank">TVアニメ「平家物語」第二弾PV　2022年1月12日(水)よりフジテレビ「+Ultra」にて毎週水曜24:55〜放送／FODにて全話先行配信中</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/n27irsU7x6c" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="54">◆<b><a href="https://pjsekai.sega.jp/" target="_blank">ぷちセカ</a></b></h2><p class="preface">
</p><a href="https://pjsekai.sega.jp/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/54.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
TOKYO MX：1/13(木) 25:00～<br />
<br />
<b><a href="https://www.youtube.com/channel/UCdMGYXL38w6htx6Yf9YJa-w" target="_blank">YouTubeプロセカ公式チャンネル</a></b>：19:00～<br />
<br />
<b>・スタッフ</b><br />
コピーライト表記：<span translate="no">&#x00A9; SEGA／&#x00A9; Colorful Palette Inc.／&#x00A9; Crypton Future Media, INC. www.piapro.net piapro All rights reserved.</span><br />
<br />
Twitter：<a href="https://twitter.com/pj_sekai" target="_blank">@pj_sekai</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/プロセカ" target="_blank">#プロセカ</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="38">◆<b><a href="https://twitter.com/anime_kakeakami/" target="_blank">あたしゃ川尻こだまだよ～デンジャラスライフハッカーのただれた生活～</a></b></h2><p class="preface">
</p><a href="https://twitter.com/anime_kakeakami/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/38.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
フジテレビ：1/13(木) 25:25～(「<b><a href="https://www.fujitv.co.jp/exitv/" target="_blank">EXITV</a></b>」内)<br />
<br />
FOD：1/13(木) 26:35～(第1話～第12話先行配信)<br />
全24話<br />
<br />
<b>・作品情報</b><br />
漫画家・川尻こだまによる不摂生生活のエッセイ漫画が原作。なお、単行本にするにあたって全編フルカラー化＆描き下ろしの追加が行われている。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4046808284/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4046808284/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/38-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
なお、Twitterでの連載をざっくりとまとめてオマケを追加した上で、作者自らが「川尻こだまのただれた生活」として電子書籍を刊行している。既刊5巻。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B09MVW2NVY/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/B09MVW2NVY/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/38-03.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・スタッフ</b><br />
原作：川尻こだま<br />
アニメーション制作：ラパントラック<br />
製作：ただれた生活委員会<br />
コピーライト表記：<span translate="no">&#x00A9;川尻こだま／KADOKAWA・ただれた生活委員会</span><br />
<br />
<b>・キャスト</b><br />
川尻こだまとだいたい全部：悠木碧<br />
<br />
Twitter：<a href="https://twitter.com/anime_kakeakami" target="_blank">@anime_kakeakami</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/川尻こだま" target="_blank">#川尻こだま</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="37">◆<b><a href="https://rymansclub.com/" target="_blank">リーマンズクラブ</a></b></h2><p class="preface">
</p><a href="https://rymansclub.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/37.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ朝日系全国24局ネット：1/22(土) 25:30～(「<b><a href="https://www.tv-asahi.co.jp/numanimation/" target="_blank">NUMAnimation</a></b>」枠)<br />
CSテレ朝チャンネル1：1/22(土) 26:30～<br />
BS朝日：1/28(金) 23:00～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>天才的な観察眼で、バドミントン選手として活躍していた白鳥尊。<br />
しかし、インターハイでのトラウマが原因で、思うようなプレーができずにいた。<br />
社会人選手として所属していた強豪チーム・ミツホシ銀行をクビになった尊。<br />
選手としての再起をかけて、サンライトビバレッジに入社した彼を待ち受けていたのは、慣れない会社員としての仕事に、結果の出せていない弱小バドミントン部…。<br />
おまけに、ガサツで、声がでかくて、やたらと距離の近いおっさん…宮澄建。<br />
元・天才の新人社会人と、豪快おっさん“バドリーマン”。<br />
何もかもが正反対のコンビが生まれたとき、諦めかけていた夢が、ふたたび幕を開ける。</blockquote><p class="preface">
<br />
サラリーマンとバドミントン選手の2面を描くオリジナルの実業団バドミントンアニメ。<br />
<br />
<b>・スタッフ</b><br />
原作：Team RMC<br />
監督：山内愛弥<br />
シリーズ構成：内海照子、山内愛弥<br />
キャラクター原案：ヤスダスズヒト<br />
キャラクターデザイン：まじろ<br />
美術設定：佐藤正浩(ヘッドワークス)<br />
美術監督：甲斐政俊<br />
美術背景：スタジオKAIMU<br />
色彩設計：小野寺笑子<br />
2Dワークス：渡部岳、中島俊、旭プロダクション デザイン部<br />
3Dワークス：FelixFilm<br />
撮影監督：長谷川奈穂<br />
撮影：旭プロダクション<br />
編集：長谷川舞(editz)<br />
音楽：fox capture plan<br />
音響監督：本山哲<br />
音響制作：グロービジョン<br />
アニメーション制作：ライデンフィルム<br />
製作：サンライトビバレッジ広報部<br />
コピーライト表記：<span translate="no">&#x00A9;Team RMC／サンライトビバレッジ広報部</span><br />
<br />
<b>・キャスト</b><br />
白鳥尊：榎木淳弥<br />
宮澄建：三木眞一郎<br />
佐伯蒼汰：石川界人<br />
佐伯橙也：逢坂良太<br />
竹田浩輝：柿原徹也<br />
<br />
Twitter：<a href="https://twitter.com/rymansclub_PR" target="_blank">@rymansclub_PR</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/リーマンズクラブ" target="_blank">#リーマンズクラブ</a><br />
<br />
OP：Novelbright「The Warrior」<br />
<br />
<b><a href="https://www.youtube.com/watch?v=4oLOpLShm8Y" target="_blank">オリジナルTVアニメ「リーマンズクラブ」ティザーPV【2022年1月放送開始！】</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/4oLOpLShm8Y" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="53">◆<b><a href="https://luminaria.tales-ch.jp/anime/" target="_blank">テイルズ オブ ルミナリア The Fateful Crossroad</a></b></h2><p class="preface">
</p><a href="https://luminaria.tales-ch.jp/anime/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/53.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
配信：1/22(土)～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>かつて、この地には山のように巨大な獣たちが生きていた。<br />
その骸の周りは「マナ」に溢れ、そのマナを求めた人々が集まり、いくつかの国が生まれた。<br />
いつしか人々はマナの源たるその獣たちを「源獣」と崇め、共に暮らすようになった。<br />
これが源獣信仰の始まりである。<br />
そして、時は流れ――。<br />
源獣信仰を基盤とした国々によるユール連邦と、独自技術によって飛躍的に発展を遂げたジルドラ帝国の間で戦争が勃発。<br />
以後、戦争は激化を辿っていく。<br />
連邦の若き騎士候補生レオはある任務のため、幼馴染で同じく候補生のセリア、教官のリゼットと共に、かつて帝国領だった国境付近の街、リュンヌを訪れる。<br />
しかし、そこに現れたのは祖国を裏切り、帝国の兵士となった友、ユーゴだった――。</blockquote><p class="preface">
<br />
原作は「テイルズ オブ」シリーズのスマートフォン向け“21の生き様が交錯するRPG”。<br />
<br />
<b>・スタッフ</b><br />
キャラクターデザイン：佐伯俊<br />
コピーライト表記：<span translate="no">&#x00A9;BANDAI NAMCO Entertainment Inc.</span><br />
<br />
<b>・キャスト</b><br />
レオ・フルカード：新井良平<br />
セリア・アルヴィエ：岡咲美保<br />
リゼット・レニエ：嶋村侑<br />
リュシアン・デュフォール：上村祐翔<br />
ユーゴ・シモン：竹田海渡<br />
アレクサンドラ・フォン・ゾンネ：大西沙織<br />
アウグスト・ヴァレンシュタイン：梅原裕一郎<br />
バスチアン・フォルジュ：竹内良太<br />
<br />
Twitter：<a href="https://twitter.com/to_luminaria" target="_blank">@to_luminaria</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/ルミナリア" target="_blank">#ルミナリア</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="35">◆<b><a href="https://chikyugai.com/" target="_blank">地球外少年少女</a></b></h2><p class="preface">
</p><a href="https://chikyugai.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/35.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
Netflix：1/28(金)～<br />
<br />
<b>・作品情報</b><br />
NHKで放送されたオリジナルアニメ「電脳コイル」を手がけた磯光雄が、2045年の宇宙を舞台に「AIがある宇宙での暮らし」を描く、全6話構成のオリジナルアニメ。3話ずつの構成で劇場上映されるのと同時にBlu-ray＆DVDが発売され、Netflixで配信される。<br />
<br />
<b>・スタッフ</b><br />
原作・脚本・監督：磯光雄<br />
キャラクターデザイン：吉田健一<br />
メインアニメーター：井上俊之<br />
美術監督：池田裕輔<br />
色彩設計：田中美穂<br />
音楽：石塚玲依<br />
音響監督：清水洋史<br />
制作：Production +h.<br />
製作：地球外少年少女製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;MITSUO ISO／avex pictures・地球外少年少女製作委員会</span><br />
<br />
Twitter：<a href="https://twitter.com/Chikyugai_BG" target="_blank">@Chikyugai_BG</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/地球外少年少女" target="_blank">#地球外少年少女</a><br />
<br />
<b><a href="https://www.youtube.com/watch?v=BPXgdELm5g8" target="_blank">オリジナルアニメ『地球外少年少女』特報／1月28日前編、2月11日後編、各2週限定劇場上映</a></b><br />
<iframe class="lazyload yt_iframe" width="560" height="315" data-src="https://www.youtube.com/embed/BPXgdELm5g8" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe><br />
<br />
</p><hr></p><p class="preface"><h2 id="36">◆<b><a href="https://eienno831.com/" target="_blank">永遠の831</a></b></h2><p class="preface">
</p><a href="https://eienno831.com/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/36.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
WOWOW：1/30(日) 20:00～<br />
<br />
WOWOWオンデマンド：1/30(日) 20:00～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>“未曾有の大災厄”により世界中が混迷を極める現代。<br />
東京で新聞奨学生として暮らす青年・スズシロウは、誰にも言えない秘密を抱えていた。高校三年の夏に起こったある事件をきっかけとして、自分の意志とは無関係に周囲の時間を止められるようになってしまったのだ。<br />
そんなある日、スズシロウは初めてその秘密を分かち合える相手と出会う。自分と同じく、心に傷を負い、時間が止められるようになってしまった少女・なずな。<br />
彼女が兄の芹によって犯罪に利用されていることを知ったスズシロウは、衝動的に手を差し伸べる。<br />
「……君、ここから一緒に逃げよう!」<br />
止まってしまった時間を、二人は再び動かすことができるのか。</blockquote><p class="preface">
<br />
「攻殻機動隊 STAND ALONE COMPLEX」「東のエデン」などを手がけてきた神山健治による新たなオリジナルアニメ。<br />
<br />
<b>・スタッフ</b><br />
監督・脚本：神山健治<br />
キャラクターデザイン：河野紘一郎<br />
CGディレクター：松浦宏樹<br />
テクニカルディレクター：田尻真輝<br />
美術監督：滝口比呂志<br />
アニメーション制作：CRAFTAR<br />
コピーライト表記：<span translate="no">&#x00A9;神山健治・ＣＲＡＦＴＡＲ・ＷＯＷＯＷ／「永遠の８３１」ＷＯＷＯＷ</span><br />
<br />
<b>・キャスト</b><br />
浅野スズシロウ：斉藤壮馬<br />
橋本なずな：M・A・O<br />
亜川芹：興津和幸<br />
アキナ：日笠陽子<br />
室戸：大塚明夫<br />
双六：粟津貴嗣<br />
坂田兄弟(兄)：近藤孝行<br />
坂田兄弟(弟)：岩崎諒太<br />
ドンキ：山下誠一郎<br />
各務恭子：五十嵐麗<br />
<br />
Twitter：<a href="https://twitter.com/eienno831" target="_blank">@eienno831</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/永遠の８３１" target="_blank">#永遠の８３１</a><br />
<br />
OP：カノエラナ<br />
ED：angela<br />
<br />
</p><hr></p><p class="preface"><h2 id="49">◆<b><a href="https://www.tv-tokyo.co.jp/anime/kinder/intro/sakuhin30/" target="_blank">テイコウペンギン</a></b></h2><p class="preface">
</p><a href="https://www.tv-tokyo.co.jp/anime/kinder/intro/sakuhin30/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/49.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ東京系6局ネット：1月～<br />
「<b><a href="https://www.tv-tokyo.co.jp/anime/kinder/" target="_blank">きんだーてれび</a></b>」内<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>ブラック企業で働くペンギンが、社会の理不尽にテイコウするアニメ。<br />
上司からの無茶振りや、同僚からの押し付け。<br />
そして、な〜んか憎めない後輩。<br />
一言言いたいけど、今後の関係を考えるとそうもいかない。<br />
そんな理不尽に対して、言いたかった一言を言ってくれるペンギンたちのお話。</blockquote><p class="preface">
<br />
原作はイラストレーター・<b><a href="https://twitter.com/torinosashimi" target="_blank">とりのささみ。</a></b>がTwitterで発表した漫画。描き下ろしを追加した書籍版が2019年に発売されている。<br />
</p><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4065152119/gigazine-22" target="_blank" onclick="ga('send', 'event', 'amazonLink', '20211215-anime-2022winter', 'exec/obidos/ASIN/4065152119/gigazine-22');"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/49-02.png" border="0" class="lazyload"></a><p class="preface">
<br />
また、すでにYouTubeでアニメ配信が行われている。テレビアニメは配信版とは違うオリジナルストーリー。<br />
<br />
<b>・スタッフ</b><br />
原作：とりのささみ。<br />
制作：株式会社Plott<br />
<br />
Twitter：<a href="https://twitter.com/teikoupenguin" target="_blank">@teikoupenguin</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/テイペン" target="_blank">#テイペン</a><br />
<br />
</p><hr></p><p class="preface"><h2 id="55">◆<b><a href="https://www.tv-tokyo.co.jp/anime/yamishibai10/" target="_blank">闇芝居(第10期)</a></b></h2><p class="preface">
</p><a href="https://www.tv-tokyo.co.jp/anime/yamishibai10/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/55.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ東京：2022年1月～<br />
<br />
<b>・作品情報</b><br />
</p><blockquote>「都市伝説」を題材として、2013年から放送されているショートアニメシリーズの第10作目。</blockquote><p class="preface">
<br />
<b>・スタッフ</b><br />
全体演出：船田晃、杉本健一<br />
各話演出：平田貢一、宮䑓直也、井上洋輔、横山恵<br />
脚本：熊本浩武、佐々木充郭、石上加奈子<br />
作画：鹿目凛、jimmy、菅谷孝一、渡邉柊、日暮桃花、風月<br />
企画：山川典夫（テレビ東京）、岩﨑拓矢（ILCA）<br />
プロデューサー：梅津智史（テレビ東京）、船田晃<br />
制作協力：yell、DRAWIZ<br />
制作：ILCA<br />
製作：「闇芝居」製作委員会<br />
コピーライト表記：<span translate="no">&#x00A9;「闇芝居」製作委員会2021</span><br />
<br />
<b>・キャスト</b><br />
紙芝居屋：津田寛治<br />
木津つばさ<br />
深澤大河<br />
篠田諒<br />
園山敬介<br />
沢井正棋<br />
ぽんず<br />
沢田敏子<br />
二又一成<br />
森レイ子<br />
千本木彩花<br />
松田利冴<br />
山崎夏菜<br />
奈日抽ねね<br />
水科葵<br />
<br />
ED：忘れらんねえよ「くだらねえな」<br />
<br />
</p><hr></p><p class="preface"><h2 id="52">◆<b><a href="https://www.toei-anim.co.jp/tv/delicious-party_precure/" target="_blank">デリシャスパーティ♡プリキュア</a></b></h2><p class="preface">
</p><a href="https://www.toei-anim.co.jp/tv/delicious-party_precure/" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/52.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
ABC・テレビ朝日系列全国ネット：2月～<br />
<br />
<b>・作品情報</b><br />
プリキュアシリーズ第19作目。<br />
<br />
<b>・スタッフ</b><br />
アニメーション制作：東映アニメーション<br />
コピーライト表記：<span translate="no">&#x00A9;ABC-A・東映アニメーション</span><br />
<br />
</p><hr></p><p class="preface"><h2 id="57">◆<b><a href="https://www.toei.co.jp/tv/donbrothers/index.html" target="_blank">暴太郎戦隊ドンブラザーズ</a></b></h2><p class="preface">
</p><a href="https://www.toei.co.jp/tv/donbrothers/index.html" target="_blank"><img data-src="https://i.gzn.jp/img/2021/12/15/anime-2022winter/57.png" border="0" class="lazyload"></a><p class="preface">
<br />
<b>・放送情報</b><br />
テレビ朝日系列：3/6(日) 9:30～<br />
<br />
<b>・作品情報</b><br />
昔話「桃太郎」をモチーフとしたスーパー戦隊。<br />
<br />
<b>・スタッフ</b><br />
脚本：井上敏樹<br />
演出：田﨑竜太<br />
コピーライト表記：<span translate="no">&#x00A9;テレビ朝日・東映AG・東映</span><br />
<br />
<b>・キャスト</b><br />
ドンモモタロウ(レッド)：未発表<br />
サルブラザー(ブルー)：未発表<br />
イヌブラザー(ブラック)：未発表<br />
キジブラザー(ピンク)：未発表<br />
オニシスター(イエロー)：未発表<br />
<br />
Twitter：<a href="https://twitter.com/Donbro_toei" target="_blank">@Donbro_toei</a><br />
ハッシュタグ：<a href="https://twitter.com/hashtag/暴太郎戦隊ドンブラザーズ" target="_blank">#暴太郎戦隊ドンブラザーズ</a> <a href="https://twitter.com/hashtag/ドンブラザーズ" target="_blank">#ドンブラザーズ</a><br />
</p>`
