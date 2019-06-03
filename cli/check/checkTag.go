package check

import (
	"sync"

	"github.com/PuerkitoBio/goquery"

	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/cli/target"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
)

var MutexDuplicate = &sync.Mutex{}

func checkTag(doc *goquery.Document, c *config.Tag, url string) (*ErrTag, error) {
	// エラーリストにタグを登録
	tag := &ErrTag{Tag:c.Tag, Target:c.Target, Attr:c.Attr}

	// メタキーワードの取得
	kws, err := pkg.KwsSplit(doc)
	if err != nil {
		return tag, err
	}

	dip := []string{}

	var err2 error
	// スクレイピングした項目ごとにチェックを行う
	doc.Find(target.Option+c.Tag).Each(func(_ int, s *goquery.Selection) {
		// タグ内のテキストコンテンツ
		if attrSpec(c, s) {
			if c.Target == "" {
				err2 = checkItem(tag, c, s.Text(), kws, url, dip)
				// name属性を指定する場合で、属性値を取得する場合
			} else {
				cont, _ := s.Attr(c.Target)
				err2 = checkItem(tag, c, cont, kws, url, dip)
			}
		}
	})
	if err2 != nil {
		return nil, err2
	}


	if c.DuplicateInPage {
		err := duplicateInPage(tag, dip)
		if err != nil {
			return nil, err
		}
	}

	return tag, nil
}

func attrSpec(c *config.Tag, s *goquery.Selection) bool {
	for k, v := range c.Attr {
		attr, _ := s.Attr(k)
		if attr != v {
			return false
		}
	}
	return true
}

func checkItem(tag *ErrTag, c *config.Tag, cont string, keywords []string, url string,
	dip []string) error {
	// コンテンツが要件を満たしているか確認
	length(tag, c, cont)
	kwCheck(tag, c, cont, keywords)
	match(tag, c.Match, cont)
	include(tag, c.Include, cont)

	// コンテンツのページ内の重複の確認
	if c.DuplicateInPage {
		dip = append(dip, cont)
	}

	// コンテンツ内での重複の確認（メタキーワード、クラス等）
	if c.MultipleContent.DuplicateInContent {
		err := duplicateInContent(tag, &c.MultipleContent, cont)
		if err != nil {
			return err
		}
	}
	multipleContent(tag, &c.MultipleContent, cont)

	MutexDuplicate.Lock()
	duplicate(c, cont, url)
	MutexDuplicate.Unlock()

	return nil
}
