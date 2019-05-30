package target

import (
	"errors"
	"log"

	"github.com/taqboz/tombo-gdn/internal/app/tombo-gdn/regex"
)

	// チェック対象
var	Input, Host, Option string

func InputTargetInfo(input string) (err error){
	// 入力されたのがURLかどうか判別する
	isXml := regex.XML.MatchString(input)

	// 未入力の場合
	if input == "" {
		err = errors.New("Enter WEB page's URL that you want to check.")

	// 引数がURLではなかった場合
	} else if !isXml {
		err = errors.New("This is not a WEB page's URL!\nplease enter another one")

	// 引数がURLだった場合
	} else {
		Input = input
		// ホストURLを抽出
		var extUrl string
		ext := regex.HostURL.FindAllString(input, -1)
		if len(ext) == 0 {
			log.Fatal("Cannot extract Host URl from the XML's URL input")
		}
		extUrl = ext[0]

		if extUrl[len(extUrl)-1:] != "/" {
			// URLの末尾"/"の補完
			extUrl = extUrl + "/"
		}

		Host = extUrl

	}

	return err
}
