package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	targetUrl := "https://www.trackingdifferences.com/ETF/ISIN/IE00B3RBWM25"

	c := colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(targetUrl)
}
