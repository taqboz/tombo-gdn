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
