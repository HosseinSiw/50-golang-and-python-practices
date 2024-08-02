package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func WriteToCsv(fileName string, scrapedDate ScrapedData) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)
	writer := csv.NewWriter(file)
	headers := []string{"title", "url", "points"}
	err = writer.Write(headers)
	if err != nil {
		return err
	}
	for i := 0; i < scrapedDate.num; i++ {
		title := scrapedDate.general["title"][i]
		url := scrapedDate.general["url"][i]
		recored := []string{title, url, scrapedDate.points}
		err = writer.Write(recored)
		if err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Printf("%d recoreds has been written into %s", scrapedDate.num, fileName)
	return nil
}
