package target

import (
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo-gdn/internal/app/tombo-gdn/pkg"
)

func TestGetFromXml(t *testing.T) {
	f, e := os.Open("../../testdata/sitemap.htm")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()
	// file, err := ioutil.ReadAll()
	doc, e := goquery.NewDocumentFromReader(f)

	a := []string{
		"http://career-navigation.co.jp/",
		"http://career-navigation.co.jp/about/",
		"http://career-navigation.co.jp/award/",
		"http://career-navigation.co.jp/blog/",
		"http://career-navigation.co.jp/blog/for_company/",
		"http://career-navigation.co.jp/blog/for_student/",
		"http://career-navigation.co.jp/event/",
		"http://career-navigation.co.jp/https://goo.gl/forms/QBOP2o9eh7JArOxH2/",
		"http://career-navigation.co.jp/https://goo.gl/forms/QBOP2o9eh7JArOxH2/complete.html",
		"http://career-navigation.co.jp/https://goo.gl/forms/QBOP2o9eh7JArOxH2/confirm.html",
		"http://career-navigation.co.jp/internship/",
		"http://career-navigation.co.jp/internship_blog/",
		"http://career-navigation.co.jp/internship_blog/sample/",
		"http://career-navigation.co.jp/internship_blog/sample2/",
		"http://career-navigation.co.jp/interview/",
		"http://career-navigation.co.jp/lp/",
		"http://career-navigation.co.jp/lp/index02.html",
		"http://career-navigation.co.jp/lp/index03.html",
		"http://career-navigation.co.jp/lp/thanks.html",
		"http://career-navigation.co.jp/lp01/",
		"http://career-navigation.co.jp/lp02/",
		"http://career-navigation.co.jp/lp03/",
		"http://career-navigation.co.jp/lp04/",
		"http://career-navigation.co.jp/lp05/",
		"http://career-navigation.co.jp/performance/",
		"http://career-navigation.co.jp/privacy_policy/",
		"http://career-navigation.co.jp/sitemap/",
		"http://career-navigation.co.jp/testpage/",
	}

	if e != nil {
		log.Fatal(e)
	}
	if err := getFromXml(doc); err != nil {
		t.Fatal(err)
	}

	for _, v := range Urls {
		ok, err := pkg.IsContain(a, v)
		if err != nil {
			log.Fatal(err)
		}

		if !ok {
			t.Fatal("failed test:",v)
		}
	}

}
