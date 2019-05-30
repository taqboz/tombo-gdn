package check

// エラー格納用配列
var ErrContsList = []*ErrList{}

// ページごとのエラー情報
type ErrList struct {
	URL string
	Tags []*ErrTag
}

// チェック項目ごとのエラー情報
type ErrTag struct {
	Tag string
	IsErr bool
	Target string
	Attr map[string]string
	Length []*ErrCont
	UseKw []*ErrUseKw
	NumMultiple *NumMultiple
	Match map[string]int
	Include map[string]int
	Duplicate []string
	DuplicateInPage map[string]int
	DuplicateInContent map[string]int
}

// エラー情報
type ErrCont struct {
	Content string
	Num int
}

type ErrUseKw struct {
	Content string
	Kw string
	Num int
}

type NumMultiple struct {
	IsErr bool
	Num int
}