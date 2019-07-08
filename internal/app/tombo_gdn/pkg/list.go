package pkg

import (
	"errors"
)

// 要素がスライスに含まれているかの確認
func IsContain(s []string, e string) (bool, error) {
	if s == nil {
		return false, errors.New("slice is nil")
	}

	for _, v := range s {
		if e == v {
			return true, nil
		}
	}
	return false, nil
}
