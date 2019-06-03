package check

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/config"
)

// ページごとエラーコンテンツの調査
func ErrCheckContents(url string, doc *goquery.Document) (*ErrList, error) {
	// 指定ページ用エラーリストを作成してURLを登録
	list := &ErrList{URL:url}

	// 設定ファイルを元にスクレイピングを行う
	for _, c := range config.Tags {
		// タグ毎に調査を行いエラーリストを作成
		l, err := checkTag(doc, c, url)
		if err != nil {
			return nil, err
		}

		// ページごとのエラーリストに作成したエラーリストを登録
		list.Tags = append(list.Tags, l)

	}

	return list, nil
}
