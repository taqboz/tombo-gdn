package data

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/helpers"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"sync"
)

var (
	CheckedPaths = make([]*Path, 0)
	MtxChecked = &sync.Mutex{}
	NotCheckedPaths = make([]string, 0)
	MtxNotChecked = &sync.Mutex{}

)

type Path struct {
	Path string
	Status int
	Doc *goquery.Document
	Checked bool
	ErrLinks []string
}

func AddPath(path *Path) {
	MtxChecked.Lock()
	for _, v := range CheckedPaths {
		if v.Path == path.Path {
			return
		}
	}
	CheckedPaths = append(CheckedPaths, path)
	fmt.Println("add" , path.Path)
	MtxChecked.Unlock()
}

func FindNotCheckedPath() {
	l := make([]string, 0)
	for _, v := range CheckedPaths {
		if !v.Checked && v.Doc != nil {
			a := helpers.ScrapingPath(v.Doc)
			for _, v := range a {
				l = append(l, v)
			}
			v.Checked = true
		}
	}

	l = pkg.RemoveDuplicate(l)

	for _, v := range l {
		isChecked, _ := IsChecked(v)
		if !isChecked {
			NotCheckedPaths = append(NotCheckedPaths, v)
		}
	}
}

func IsChecked(s string) (bool, int) {
	for _, v := range CheckedPaths {
		if v.Path == s {
			return true, v.Status
		}
	}

	return false, 0
}
