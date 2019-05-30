package config

import "time"

// 設定
var (
	Tags []*Tag
	CheckPageParallel int
	GetLinksParallel int
	GetLinksTimeSleep time.Duration
)

type Config struct {
	CheckPageParallel int `json:"check_page_parallel"`
	GetLinksParallel int `json:"get_links_parallel"`
	GetLinksTimeSleep time.Duration `json:"get_links_time_sleep"`
	Tag []*Tag `json:"tag"`
}

// チェック設定の構造体の定義
type Tag struct {
	Tag string `json:"tag"`
	Target string `json:"target"`
	Attr map[string]string `json:"attr"`
	Min int `json:"min"`
	Max int `json:"max"`
	KwMin int `json:"kw_min"`
	KwMax int `json:"kw_max"`
	MultipleContent MultipleContent `json:"multiple_content"`
	Duplicate bool `json:"duplicate"`
	DuplicateInPage bool `json:"duplicate_in_page"`
	Include []string `json:"include"`
	Match []string `json:"match"`
}

type MultipleContent struct {
	DuplicateInContent bool `json:"duplicate_in_content"`
	Split string `json:"split_point"`
	Min int `json:"min"`
	Max int `json:"max"`
}