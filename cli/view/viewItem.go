package view

import (
	"fmt"
	"sync"

	"github.com/taqboz/tombo_gdn/cli/check"
)

// 文字数、キーワードの使用回数のエラー内容の表示
func cont(list []*check.ErrCont, msg string) {
	// エラーが存在するときに表示
	if len(list) > 0 {
		// 項目名
		fmt.Println(msg)
		for _, v := range list {
			// エラー内容の表示
			fmt.Println(v.Content+":",v.Num)
		}
		fmt.Println()
	}
}

// 文字の完全一致、部分一致のエラー内容の表示
func match(list map[string]int, msg string) {
	// エラーが存在するときに表示
	if len(list) > 0 {
		// 項目名
		fmt.Println(msg)
		for k, v := range list {
			// エラー内容の表示
			fmt.Print(`"`+k+`"：`)
			fmt.Println(v)
		}
		fmt.Println()
	}
}

// コンテンツの重複エラーの表示
func duplicate(list map[string]int, msg string) {
	var once sync.Once
	// エラーが存在するときに表示
	if len(list) > 0 {
		for k, v := range list {
			if v > 1 {
				// 項目名
				once.Do(func() {
					fmt.Println(msg)
				})
				// エラー内容の表示
				fmt.Print(`"`+k+`"：`)
				fmt.Println(v)
			}
		}
		fmt.Println()
	}
}

func useKw(list []*check.ErrUseKw) {
	var once sync.Once
	if len(list) > 0 {
		once.Do(func() {
			fmt.Println("Use Kws")
		})
		for _, v := range list {
			fmt.Println(v.Content)
			fmt.Println(v.Kw+":", v.Num)
		}
	}
}

func numMultiple(list *check.NumMultiple) {
	if list != nil {
		if list.IsErr {
			fmt.Println("Numbers:",list.Num)
		}
	}
}
