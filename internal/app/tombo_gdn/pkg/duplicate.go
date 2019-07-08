package pkg

// 重複の確認
func Duplicate(s []string, c string) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}

	return false
}


