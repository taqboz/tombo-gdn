package pkg

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// メタキーワードを取得して単語ごとに分割
func KwsSplit(doc *goquery.Document) ([]string, error) {
	if doc == nil {
		return nil, errors.New("KwsSplit: documents is nil")
	}

	// メタキーワードの取得
	var metaKeywords string
	doc.Find("meta").Each(func(_ int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "keywords" {
			attr, _ := s.Attr("content")
			metaKeywords = attr
		}
	})

	// 分割して配列に格納
	keywords := []string{}
	if metaKeywords == "" {
		keywords = strings.Split(metaKeywords, ",")
	}

	return keywords, nil
}
