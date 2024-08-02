package main

/*
In this example we will learn how to scrape data from a web page.
*/

import (
	"fmt"
	"log"
)

func main() {
	result := ScrapeHackerNews("https://news.ycombinator.com/")
	// showing the result in the console, Note that instead you can write the results into a file, eg a csv file.
	for i := 0; i < result.num; i++ {
		fmt.Println("Generals")
		msg := fmt.Sprintf("[Title]: %s\n[URL]: %s", result.general["title"][i], result.general["url"][i])
		fmt.Println(msg)
		fmt.Println("Details")
		msg = fmt.Sprintf("[points]: %s\n", result.points)
		fmt.Println(msg)
	}
	err := WriteToCsv("hackernews.csv", result)
	if err != nil {
		log.Fatalln(err)
	}
}
