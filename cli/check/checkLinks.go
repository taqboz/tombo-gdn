package check

import (
	"sync"

	"golang.org/x/sync/errgroup"
	"github.com/PuerkitoBio/goquery"

	"github.com/taqboz/tombo/cli/config"
	"github.com/taqboz/tombo/cli/target"
	"github.com/taqboz/tombo/internal/app/tombo/pkg"
	"github.com/taqboz/tombo/internal/app/tombo/regex"
)

// エラーリンクの確認
func ErrCheckLink (doc *goquery.Document) ([]*ErrLink, error) {
	list := []*ErrLink{}
	sccessMtx := &sync.Mutex{}
	errMtx := &sync.Mutex{}

	// httpリクエストを多量に行うため、goroutineを使う
	var eg errgroup.Group
	var mtx sync.Mutex
	// ページあたりのリンクチェックの並行数
	c := make(chan struct{}, config.GetLinksParallel)

	// リンクを収集する
	link, err := gatherLinks(doc)
	if err != nil {
		return nil, err
	}


	for _, v := range link {
		v2 := v

		eg.Go(func() error {
			c<-struct{}{}
			defer func(){
				<-c
			}()

			if info, ok := isContainLink(BadLinks, v2); ok {
				mtx.Lock()
				list = append(list, info)
				mtx.Unlock()

				// リストに乗っていない正しいリンク
			} else if c, err := pkg.IsContain(SuccessLinks, v2); c{
				if err != nil {
					return err
				}

			} else {
				stcode, err := pkg.GetStatus(v2)
				if err != nil {
					return err
				}

				switch stcode {
				case 200:
					sccessMtx.Lock()
					SuccessLinks = append(SuccessLinks, v2)
					sccessMtx.Unlock()

				default:
					// リストに乗っていないエラーリンクの処理
					errInfo := &ErrLink{v2, stcode}

					errMtx.Lock()
					BadLinks = append(BadLinks, errInfo)
					errMtx.Unlock()

					mtx.Lock()
					list = append(list, errInfo)
					mtx.Unlock()
				}
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return list, nil
}

// ページ内のリンクを絶対リンクとして収集する
func gatherLinks(doc *goquery.Document) ([]string, error) {
	bfLink := []string{}
	link := []string{}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		attr, _ := s.Attr("href")
		// 相対リンクかどうかをチェック
		isRltv := regex.Relative.MatchString(attr)
		if isRltv {
			attr = relativeToAbsolute(attr)
		}
		bfLink = append(bfLink, attr)
	})

	for _, v := range bfLink {
		// ページリンクかどうかをチェック
		// tel, e-mail等を除外
		isPl := regex.PageLink.MatchString(v)

		ok, err := pkg.IsContain(link, v)
		if err != nil {
			return nil, err
		}

		if !ok && isPl {
			link = append(link, v)
		}
	}

	return link, nil
}

// 相対URLを絶対URLに変換
func relativeToAbsolute (url string) string {
	host := target.Host
	// ホストURLの最後が"/"の場合、絶対リンク作成用に編集
	if host[len(host)-1:] == "/" {
		host = host[:len(host)-1]
	}

	return host + url
}