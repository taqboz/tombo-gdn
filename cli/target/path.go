package target

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"net/url"
	"sort"
)

var (
	PathList = []string{}
	PageList = []string{}
)

// 調査対象のページを読み込む
func GetPaths(doc *goquery.Document) error {
	l := pkg.Scraping(doc, "loc","", nil)
	for _, v := range l {
		u, err := url.Parse(v)
		if err != nil {
			return err
		}

		PathList = append(PathList, u.Path)

		p, err := pkg.ResolveURL(Host, u.Path)
		if err != nil {
			return err
		}

		if !pkg.IsContain(config.Ignore.IgnorePaths, u.Path) {
			PageList = append(PageList, p)
		}
	}

	sort.Strings(PathList)

	return nil
}
