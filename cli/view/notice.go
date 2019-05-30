package view

import (
	"fmt"

	"github.com/taqboz/tombo/cli/target"
)

const split string = "\n###############################################################################\n"

func FromXml() {
	fmt.Println("getting information of target website: " + target.Input + "...\n")
}

// 調査開始前の表示
func InfoGot()  {
	fmt.Println("Check Target: " + target.Host)
	fmt.Println("Detected & Check pages: ", len(target.Urls))
	fmt.Print(split)
	fmt.Print("Checking...")
}

// 調査結果の表示
func Result()  {
	// 調査終了の通知
	fmt.Println("Completed")
	fmt.Println(split)
	// 調査結果の表示
	fmt.Print("<Check Result>\n\n")
}