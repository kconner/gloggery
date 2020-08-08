package main

import (
	"fmt"
	"log"
	"regexp"
	"time"
)

type post struct {
	ModifiedTime time.Time
	Filename string
	URL      string
	Date     string
	ISODate  string
	ReadBody func() string
}

func newPost(folder string, item folderItem, indexURL string) *post {
	geminiFilename := fmt.Sprintf("%v.gmi", item.Filename)

	url := fmt.Sprintf("%v/%v", indexURL, geminiFilename)

	date, isoDate := parseFilenameDate(item.Filename)

	readBody := func() string {
		return string(readFile(folder, item.Filename))
	}

	return &post{
		ModifiedTime: item.ModifiedTime,
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

func (p *post) ShouldBuild(geminiFolder string) bool {
	geminiModifiedTime, ok := findModifiedTime(geminiFolder, p.Filename)
	if !ok {
		return true
	}

	return geminiModifiedTime.Before(p.ModifiedTime)
}

func (p *post) Body() string {
	return p.ReadBody()
}
