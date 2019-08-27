package pkg

// 要素がスライスに含まれているかの確認
func IsContain(s []string, e string) bool {
	if s == nil {
		return false
	}

	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

// 重複の確認
func Duplicate(s []string, c string) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}

	return false
}

func RemoveDuplicate(s []string) []string {
	l := []string{}
	for _, v := range s {
		if !Duplicate(l, v) {
			l = append(l, v)
		}
	}

	return l
}
