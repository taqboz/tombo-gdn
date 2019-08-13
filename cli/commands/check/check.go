package check

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/taqboz/tombo_gdn/cli/config"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
)

func CheckAllCondetions(doc *goquery.Document, cl []config.CheckItem, u string) ([]*IncorrectTag, error) {
	add := []*IncorrectTag{}
	kws := pkg.Scraping(doc, "meta","contents" ,map[string]string{"name" : "keywords"})
	for _, v := range cl {
		a := &IncorrectTag{Tag:v.Tag}
		l := pkg.Scraping(doc, v.Tag, v.Target, v.Attr)
		a.Duplicate = Duplicate(l)
		agr := pkg.RemoveDuplicate(l)

		for _, v2 := range agr {
			if i, ic := CheckContConditions(v2, v, kws); ic {
				a.Incorrect = append(a.Incorrect, i)
			}
		}

		if len(a.Duplicate) > 0 || len (a.Incorrect) > 0 {
			add = append(add, a)
		}
	}

	return add, nil
}

func CheckContConditions(s string, c config.CheckItem, kws []string) (*IncorrectCont, bool){
	add := &IncorrectCont{}
	var ic bool

	if c.Length.Check != "" {
		add.Length = Length(s, c.Length)
		if add.Length != nil {
			ic = true
		}
	}
	if c.Length.Check != "" {
		add.UseKws = UseKws(s, c.UseKws, kws)
		if add.UseKws != nil {
			ic = true
		}
	}
	if len(c.NotPermit) > 0 {
		add.NotPermit = NotPermit(s, c.NotPermit)
		if add.Length != nil {
			ic = true
		}
	}
	if len(c.NotInclude) > 0 {
		add.NotInclude = NotInclude(s, c.NotInclude)
		if add.NotInclude != nil {
			ic = true
		}
	}


	if c.MultipleNum.Check != "" {
		add.MultipleNum = MultipleNum(s, c.MultipleNum)
		if add.MultipleNum != nil {
			ic = true
		}
	}
	if c.MultipleDup.Check != "" {
		add.MultipleDup = MultipleDup(s, c.MultipleDup)
		if add.MultipleNum != nil {
			ic = true
		}
	}

	return add, ic
}
