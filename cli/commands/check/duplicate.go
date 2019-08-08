package check

import (
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"strings"
)

func MultipleDuplicate(s string, c config.CheckLength, split string) *NumIncorrectList {
	add := &NumIncorrectList{s,[]*NumIncorrect{}}
	agr := strings.Split(s, split)
	l := pkg.RemoveDuplicate(agr)
	for _, v := range l {
		n := DuplicateNum(v, agr)
		if n > 1 {
			cont := &NumIncorrect{v, n}
			add.Incorrect = append(add.Incorrect, cont)
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
