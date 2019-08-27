package pkg

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// リクエストの実行
func DoRequest(req *http.Request) (*http.Response, error){
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetDocument(input string) (*goquery.Document, int, error) {
	var doc *goquery.Document

	res, err := http.Get(input)
	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		doc, err = goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			return nil, 0, err
		}
	}

	return doc, res.StatusCode, nil
}
