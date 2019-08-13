package check

import "strings"

func NotPermit(s string, c []string) *MatchIncorrect {
	for _, v := range c {
		if v == s {
			return &MatchIncorrect{v}
		}
	}

	return nil
}

func NotInclude(s string, c []string) []*NumIncorrect {
	add := []*NumIncorrect{}
	for _, v := range c {
		n := strings.Count(s, v)
		if n > 0 {
			cont := &NumIncorrect{s, n}
			add = append(add, cont)
		}
	}

	return add
}

