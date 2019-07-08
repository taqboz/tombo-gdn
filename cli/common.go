package cli

import (
	"github.com/taqboz/tombo_gdn/cli/request"
	"github.com/taqboz/tombo_gdn/cli/target"
)

func common(input string) error {
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
