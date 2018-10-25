package main

import (
	"fmt"
	"github.com/kobtea/go-spiderbaby/stringplay"
	"github.com/pkg/browser"
)

func main() {
	url, err := stringplay.SpiderBaby()
	if err != nil {
		fmt.Println(err)
	}
	if err = browser.OpenURL(url); err != nil {
		fmt.Println(err)
	}
}
