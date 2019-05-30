package cli

import (
	"fmt"
	"log"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/taqboz/tombo-gdn/cli/check"
	"github.com/taqboz/tombo-gdn/cli/config"
	"github.com/taqboz/tombo-gdn/cli/target"
	"github.com/taqboz/tombo-gdn/cli/view"
	"github.com/taqboz/tombo-gdn/internal/app/tombo-gdn/pkg"
)

var mtCont, mtLink *sync.Mutex

// エラーコンテンツ、エラーリンク両方のチェック
func process(o string) error {
	config.AsignConfig("config/config.json")
	// httpリクエストを多量に行うため、goroutineを使う
	var eg errgroup.Group

	c := make(chan struct{}, config.CheckPageParallel)
	for _, v := range target.Urls {
		v2 := v
		eg.Go(func() error {
			c<-struct{}{}
			defer func(){
				<-c
			}()

			doc, stcode, err := pkg.GetHTTP(v2)
			if err != nil {
				return err
			}

			if stcode == 200 {
				switch o {
				case "contents":
					// エラーコンテンツの調査
					err := contents(doc, v2)
					if err != nil {
						return err
					}

				default:
					return errors.New("this option is error")
				}

			} else {
				fmt.Print("\n"+v2+":", stcode, "\n")
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	view.Result()
	switch o {
	case "contents":
		view.ErrContents()

	case "links":
		view.ErrLinks()

	case "all":
		view.ErrContents()
		view.ErrLinks()

	default:
		log.Fatal("this option is error")
	}

	return nil
}

func contents(doc *goquery.Document, v string) error {
	c, err := check.ErrCheckContents(v, doc)
	if c != nil {
		return err
	}

	mtCont.Lock()
	check.ErrContsList = append(check.ErrContsList, c)
	mtCont.Unlock()
	return nil
}
