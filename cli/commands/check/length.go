package check

import (
	"github.com/taqboz/tombo_gdn/cli/config"
	"strings"
	"unicode/utf8"
)

func common(s string, i int, c config.CheckLength) *NumIncorrect {
	var isIncorrect bool
	switch c.Check {
	case "all":
		isIncorrect = i < c.Min || i > c.Max
	case "min":
		isIncorrect = i < c.Min
	case "max":
		isIncorrect = i > c.Max
	}

	if isIncorrect {
		return &NumIncorrect{s,i}
	}

	return nil
}

func Length(s string, c config.CheckLength, l []*NumIncorrect) []*NumIncorrect {
	i := utf8.RuneCountInString(s)
	add := common(s, i, c)
	if add != nil {
		l = append(l, add)
	}

	return l
}

func UseKws(s string, c config.CheckLength, l []*NumIncorrectList, kws []string) []*NumIncorrectList {
	add := &NumIncorrectList{s,[]*NumIncorrect{}}
	for _, v := range kws {
		i := strings.Count(s, v)
		cont := common(s, i, c)
		if cont != nil {
			add.Incorrect = append(add.Incorrect, cont)
		}
	}

	if len(add.Incorrect) > 0 {
		l = append(l, add)
	}

	return l
}

func MultipleCount(s string, c config.CheckLength, l []*NumIncorrect, split string) []*NumIncorrect {
	i := len(strings.Split(s, split))
	add := common(s, i, c)
	if add.Length > 0 {
		l = append(l, add)
	}
	return l
}
