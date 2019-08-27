package pkg

import (
	"net/url"
)

// ページURLかどうか確認
func IsURL(input string) (bool, error) {
	u , err := url.Parse(input)
	if err != nil {
		return false, err
	}

	if u.Scheme != "" && u.Host != "" {
		return false, nil
	}

	return true, nil
}

// 相対リンクを絶対リンクにする処理
func ResolveURL(input string, path string) (string, error) {
	base, err := url.Parse(input)
	if err != nil {
		return "", err
	}

	reference, err := url.Parse(path)
	endpoint := base.ResolveReference(reference).String()

	return endpoint, nil
}
