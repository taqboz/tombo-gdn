package pkg

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/taqboz/tombo_gdn/cli/config"
)

var Count int

// HTMLドキュメント(goquery)の取得とレスポンスのコードの収録
func GetHTTP(url string) (*goquery.Document, int, error) {
	var i int
	doc, stcode, err := getHTTP(url, &i)
	if err != nil {
		return nil, 0, err
	}

	return doc, stcode, nil
}

// 503レスポンスの回避と200以外のエラーの生成
func getHTTP(url string, i *int) (*goquery.Document, int, error){
	Count++
	// http情報の取得
	doc, status, err := getDocument(url)
	if err != nil {
		return nil, 0, err
	}

	if status == 200 {
		return doc, status, nil
		// 503だった場合は再リクエストを行う
	} else if status >= 500 && status <= 510  {
		*i++
		if *i < 1000 {
			return getHTTP(url, i)
		} else {
			return nil, status, nil
		}
	} else {
		return nil, status, nil
	}
}

// goquery.Documentの生成
func getDocument(url string) (*goquery.Document, int, error) {
	var doc *goquery.Document
	// HTMLページにリクエストを送る
	res, err := http.Get(url)
	if err != nil {
		// log.Fatal(err)
		return nil, 0, err
	}
	defer res.Body.Close()

	// goquery.Documentの生成
	if res.StatusCode == 200 {
		doc, err = goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			// log.Fatal(err)
			return nil, 0, err
		}
	}
	return doc, res.StatusCode, nil
}

func GetStatus(url string) (int, error) {
	var i int
	stcode, err := getStatus(url, &i)
	if err != nil {
		return 0, err
	}

	return stcode, nil
}

func getStatus(url string, i *int) (int, error) {
	Count++
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return res.StatusCode, err

	} else if res.StatusCode >= 500 && res.StatusCode <= 510  {
		*i++
		if *i < 1000 {
			time.Sleep(config.GetLinksTimeSleep * time.Millisecond)
			return getStatus(url, i)
		} else {
			return res.StatusCode, err
		}
	} else {
		return res.StatusCode, err
	}
}
