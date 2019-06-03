# Tombo for GDN
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
$ brew tap taqboz/tombo_gdn
$ brew install tombo_gdn
````
* ビルド済みデータ：<br>以下のリンクからダウンロードしてください<br>
[【Dropbox】tombo_gdn (Windows 64bit)](https://www.dropbox.com/s/0qhxq3cos52tp4n/tombo_gdn.zip?dl=0)

### 起動方法：
* ターミナル等で以下のコードを実行 <br>
``$ tombo_gdn http://○○○.△△/``<br>

* もしくはビルド済みのファイルから起動。

#### オプション
`[半角スペース]調査対象のサイトマップ(.xml)URL`:すべての要素をチェック<br>
`-h`:バージョン情報、使い方の表示

### できること
* sitemap.xmlのチェック
* 文字数のチェック
* タグ内のコンテンツ内でのキーワードの最小使用回数、最大使用回数のチェック
* コンテンツが特定の文字列と合致するか、または含むかチェック
* ページ内でのコンテンツの重複チェック（alt属性の重複など）
* 全ページにおいてのコンテンツの重複チェック（titleの重複など）

##### メタキーワードなど複数要素を内包する場合、
* 重複のチェック
* 個数のチェック


## チェック項目・設定
### 設定ファイル
* config/tag_config.jsonに記入<br>
* 設定項目は以下の通り
```
[
  "check_page_parallel": 10, (並列アクセスの最大値)
  "retry_at_err": 1000, (503等エラー発生時、再度アクセスを
                         試みるまで待つ時間(単位：MicroSecond))

  // 以下、調査対象の設定
  "tags" : [
    {
      "tag" : "meta（タグ名）",
      "target": "content",　(取得したい情報が属性値の場合は設定)
      "attr": {
        "name": "keywords", (属性値での絞り込み)
      },
            
      "min" : 20（最小文字数）,
      "max" : 30（最大文字数）,
      
      "kw_min" : 1（キーワードの最小使用回数）,
      "kw_max" : 1（キーワードの最大使用回数）,
      
      "uniq_in_page" : true（ページ内での重複を検知）,
      "uniq" : true（全ページでの重複を検知）,
      
      "match" : [
        ""（ページ内での全文検索）
      ],
      
      include : [
        "" (ページ内での部分検索)
      ],
      
      // 調査対象が複数の要素から形成されている場合
      "multiple_content": {
        "duplicate_in_content": true,（要素の重複を検知）
        "split_point": ",",（区切り文字）
        "min": 2,（要素の最小数）
        "max": 7（要素の最大数）
      },
    },
    
    .....
    
  ]
]  
```

### 今後の機能追加予定
* Basic認証への対応
* チェックする要素の絞り込み機能
* ページ内部のエラーリンクのチェック
* GUI化 (希望が多ければ。。。)

## ライセンス
MIT license
