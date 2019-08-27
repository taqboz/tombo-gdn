package cli

import (
	"errors"
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/commands"
	"github.com/taqboz/tombo_gdn/cli/data"
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
	fmt.Printf("\n調査対象の設定：%s\n", input)

	// Homeディレクトリからリンクを取得
	doc, _, err := helpers.GetRequestBasicAuth(input)
	if err != nil {
		return err
	}

	if doc == nil {
		return errors.New("there is no content in top page")
	}


	data.NotCheckedPaths = pkg.RemoveDuplicate(helpers.ScrapingPath(doc))

	// 各ディレクトリにアクセスしてリンクチェックを開始
	if err := commands.Check(input, 2); err != nil {
		return err
	}

	// 結果の表示
	fmt.Printf("\n%d件のエラーリンクを検出\n", len(commands.ErrLinks))
	for _, v := range commands.ErrLinks {
		fmt.Println(v.Destination.Path, ":", v.Destination.Status)
		for _, v2 := range v.Sources {
			fmt.Println(v2)
		}
	}


	return nil
}
