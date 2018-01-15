package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// If we get error just after the request, this callback will be called
	// and colly will return at this point, other callbacks will not be called
	c.OnError(func (resp *colly.Response, e error) {
		fmt.Println(e.Error())
	})

	// After a success response, this callback will be called
	c.OnResponse(func (resp *colly.Response) {
		fmt.Printf("response with status code: %d\n", resp.StatusCode)
	})

	// this is the final callback, we usually doing some teardown jobs
	// e.g. write to file, report to api...
	c.OnScraped(func (resp *colly.Response) {
		fmt.Printf("crawling url: %s done\n", resp.Request.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
