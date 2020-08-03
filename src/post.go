package main

import (
	"log"
	"path/filepath"
	"regexp"
	"time"
)

type post struct {
	Filename string
	Date string
	Body string
}

func newPost(folder string, filename string) *post {
	path := filepath.Join(folder, filename)

	date := parseFilenameDate(filename)
	body := readFile(path)

	return &post{
		Filename: filename,
		Date: date,
		Body: string(body),
	}
}

var filenameDateRegex = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}-")

func parseFilenameDate(filename string) string {
	match := filenameDateRegex.FindString(filename)
	if match == "" {
		log.Fatalf("can't parse date from post filename %v", filename)
	}

	date, err := time.Parse("2006-01-02-", match)
	if err != nil {
		log.Fatalf("can't parse date from post filename %v", filename)
	}

	return date.Format("2 January 2006")
}
