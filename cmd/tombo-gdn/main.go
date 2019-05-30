package tombo_gdn

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/taqboz/tombo-gdn/cli"
)

func Main() {
	// 時間の測定の開始
	bf := time.Now()

	// 実行
	if err := cli.New().Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// 時間測定終了
	af := time.Now()
	fmt.Println(af.Sub(bf))
}
