# tombo
## 概要
[taqboz/tombo: WEBサイトのmeta情報、bodyの中身を高速でチェックするツール](https://github.com/taqboz/tombo/)のガーディアン社員向け改修版です。<br>
ホームページ制作後の品質チェック用CLIです。<br>
入稿情報、META情報等のチェックをページごとに検知した修正が必要な項目を表示します。


## 使い方
### インストール
* ソースコードを直接扱う場合
````
$ go get github.com/taqboz/tombo
````
* Homebrewを扱う場合
````
$ brew tap taqboz/tombo-gdn
$ brew install tombo-gdn
````
* ビルド済みデータ：<br>以下のリンクからダウンロードしてください 

### 起動方法：
* ターミナル等で以下のコードを実行 <br>
``$ tombo http://○○○.△△/``<br>
* もしくはビルド済みのファイルから起動。

#### オプション
`調査対象のURL`:すべての要素をチェック<br>
`-h 調査対象のURL`:head要素のみチェック<br>
`-b 調査対象のURL`:body要素のみチェック<br>

`-h`:バージョン情報、使い方の表示

### できること
* sitemap.xmlのチェック
* 文字数のチェック
* タグ内のコンテンツ内でのキーワードの最小使用回数、最大使用回数のチェック
* コンテンツが特定の文字列と合致するか、または含むかチェック
* ページ内でのコンテンツの重複チェック（alt属性の重複など）
* 全ページにおいてのコンテンツの重複チェック（titleの重複など）

## チェック項目・設定
### 設定ファイル
* config/tag_config.jsonに記入<br>
* 設定項目は以下の通り
```
[
  {
    "tag" : "title（タグ名）",
    "property" : "",（属性の値を調べる場合は設定）
    "min" : 20（最小文字数）,
    "max" : 30（最大文字数）,
    "kw_min" : 1（キーワードの最小使用回数）,
    "kw_max" : 1（キーワードの最大使用回数）,
    "uniq_in_page" : true（ページ内での重複を検知）,
    "uniq" : true（全ページでの重複を検知）,
    "not_match" : [
      ""（ページ内での全文検索）
    ]
  },
  
  {
    "tag" : "h1",
    ...
  }
]  
```

### 今後の機能追加予定
* ページ内部のエラーリンクのチェック
* GUI化 (希望が多ければ。。。)

## ライセンス
MIT license
