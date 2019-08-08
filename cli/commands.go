package cli

import (
	"github.com/taqboz/tombo_gdn/cli/commands"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/cli/request"
	"github.com/taqboz/tombo_gdn/cli/target"
	"github.com/taqboz/tombo_gdn/cli/ui"
	"golang.org/x/sync/errgroup"
)

func seo(input string) error {
	err := common(input)
	if err != nil {
		return err
	}

	var eg errgroup.Group
	c := make(chan struct{}, config.SimultaneousAccess)

	// チェック項目設定の読み込み
	cl, err := config.CheckInit("config/items/seo.json")
	if err != nil {
		return err
	}

	for _, v := range target.PageList {
		v2 := v
		eg.Go(func() error {
			c<-struct{}{}
			defer func() {
				<-c
			}()

			doc, err := request.GetRequestBasicAuth(v2)
			if err != nil{
				return err
			}

			if err := commands.SEO(doc, cl); err != nil {
				return err
			}

			return nil
		})
	}

	// エラーハンドリング
	if err := eg.Wait(); err != nil {
		return err
	}

	ui.Finish()
	return nil
}
