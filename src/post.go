package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

type post struct {
	Index        *postIndex
	NextPost     *post
	PreviousPost *post
	ModifiedTime time.Time
	Filename     string
	URL          string
	PostTime     time.Time
	Title        string
	Body         string
}

func newPost(index *postIndex, folder string, item folderItem, indexURL string) *post {
	postTime, slug := parseFilename(item.Filename)
	postDateString := postTime.Format("2006-01-02")
	geminiFilename := fmt.Sprintf("%v-%v.gmi", postDateString, slug)

	title := "Untitled"
	body := string(readFile(folder, item.Filename))

	splitBody := strings.SplitN(body, "\n\n", 2)
	if len(splitBody) == 2 {
		title = splitBody[0]
		body = splitBody[1]
	}

	return &post{
		Index:        index,
		ModifiedTime: item.ModifiedTime,
		Filename:     geminiFilename,
		URL:          fmt.Sprintf("%v/%v", indexURL, geminiFilename),
		PostTime:     postTime,
		Title:        title,
		Body:         body,
	}
}

var filenameRegex = regexp.MustCompile("^(\\d{4}-\\d{2}-\\d{2}-\\d{4})-(.*)")

func parseFilename(filename string) (postTime time.Time, slug string) {
	matches := filenameRegex.FindStringSubmatch(filename)
	if len(matches) == 0 {
		log.Fatalf("can't parse post filename %v", filename)
	}

	postTime, err := time.Parse("2006-01-02-1504", matches[1])
	if err != nil {
		log.Fatalf("can't parse post time from post filename %v", filename)
	}

	slug = matches[2]
	return
}

func (p *post) ShouldBuild(geminiFolder string) bool {
	return p.HasChanged(geminiFolder) ||
		(p.NextPost != nil && p.NextPost.HasChanged(geminiFolder)) ||
		(p.PreviousPost != nil && p.PreviousPost.HasChanged(geminiFolder))
}

func (p *post) HasChanged(geminiFolder string) bool {
	geminiModifiedTime, ok := findModifiedTime(geminiFolder, p.Filename)
	if !ok {
		return true
	}

	return geminiModifiedTime.Before(p.ModifiedTime)
}

func (p *post) ReadableDate() string {
	return p.PostTime.Format("2 January 2006")
}

func (p *post) ISODate() string {
	return p.PostTime.Format("2006-01-02")
}

func (p *post) ISOTime() string {
	return p.PostTime.Format(time.RFC3339)
}
