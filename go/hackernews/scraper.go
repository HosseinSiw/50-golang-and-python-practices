package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type ScrapedData struct {
	// Creating a struct which contains all of our information
	general map[string][]string
	num     int
	points  string
}

func ScrapeHackerNews(url string) ScrapedData {
	// Creating a map that contains the information about a single news
	var result = map[string][]string{
		"title": {},
		"url":   {},
	}
	scrapedData := ScrapedData{
		general: result,
		num:     0,
	}
	// Initializing new collector
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"),
	)
	// This method will call exactly after the response backs
	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
	// This method will call if the response contains an HTML File
	collector.OnHTML("table#hnmain tr td table tbody tr td.title span.titleline a",
		func(element *colly.HTMLElement) {
			result["url"] = append(result["url"], element.Attr("href"))
			result["title"] = append(result["title"], element.Text)
			scrapedData.num += 1
		})

	collector.OnHTML("table#hnmain tr td table tbody tr td.subtext span.subline span.score",
		func(e *colly.HTMLElement) {
			scrapedData.points = e.Text
		})
	err := collector.Visit(url)
	if err != nil {
		panic(err)
	}
	return scrapedData
}
