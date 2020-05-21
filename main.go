package main

import (
	"fmt"
	"log"
	"os"

	"github.com/avast/retry-go"
	"github.com/salaleser/scraper"
)

// ru games 134440

type row struct {
	ID    string
	Value string
}

var (
	err      error
	c        chan row
	filename string
	location string
	language string
)

func main() {
	genreIDs := []int{
		36, 6014,
	}

	for _, genreID := range genreIDs {
		filename = fmt.Sprintf("result/%d.csv", genreID)
		c = make(chan row, 20)

		locations := []string{
			"us",
			"br",
			"no",
			"vn",
			"hu",
			"tr",
			"pt",
			"id",
			"fi",
			"pl",
			"ar",
			"ie",
			"at",
			"se",
			"ua",
			"ru",
			"es",
			"au",
			"gb",
			"nl",
			"ch",
			"be",
			"ca",
			"co",
			"cz",
			"dk",
			"it",
			"nz",
			"za",
			"mx",
			"fr",
			"de",
		}

		log.Printf("Start %d.", genreID)
		for _, location := range locations {
			go scrape(genreID, location)
		}

		rows := make([]row, len(locations))
		for i := 0; i < len(rows); i++ {
			rows[i] = <-c
			fmt.Printf("%d. %s: %s\n", i+1, rows[i].ID, rows[i].Value)
		}

		save(rows)

		log.Printf("Done with %d rows.", len(rows))
	}
}

func scrape(genreID int, location string) {
	var err error
	var page scraper.Page
	err = retry.Do(
		func() error {
			page, err = scraper.AsGenre(genreID, location)
			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(5),
	)
	if err != nil {
		log.Println(genreID, location, err)
	}

	c <- row{location, page.PageData.MetricsBase.PageID}
}

func save(rows []row) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("open file %q: %v", filename, err)
	}
	defer f.Close()

	for _, row := range rows {
		_, err = f.WriteString(fmt.Sprintf("%q,%q\n", row.ID, row.Value))
		if err != nil {
			log.Fatalf("write row %v: %v", row, err)
		}
	}
}
