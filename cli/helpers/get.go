package helpers

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"io"
	"net/http"
)

func GetWithBAtoDocument(s string) (*goquery.Document, int, error){
	res, err := GetRequestBasicAuth(s)
	if err != nil {
		return nil, 0, err
	}

	switch res.StatusCode {
	case 200:
		doc, err := BodyToDocument(res.Body)
		if err != nil {
			return doc, 200, nil
		}
	}

	return nil, res.StatusCode, nil
}

// 基本認証付きGETリクエスト
func GetRequestBasicAuth(s string) (*http.Response, error) {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		return nil,  err
	}

	req = pkg.BasicAuth(req, config.BasicAuth.UserName, config.BasicAuth.Passwords)

	res, err := pkg.DoRequest(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return res, err
}

func BodyToDocument(body io.ReadCloser) (*goquery.Document, error){
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
