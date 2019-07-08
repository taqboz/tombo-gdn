package ui

import (
	"fmt"
	"github.com/taqboz/tombo_gdn/cli/target"
)

const (
	long = "-----------------------------------------------------------"
)
func SetTarget()  {
	fmt.Println("チェック対象：", target.Host)
}

func CheckStart() {
	fmt.Println(long)
	fmt.Println("ページ検知数: ", len(target.PathList))
	fmt.Println("調査対象ページ数: ", len(target.PageList))
	fmt.Println(long)
	fmt.Println("調査開始...")
}

func LoadConfigStart()  {
	fmt.Print("設定の読み込み開始")
	fmt.Print("...")
}

func LoadConfigSiteMap()  {
	fmt.Print("sitemap.xmlの読み込み開始")
	fmt.Print("...")
}

func Finish()  {
	fmt.Println("完了")
}
