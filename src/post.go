package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

type post struct {
	ModifiedTime time.Time
	Filename     string
	URL          string
	Date         time.Time
	Title        string
	Body         string
}

func newPost(folder string, item folderItem, indexURL string) *post {
	geminiFilename := fmt.Sprintf("%v.gmi", item.Filename)

	title := "Untitled"
	body := string(readFile(folder, item.Filename))

	splitBody := strings.SplitN(body, "\n\n", 2)
	if len(splitBody) == 2 {
		title = splitBody[0]
		body = splitBody[1]
	}

	return &post{
		ModifiedTime: item.ModifiedTime,
		Filename:     geminiFilename,
		URL:          fmt.Sprintf("%v/%v", indexURL, geminiFilename),
		Date:         parseFilenameDate(item.Filename),
		Title:        title,
		Body:         body,
	}
}

var filenameDateRegex = regexp.MustCompile("^(\\d{4}-\\d{2}-\\d{2})-")

func parseFilenameDate(filename string) (date time.Time) {
	matches := filenameDateRegex.FindStringSubmatch(filename)
	if len(matches) == 0 {
		log.Fatalf("can't parse date from post filename %v", filename)
	}

	date, err := time.ParseInLocation("2006-01-02", matches[1], time.Local)
	if err != nil {
		log.Fatalf("can't parse date from post filename %v", filename)
	}
	return
}

func (p *post) ShouldBuild(geminiFolder string) bool {
	geminiModifiedTime, ok := findModifiedTime(geminiFolder, p.Filename)
	if !ok {
		return true
	}

	return geminiModifiedTime.Before(p.ModifiedTime)
}

func (p *post) ReadableDate() string {
	return p.Date.Format("2 January 2006")
}

func (p *post) ISODate() string {
	// 8 PM
	return p.Date.Add(time.Hour * 20).Format(time.RFC3339)
}
