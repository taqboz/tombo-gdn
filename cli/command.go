package cli

import (
	"errors"
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/helpers"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
)

func check(input string) error {
	// 設定の読み込み
	if err := common(input); err != nil {
		return err
	}

	// 入力されたURLの確認
	// TODO

	// Homeディレクトリからリンクを取得
	doc, _, err := helpers.GetWithBAtoDocument(input)
	if err != nil {
		return err
	}

	if doc == nil {
		return errors.New("there is no content in top page")
	}

	links := pkg.RemoveDuplicate(pkg.Scraping(doc, "a", "href", nil))

	for _, v := range links {
		fmt.Println(v)
	}

	// 各ディレクトリにアクセスしてリンクチェックを開始
	// TODO

	// 結果の表示
	// TODO


	return nil
}
