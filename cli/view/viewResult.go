package view

import (
	"fmt"
	"sync"

	"github.com/taqboz/tombo/cli/check"
	"github.com/taqboz/tombo/internal/app/tombo/pkg"
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
				// viewCont(v2.UseKw, "UseKw")
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
				if len(k2) > 150 {
					fmt.Println(k2[0:150] + "...")
				} else {
					fmt.Println(k2 + "|")
				}
				for _, v3 := range v2 {
					fmt.Println(v3)
				}
				fmt.Println("-")
			}
		}
	}
}

// リンクのエラー内容の表示
func ErrLinks() {
	fmt.Println("\n" + "### Error Links ###")
	for _, v := range check.ErrLinksList{
		fmt.Println(v.URL)
		for _, v2 := range v.ErrLink {
			// エラー内容の表示
			fmt.Println(">", v2.URL, ":", v2.Status)
		}
		fmt.Println()
	}

	fmt.Println("access:", pkg.Count)
	defer fmt.Println("Errors are ditected from pages:", len(check.ErrLinksList), "page")
}