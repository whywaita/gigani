# gigani

parser of GIGAZINE Anime page

## Install

### Download binary

You can download binary in [release page](https://github.com/whywaita/gigani/releases).

### Build on your computer

```
$ git clone https://github.com/whywaita/gigani
$ go build .
$ ./gigani
```

## Usage

### Option

- `-url`: [required] set target URL (in GIGAZINE).
- `-output`: [required] set output format. markdown / json
- `-sort`: [optional] set sort rule. post(default) / time

## Example

```
$ ./gigani -url "http://gigazine.net/news/20170917-anime-2017autumn/" -output markdown
# 録画リスト
| 名前 | 局 | 放送時間 | done |
|---|---|---|---|
|[アイドルマスター SideM](http://imas-sidem.com/) |TOKYO MX|10/7(土) 23:30～||
|[Infini-T Force(インフィニティフォース)](http://www.infini-tforce.com/) |日本テレビ|10/3(火) 25:59～||
|[ラブライブ！サンシャイン!! 第2期](http://www.lovelive-anime.jp/uranohoshi/) |TOKYO MX|10/7(土) 22:30～||
|[アニメガタリズ](http://animegataris.com/) |TOKYO MX|10/8(日) 22:00～||
|[UQ HOLDER!～魔法先生ネギま！2～](http://uqholder.jp/) |tvk|10/2(月) 25:00～||
etc etc...
```

Output Example: https://gist.github.com/whywaita/2ecc3086f98b5831d9e92cc5c72b8db1
