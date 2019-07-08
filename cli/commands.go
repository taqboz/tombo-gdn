package cli

import (
	"fmt"
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

	for _, v := range target.PageList {
		eg.Go(func() error {
			doc, err := request.GetRequestBasicAuth(v)
			if err != nil{
				return err
			}

			if doc != nil {
				fmt.Println(v)
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
