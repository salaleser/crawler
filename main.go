package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/avast/retry-go"
	"github.com/salaleser/scraper"
)

// ru games 134440

type row struct {
	ID    string
	Value string
}

const (
	location = "ru"
	language = "ru"
)

var (
	err      error
	c        chan row
	filename string
	start    int
	end      int
)

func main() {
	if len(os.Args) == 3 {
		start, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		end, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	filename = fmt.Sprintf("result/%s-%s-%d-%d.csv", location, language, start, end)
	c = make(chan row, 10)

	log.Printf("Start with %d rows.", end-start)
	for i := start; i < end; i++ {
		go scrape(strconv.Itoa(i), location, language)
	}

	rows := make([]row, end-start)
	for i := 0; i < len(rows); i++ {
		rows[i] = <-c
		if rows[i].Value == "" {
			log.Printf("#%d [MISS] %s", i+1, rows[i].ID)
		} else {
			log.Printf("#%d [HIT!] %s (%s)", i+1, rows[i].ID, rows[i].Value)
		}
	}

	save(rows)

	log.Printf("Done with %d rows.", len(rows))
}

func scrape(id string, location string, language string) {
	var err error
	var body []byte
	err = retry.Do(
		func() error {
			body, err = scraper.AsGrouping(id, location, language)
			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(25),
	)
	if err != nil {
		log.Printf("retry: %v", err)
	}

	value, err := parse(body)
	if err != nil {
		c <- row{id, ""}
	} else {
		c <- row{id, value}
	}
}

func parse(body []byte) (string, error) {
	var page scraper.Page
	err := json.Unmarshal(body, &page)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(page.PageData.GenreID), nil
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
