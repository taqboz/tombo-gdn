package check

import (
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"strings"
	"unicode/utf8"
)

func common(i int, c config.CheckLength) *LengthIncorrect {
	var sts string
	var dif int
	switch c.Check {
	case "all":
		if i < c.Min {
			sts = "not enough"
			dif = c.Min - i
		} else if i > c.Max {
			sts = "over"
			dif = i - c.Max
		}
	case "min":
		if i < c.Min {
			sts = "not enough"
			dif = c.Min - i
		}
	case "max":
		if i > c.Max {
			sts = "over"
			dif = i - c.Max
		}
	}

	if sts != "" {
		return &LengthIncorrect{i,sts, dif}
	}

	return nil
}

func Length(s string, c config.CheckLength) *LengthIncorrect {
	i := utf8.RuneCountInString(s)
	add := common(i, c)

	return add
}

func UseKws(s string, c config.CheckLength, kws []string) []*LengthIncorrectCont {
	add := []*LengthIncorrectCont{}
	for _, v := range kws {
		i := strings.Count(s, v)
		cont := common(i, c)
		if cont != nil {
			addCont := &LengthIncorrectCont{s,cont}
			add = append(add, addCont)
		}
	}

	return add
}

func MultipleNum(s string, c config.MultipleLength) *LengthIncorrect {
	i := len(strings.Split(s, c.Split))
	conf := config.CheckLength{c.Check,c.Min,c.Max}
	add := common(i, conf)

	return add
}

func MultipleDup(s string, c config.MultipleLength) []*NumIncorrect {
	add := []*NumIncorrect{}
	agr := strings.Split(s, c.Split)
	l := pkg.RemoveDuplicate(agr)
	for _, v := range l {
		n := DuplicateNum(v, agr)
		if n > 1 {
			cont := &NumIncorrect{v, n}
			add = append(add, cont)
		}
	}

	return add
}
