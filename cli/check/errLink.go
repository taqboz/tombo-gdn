package check

// エラーリンク格納用配列
var (
	ErrLinksList = []*ErrLinks{}
	BadLinks = []*ErrLink{}
	SuccessLinks = []string{}
)

// ページごとのエラーリンクの情報
type ErrLinks struct {
	URL string
	ErrLink []*ErrLink
}

// エラーリンクの情報
type ErrLink struct {
	URL string
	Status int
}

func isContainLink(list []*ErrLink, info string) (*ErrLink, bool) {
	for _, v := range list {
		if info == v.URL {
			return v, true
		}
	}
	return nil, false
}