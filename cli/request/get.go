package request

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"net/http"
)

func GetRequestBasicAuth(s string) (*goquery.Document, error) {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil, err
	}

	config.AuthInit()
	req = pkg.BasicAuth(req, config.BasicAuth.UserName, config.BasicAuth.Passwords)

	res, err := pkg.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var doc *goquery.Document
	if res.StatusCode == 200 {
		doc, err = goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, err
		}
	}

	return doc, nil
}
