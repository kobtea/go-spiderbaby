package stringplay

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	size = 20
)

func SpiderBaby() (string, error) {
	url := fmt.Sprintf("https://search.yahoo.co.jp/image/search?p=ストリングプレイスパイダーベイビー&n=%d", size)
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	var links []string
	doc.Find("#gridlist").Find("a > img").Each(func(i int, s *goquery.Selection) {
		if l, ok := s.Attr("src"); ok {
			links = append(links, l)
		}
	})

	return links[rand.Intn(size)], nil
}
