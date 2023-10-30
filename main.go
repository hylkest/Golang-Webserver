package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %v\nFailed with response: %v\nError: %v\n", r.Request.URL, r, err)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
	})

	err := c.Visit("https://cssight.com")
	if err != nil {
		log.Fatal(err)
	}
}
