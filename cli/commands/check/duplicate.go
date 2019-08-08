package check

import (
	"github.com/taqboz/tombo_gdn/cli/config"
	"strings"
)

func MultipleDuplicate(s string, c config.CheckLength, l []*NumIncorrectList, split string) []*NumIncorrectList {
	add := &NumIncorrectList{s,[]*NumIncorrect{}}
	agr := strings.Split(s, split)
	for _, v := range agr {
		n := DuplicateNum(v, agr)
		if n > 1 {
			cont := &NumIncorrect{v, n}
			add.Incorrect = append(add.Incorrect, cont)
		}
	}

	if len(add.Incorrect) > 0 {
		l = append(l, add)
	}

	return l
}

func DuplicateNum(t string, s []string) int {
	i := 1
	for _, v := range s {
		if v == t {
			i++
		}
	}

	return i
}
