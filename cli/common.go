package cli

import (
	"github.com/taqboz/tombo_gdn/cli/config"
)

func common(input string) error {
	// 基本認証の設定
	if err := config.LoadInit("config/config.ini"); err != nil {
		return err
	}

	return nil
}
