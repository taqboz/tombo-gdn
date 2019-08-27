package helpers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"net/http"
	"net/url"
	"sort"
)

// 基本認証付きGETリクエスト
func GetRequestBasicAuth(s string) (*goquery.Document, int, error) {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil, 0,  err
	}

	req = pkg.BasicAuth(req, config.BasicAuth.UserName, config.BasicAuth.Passwords)

	res, err := pkg.DoRequest(req)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, 0, err
		}
		return doc, res.StatusCode, nil
	}

	return nil, res.StatusCode, nil

}

func ScrapingPath(doc *goquery.Document) []string {
	l := make([]string, 0)

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		c, ok := s.Attr("href")
		if ok {
			p, _ := url.Parse(c)
			if p.Path != "" {
				l = append(l, p.Path)
			}
		}
	})

	sort.Strings(l)
	return l
}


