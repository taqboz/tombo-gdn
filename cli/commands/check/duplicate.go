package check

import "github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"

func Duplicate(s []string) []*NumIncorrect {
	add := []*NumIncorrect{}
	agr := pkg.RemoveDuplicate(s)
	for _, v := range agr {
		n := DuplicateNum(v, s)
		if n > 1 {
			add = append(add, &NumIncorrect{v, n})
		}
	}

	return add
}

func DuplicateNum(t string, s []string) int {
	i := 0
	for _, v := range s {
		if v == t {
			i++
		}
	}

	return i
}
