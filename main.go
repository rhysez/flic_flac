package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	targetUrl := getUrl()
	urlPrefix := targetUrl[0:4]
	if urlPrefix != "www." {
		panic("Please enter a valid URL")
	}

	targetElement := getElement()
	// demoUrl := "https://www.trackingdifferences.com/ETF/ISIN/IE00B3RBWM25"

	c := colly.NewCollector(colly.AllowedDomains("www.trackingdifferences.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(targetElement, func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit(targetUrl)
}

func getUrl() string {
	var targetUrl string
	fmt.Println("Enter a target URL to deploy flic flac:")
	fmt.Scanln(&targetUrl)

	return targetUrl
}

func getElement() string {
	var targetElement string
	fmt.Println("Enter a target HTML element for flic flac to extract from:")
	fmt.Scanln(&targetElement)
	return targetElement
}
