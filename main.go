package main

import (
	"fmt"
	"github.com/kobtea/go-spiderbaby/stringplay"
	"github.com/pkg/browser"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	size = 20
)

func main() {
	urls, err := stringplay.SpiderBaby(size)
	if err != nil {
		fmt.Println(err)
	}
	if err = browser.OpenURL(urls[rand.Intn(size)]); err != nil {
		fmt.Println(err)
	}
}
