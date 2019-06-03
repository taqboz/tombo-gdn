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

func IsDuplicate(s []string, e string) (bool, error) {
	if s == nil {
		return false, errors.New("slice is nil")
	}

	var i int
	for _, v := range s {
		if e == v {
			i++
		}
	}

	if i > 1 {
		return true, nil
	}

	return false, nil
}