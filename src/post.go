package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"time"
)

type postIndex struct {
	URL string
	Title string
	Posts []*post
}

type post struct {
	Filename string
	URL string
	Date string
	ISODate string
	ReadBody func() string
}

func loadPostIndex(folder, title, glogURL string, result chan *postIndex) {
	filenames := listFolderItemsReverse(folder)

	if len(filenames) == 0 {
		fmt.Printf("no posts in %v\n", folder)
		os.Exit(0)
	}

	posts := make([]*post, 0, len(filenames))
	for _, filename := range filenames {
		posts = append(posts, newPost(folder, filename, glogURL))
	}

	result <- &postIndex{
		Title: title,
		URL: glogURL,
		Posts: posts,
	}
}

func (pi *postIndex) LatestPostISODate() string {
	if (len(pi.Posts) == 0) {
		return ""
	}

	return pi.Posts[0].ISODate
}

func newPost(folder, filename, glogURL string) *post {
	geminiFilename := fmt.Sprintf("%v.gmi", filename)

	url := path.Join(glogURL, geminiFilename)

	date, isoDate := parseFilenameDate(filename)

	readBody := func() string {
		return string(readFile(folder, filename))
	}

	return &post{
		Filename: geminiFilename,
		URL: url,
		Date:     date,
		ISODate:     isoDate,
		ReadBody: readBody,
	}
}

func (p *post) Body() string {
	return p.ReadBody()
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
