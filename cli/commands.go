package cli

import (
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/request"
	"github.com/taqboz/tombo_gdn/cli/target"
)

func check(input string) error {
	err := common(input)
	if err != nil {
		return err
	}

	for _, v := range target.PageList {
		doc, err := request.GetRequestBasicAuth(v)
		if err != nil{
			return err
		}

		if doc != nil {
			fmt.Println(v)
		}
	}

	return nil
}
