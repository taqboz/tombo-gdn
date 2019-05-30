package check

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/taqboz/tombo/cli/config"
)

// 文字列の長さのエラーの確認
func length(tag *ErrTag, c *config.Tag, cont string) {
	contLen := utf8.RuneCountInString(cont)

	isMin := c.Min > 0
	isMax := c.Max > 0

	addLengthErr := &ErrCont{omission(cont), contLen}

	switch {
	case isMin && isMax && (contLen < c.Min || contLen > c.Max):
		tag.Length = append(tag.Length, addLengthErr)

	case isMin && contLen < c.Min:
		tag.Length = append(tag.Length, addLengthErr)

	case isMax && contLen > c.Max:
		tag.Length = append(tag.Length, addLengthErr)
	}

}

// kw使用回数チェック
func kwCheck(tag *ErrTag,  c *config.Tag, cont string, keywords []string) {
	if cont != "" {
		isMin := c.KwMax > 0
		isMax := c.KwMax > 0

		if isMin || isMax {
			for _, v := range keywords {
				r := regexp.MustCompile(v)
				fs := r.FindAllString(cont, -1)
				num := len(fs)

				addUseKwErr := &ErrUseKw{omission(cont),v, num}

				switch {
				case isMin && isMax && (num < c.KwMin || num > c.KwMax):
					tag.UseKw = append(tag.UseKw, addUseKwErr)

				case isMin && num < c.KwMin:
					tag.UseKw = append(tag.UseKw, addUseKwErr)

				case isMax && num > c.KwMax:
					tag.UseKw = append(tag.UseKw, addUseKwErr)
				}
			}
		}
	}
}

func omission(cont string) string {
	// 長すぎる場合は短くする
	if utf8.RuneCountInString(cont) < 50 {
		return cont
	}

	return string([]rune(cont)[0:50]) + "..."
}

func multipleContent(tag *ErrTag, m *config.MultipleContent, cont string) {
	if m.DuplicateInContent {
		mc := strings.Split(cont, m.Split)
		num := len(mc)

		if cont == "" {
			num = 0
		}

		isMin := m.Min > 0
		isMax := m.Max > 0

		if tag.NumMultiple == nil {
			tag.NumMultiple = &NumMultiple{}
		}

		addErr := &NumMultiple{true, num}

		switch {
		case isMin && isMax && (num < m.Min || num > m.Max):
			tag.NumMultiple = addErr

		case isMin && num < m.Min:
			tag.NumMultiple = addErr

		case isMax && num > m.Max:
			tag.NumMultiple = addErr
		}
	}
}

// 特定の文字列と完全一致するかチェック
func match(tag *ErrTag, match []string, cont string) {
	for _, v := range match {
		if cont == v{
			// 初期化
			if tag.Match == nil {
				tag.Match = map[string]int{}
			}
			// エラー内容の登録
			tag.Match[v]++
			// エラーがあることを記録
			if tag.Match[v] > 1 {
				tag.IsErr = true
			}
		}
	}

}

// 特定の文字列と部分一致するかチェック
func include(tag *ErrTag, include []string, cont string) {
	for _, v := range include {
		// 正規表現
		r := regexp.MustCompile(v)
		match := r.FindAllString(cont, -1)
		if len(match) > 0 {
			// 初期化
			if tag.Include == nil {
				tag.Include = map[string]int{}
			}
			// エラー内容の登録
			tag.Include[v] = len(match)
			// エラーがあることを記録
			if tag.Include[v] > 1 {
				tag.IsErr = true
			}
		}
	}
}