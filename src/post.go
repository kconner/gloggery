package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

type post struct {
	Filename string
	URL      string
	Date     string
	ISODate  string
	ReadBody func() string
}

func newPost(folder, filename, indexURL string) *post {
	geminiFilename := fmt.Sprintf("%v.gmi", filename)

	url := fmt.Sprintf("%v/%v", indexURL, geminiFilename)

	date, isoDate := parseFilenameDate(filename)

	readBody := func() string {
		return string(readFile(folder, filename))
	}

	return &post{
		Filename: geminiFilename,
		URL:      url,
		Date:     date,
		ISODate:  isoDate,
		ReadBody: readBody,
	}
}

var filenameDateRegex = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}-")

func parseFilenameDate(filename string) (readableDate string, isoDate string) {
	match := filenameDateRegex.FindString(filename)
	if match == "" {
		log.Fatalf("can't parse date from post filename %v", filename)
	}

	date, err := time.Parse("2006-01-02-", match)
	if err != nil {
		log.Fatalf("can't parse date from post filename %v", filename)
	}

	readableDate = date.Format("2 January 2006")
	isoDate = date.Format(time.RFC3339)
	return
}

func (p *post) Body() string {
	return p.ReadBody()
}
