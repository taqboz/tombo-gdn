package pkg

import (
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestKwsSplit(t *testing.T) {
	f, e := os.Open("../../../../testdata/gotesting.html")
	if e != nil {
		t.Fatal(e)
	}
	defer f.Close()

	// file, err := ioutil.ReadAll()
	doc, e := goquery.NewDocumentFromReader(f)
	if e != nil {
		log.Fatal(e)
	}

	a := []string{"長期","インターンシップ","営業","企画"}

	s, err := KwsSplit(doc)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range s {
		if v != a[k] {
			t.Fatal("failed test")
		}
	}
}