package check

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"

	"github.com/taqboz/tombo-gdn/cli/config"
)

func TestErrCheckContents(t *testing.T) {
	t.Helper()

	config.AsignConfig("../../testdata/config.json")

	tests := []struct{
		url string
		doc string
		out ErrList
	}{
		{"http://career-navigation.co.jp/","../../testdata/gotesting.html",
			ErrList{URL:"http://career-navigation.co.jp/"}},
	}

	for _, v := range tests {
		f, e := os.Open(v.doc)
		if e != nil {
			t.Fatal(e)
		}
		defer func(){
			err := f.Close()
			if err != nil {
				t.Fatal(err)
			}
		}()

		doc, e := goquery.NewDocumentFromReader(f)
		if e != nil {
			t.Fatal(e)
		}

		list, err := ErrCheckContents(v.url, doc)
		if err != nil {
			log.Fatal(err)
		}

		if reflect.DeepEqual(list, v.out) {
			t.Fatal("failed test")
		}
	}

}
