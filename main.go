package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func main() {
	targetUrl := getUrl()
	targetElement := getElement()
	//demoUrl := "https://onepace.me/series/one-pace-english-sub/"

	c := colly.NewCollector()

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
