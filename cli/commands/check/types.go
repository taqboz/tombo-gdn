package check

type IncorrectTag struct {
	Tag string
	Incorrect []*IncorrectCont
	Duplicate []*NumIncorrect
}

type IncorrectCont struct {
	Content string
	Length *LengthIncorrect
	UseKws []*LengthIncorrectCont
	NotPermit *MatchIncorrect
	NotInclude []*NumIncorrect
	MultipleNum *LengthIncorrect
	MultipleDup []*NumIncorrect
}

type LengthIncorrect struct {
	Length int
	Status string
	Difference int
}

type LengthIncorrectCont struct {
	Content string
	Incorrect *LengthIncorrect
}

type MatchIncorrect struct {
	Content string
}

type NumIncorrect struct {
	Content string
	Num int
}

type PageCont struct {
	Tag string
	Attr map[string]string
	Contents *[]NumIncorrect
}

type AllPageCont struct {
	Tag string
	Attr map[string]string
	Contents map[string][]string
}

type Cont struct {
	Content string
	Pages []string
}
