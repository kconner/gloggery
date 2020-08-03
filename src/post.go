package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

type post struct {
	Filename string
	Read     func() *postContent
}

type postContent struct {
	Date string
	Body string
}

func loadPosts(folder string, result chan []*post) {
	filenames := listFolderItemsReverse(folder)

	if len(filenames) == 0 {
		fmt.Printf("no posts in %v\n", folder)
		os.Exit(0)
	}

	posts := make([]*post, 0, len(filenames))
	for _, filename := range filenames {
		posts = append(posts, newPost(folder, filename))
	}

	result <- posts
}

func newPost(folder, filename string) *post {
	read := func() *postContent {
		date := parseFilenameDate(filename)
		body := readFile(folder, filename)

		return &postContent{
			Date: date,
			Body: string(body),
		}
	}

	return &post{
		Filename: filename,
		Read:     read,
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
