package check

import "strings"

func NotPermit(s string, c []string) *StrIncorrect {
	add := &StrIncorrect{s, []string{}}
	for _, v := range c {
		if v == s {
			add.Match = append(add.Match, v)
		}
	}

	//if len(add.Match) > 0 {
	//	l = append(l, add)
	//}

	return add
}

func NotInclude(s string, c []string) *NumIncorrectList {
	add := &NumIncorrectList{s,[]*NumIncorrect{}}
	for _, v := range c {
		n := strings.Count(s, v)
		if n > 0 {
			cont := &NumIncorrect{v, n}
			add.Incorrect = append(add.Incorrect, cont)
		}
	}

	//if len(add.Incorrect) > 0 {
	//	l = append(l, add)
	//}

	return add
}

