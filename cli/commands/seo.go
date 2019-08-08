package commands

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
)

func SEO(doc *goquery.Document, cl []config.CheckItem) error {
	for _, v := range cl {
		l := pkg.Scraping(doc, v.Tag, v.Target, v.Attr)
		for _, v2 := range l {
			fmt.Println(v2)
		}
	}

	return nil
}
