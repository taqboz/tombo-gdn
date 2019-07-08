package cli

import (
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/cli/request"
	"github.com/taqboz/tombo_gdn/cli/target"
)

func common(input string) error {
	// コンフィグの読み込み
	if err := loadConfig(); err != nil {
		return err
	}

	// ターゲットとsitemap.xmlのパスの設定
	siteMap, err := target.SetTarget(input)
	if err != nil {
		return err
	}

	// goqueryDocument（sitemap.xml）の作成
	doc, err := request.GetRequestBasicAuth(siteMap)
	if err != nil {
		return err
	}

	// 存在するパスをリスト化
	if err := target.GetPaths(doc); err != nil {
		return err
	}

	return nil
}

func loadConfig() error {
	// 基本認証の設定
	if err := config.AuthInit("config/auth.ini"); err != nil {
		return nil
	}

	// 調査対象外の項目の設定
	if err := config.IgnoreInit("config/ignore.json"); err != nil {
		return err
	}

	return nil
}
