package pkg

import "github.com/PuerkitoBio/goquery"


func Scraping(doc *goquery.Document, tag string, attr string, idt map[string]string) []string {
	list := []string{}

	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		var c string
		var ok bool

		if identify(s, idt) {
			switch attr {
			case "":
				c = s.Text()
			default:
				c, ok = s.Attr(attr)
				if !ok {
					c = ""
				}
			}

			if c != "" && !Duplicate(list, c) {
				list = append(list, c)
			}
		}
	})

	return list
}

func identify(s *goquery.Selection, idt map[string]string) bool {
	if idt == nil {
		return true
	}

	for k, v := range idt {
		c, ok := s.Attr(k)
		if !ok {
			return false
		}

		if c == v {
			return true
		}
	}

	return false
}
