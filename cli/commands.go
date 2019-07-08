package cli

import (
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/cli/request"
	"github.com/taqboz/tombo_gdn/cli/target"
	"golang.org/x/sync/errgroup"
)

func check(input string) error {
	err := common(input)
	if err != nil {
		return err
	}

	var eg errgroup.Group
	c := make(chan struct{}, config.SimultaneousAccess)

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

			if doc != nil {
				fmt.Println(v2)
			}

			return nil
		})
	}

	// エラーハンドリング
	if err := eg.Wait(); err != nil {
		return err
	}


	return nil
}
