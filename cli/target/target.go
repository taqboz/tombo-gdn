package target

import (
	"github.com/pkg/errors"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"net/url"
)

var Host string

func SetTarget(input string)(string, error) {
	// 入力された文字列のパース
	u, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	// 入力された内容がURLかどうか確認
	if u.Scheme == "" && u.Host == "" {
		return "", errors.New("This isn't URL!!!")
	}
	Host = input

	// 入力された内容が文字列かどうか
	s, err := pkg.ResolveURL(input, "/sitemap.xml")
	if err != nil {
		return "", err
	}

	return s, nil
}
