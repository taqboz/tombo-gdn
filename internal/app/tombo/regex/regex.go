package regex

import (
	"regexp"
)

var (
	HostURL = regexp.MustCompile(`^https?://(([\w-]+\.)+[\w-]+/?|localhost:([\d-]*)+/?)`) // ホストURL
	PageLink  = regexp.MustCompile(`^https?://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$|(localhost:[\d-]*)+(/[\w-./?%&=]*)$`) // ページリンク
	XML = regexp.MustCompile(`^https?://(([\w-]+\.)+[\w-]+(/[\w-./?%&=]*).xml$|localhost:+([\d-]*)+(/[\w-]*).xml$)`)
	Relative = regexp.MustCompile(`^/(?)`) // 相対リンク
)