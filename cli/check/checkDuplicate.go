package check

import (
	"strings"

	"github.com/taqboz/tombo-gdn/cli/config"
	"github.com/taqboz/tombo-gdn/internal/app/tombo-gdn/pkg"
)

// ベージ内のコンテンツ重複チェック
func duplicateInPage(tag *ErrTag, duplicateInPage []string) error {
	for _, v := range duplicateInPage {
		if v != "" {
			ok, err := pkg.IsDuplicate(duplicateInPage, v)
			if err != nil {
				return err
			}

			switch {
			// 初期化
			case tag.DuplicateInPage == nil:
				tag.DuplicateInPage = map[string]int{}
				fallthrough

			case ok:
				tag.DuplicateInPage[v]++
				fallthrough

			case tag.DuplicateInPage[v] > 1:
				tag.IsErr = true
			}
		}
	}

	return nil
}

// コンテンツ内での重複の確認（メタキーワード、クラス等）
func duplicateInContent(tag *ErrTag, m *config.MultipleContent, cont string) error {
	if m.DuplicateInContent {
		// 配列に各コンテンツをいれる
		contents := strings.Split(cont, m.Split)

		for _, v := range contents {
			if v != "" {
				ok, err := pkg.IsContain(contents, v)
				if err != nil {
					return err
				}

				switch {
				// 初期化
				case tag.DuplicateInContent == nil:
					tag.DuplicateInContent = map[string]int{}
					fallthrough

				case ok:
					tag.DuplicateInContent[v]++
					fallthrough

				case tag.DuplicateInContent[v] > 1:
					tag.IsErr = true
				}
			}
		}

	}

	return nil
}

// 全ページにおけるコンテンツの重複チェック
func duplicate(c *config.Tag, cont string, url string) {
	if c.Duplicate {
		// タグ、属性等の情報の作成（キー)
		tagAttr := c.Tag + "." + c.Target + "|"
		for k, v := range c.Attr {
			tagAttr = tagAttr + k + ":" + v + ","
		}

		// エラー内容の登録
		if cont != "" {
			switch {
			// 初期化
			case ErrDupulicateList[tagAttr] == nil:
				ErrDupulicateList[tagAttr] = map[string][]string{}
				fallthrough

			case ErrDupulicateList[tagAttr][cont] == nil:
				ErrDupulicateList[tagAttr][cont] = []string{}
			}

			ErrDupulicateList[tagAttr][cont] = append(ErrDupulicateList[tagAttr][cont], url)
		}
	}
}
