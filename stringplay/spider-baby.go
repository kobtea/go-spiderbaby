package stringplay

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func SpiderBaby(size int) ([]string, error) {
	url := fmt.Sprintf("https://search.yahoo.co.jp/image/search?p=ストリングプレイスパイダーベイビー&n=%d", size)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	doc.Find("#gridlist").Find("a > img").Each(func(i int, s *goquery.Selection) {
		if l, ok := s.Attr("src"); ok {
			links = append(links, l)
		}
	})

	return links, nil
}
