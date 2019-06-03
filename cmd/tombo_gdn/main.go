package tombo_gdn

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/taqboz/tombo_gdn/cli"
)

func Main() {
	// 実行
	var sc = bufio.NewScanner(os.Stdin)
	input := os.Args
	if len(input) == 1 {
		fmt.Println("[半角スペース] or ｢--｣ + [サイトマップ]で入力してください。")
		//fmt.Println("-- :すべてのコンテンツのみチェックを行います")
		//fmt.Println("--meta -m    head要素のみチェックを行います")
		//fmt.Println("--body, -b   body要素のみチェックを行います")
		fmt.Println("※「-help」もしくは「-h] + と入力するとヘルプを表示します")
		fmt.Println("----------------------------------------------------------")
		fmt.Println("例: -- https://guardian.jpn.com/sitemap.xml")
		fmt.Print("Please Enter -[Option] [Sitemap's URL]: ")
		sc.Scan()
		input = strings.Split(sc.Text(), " ")
	}

	// 時間の測定の開始
	bf := time.Now()
	if err := cli.New().Run(input); err != nil {
		log.Fatal(err)
	}

	// 時間測定終了
	af := time.Now()
	fmt.Println(af.Sub(bf))
}
