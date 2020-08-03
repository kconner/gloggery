package main

import (
	"log"
	"path/filepath"
	"regexp"
	"time"
)

type Post struct {
	Filename string
	Date string
	Body string
}

func readPost(folder string, filename string) *Post {
	path := filepath.Join(folder, filename)

	date := dateStringFromFilename(filename)
	body := readFile(path)

	return &Post{
		Filename: filename,
		Date: date,
		Body: string(body),
	}
}

var filenameDateRegex = regexp.MustCompile("^\\d{4}-\\d{2}-\\d{2}-")

func dateStringFromFilename(filename string) string {
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
