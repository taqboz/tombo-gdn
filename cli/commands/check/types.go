package check

type IncorrectTag struct {
	Tag string
	Incorrect []*IncorrectCont
}

type IncorrectCont struct {
	Content string
	Length int
	UseKws []*NumIncorrect
	NotPermit []string
	NotInclude []*NumIncorrect
	MultipleNum int
	MultipleDuplicate []*NumIncorrect
}

type NumIncorrect struct {
	Content string
	Length int
}

type NumIncorrectList struct {
	Content string
	Incorrect []*NumIncorrect
}

type StrIncorrect struct {
	Content string
	Match []string
}
