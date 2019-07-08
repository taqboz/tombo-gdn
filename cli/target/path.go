package target

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"net/url"
	"sort"
)

var (
	PathList = []string{}
	PageList = []string{}
)

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
		PageList = append(PageList, p)
	}

	sort.Strings(PathList)

	fmt.Println(PathList)

	return nil
}
