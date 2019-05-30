package target

import (
	"errors"

	"github.com/PuerkitoBio/goquery"

	"github.com/taqboz/tombo/internal/app/tombo/pkg"
)

	// URLのリスト
var	Urls = []string{}

// URLリストの作成
func GetPageDirectory() error {
	doc, stcode, err := pkg.GetHTTP(Input)
	if err != nil{
		return err
	}

	switch stcode {
	case 200:
		if err := getFromXml(doc); err!= nil {
			return err
		}
		
	case 530:
		return errors.New("access to "+ Input +", but 503 responce returned, I'll continue check page")

	default:
		return errors.New("There is not simemap.xml in host, I'll continue check page")
	}

	return nil
}

// XMLからURLリストの取得
func getFromXml(doc *goquery.Document) error {
	// locタグからURLを取得する
	doc.Find("loc").Each(func(_ int, s *goquery.Selection) {
		Urls = append(Urls, s.Text())
	})

	// サイトマップにページが存在しない場合
	if len(Urls) == 0 {
		err := errors.New("There is no page")
		return err
	}

	return nil
}