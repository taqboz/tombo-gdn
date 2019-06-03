package view

import (
	"fmt"
	"sync"

	"github.com/taqboz/tombo_gdn/cli/check"
)

func ErrContents() {
	for _, v := range check.ErrContsList {
		var once sync.Once
		for _, v2 := range v.Tags {
			if v2.IsErr {
				//　エラーが会ったページのURLの表示
				once.Do(func() {
					fmt.Println("-----------------------------------------------------------------------------------")
					fmt.Println(v.URL)
				})
				// タグごとの結果の表示
				fmt.Println("---------------")
				fmt.Print("タグ："+v2.Tag+"."+v2.Target+"|")
				for k, v := range v2.Attr {
					fmt.Print(k+":"+v+",")
				}
				fmt.Println()
				cont(v2.Length, "Length")
				useKw(v2.UseKw)
				numMultiple(v2.NumMultiple)
				match(v2.Match, "Match：")
				match(v2.Include, "Include：")
				duplicate(v2.DuplicateInPage, "Duplicate")
				duplicate(v2.DuplicateInContent, "DuplicateInContent")
			}
		}
	}
	// エラーのあったページの数の表示
 	defer fmt.Print("\nErrors are ditected from pages:", len(check.ErrContsList), "page\n\n")

	duplicateAll()
}

// ページ全体でのコンテンツの重複エラーの表示
func duplicateAll()  {
	fmt.Println("\n" + "### Duplicate ###")
	for k, v := range check.ErrDupulicateList {
		var once sync.Once
		for k2, v2 := range v {
			if len(v2) > 1 {
				// タグと属性の表示
				once.Do(func() {
					fmt.Println("-----------------------------------------------------------------------------------")
					fmt.Println(k+":")
				})
				// エラー内容の表示
				fmt.Println(k2 + "|")
				for _, v3 := range v2 {
					fmt.Println(v3)
				}
				fmt.Println("-")
			}
		}
	}
}
