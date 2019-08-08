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

func Length(s string, c config.CheckLength) *NumIncorrect {
	i := utf8.RuneCountInString(s)
	add := common(s, i, c)

	return add
}

func UseKws(s string, c config.CheckLength, kws []string) *NumIncorrectList {
	add := &NumIncorrectList{s,[]*NumIncorrect{}}
	for _, v := range kws {
		i := strings.Count(s, v)
		cont := common(v, i, c)
		if cont != nil {
			add.Incorrect = append(add.Incorrect, cont)
		}
	}

	return add
}

func MultipleCount(s string, c config.CheckLength, split string) *NumIncorrect {
	i := len(strings.Split(s, split))
	add := common(s, i, c)

	return add
}
