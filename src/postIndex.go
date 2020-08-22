package main

import (
	"fmt"
	"os"
	"time"
)

type postIndex struct {
	URL           string
	Title         string
	SiteTitle     string
	FeedTitle     string
	Posts         []*post
	GeneratedTime time.Time
}

func loadPostIndex(folder, url, title, siteTitle, feedTitle string, result chan *postIndex) {
	generatedTime := time.Now().In(time.UTC)

	items := listFolderItemsReverse(folder)

	if len(items) == 0 {
		fmt.Printf("no posts in %v\n", folder)
		os.Exit(0)
	}

	taskDone := make(chan int)
	taskCount := 0

	index := postIndex{
		URL:           url,
		Title:         title,
		SiteTitle:     siteTitle,
		FeedTitle:     feedTitle,
		Posts:         make([]*post, len(items), len(items)),
		GeneratedTime: generatedTime,
	}

	for i, item := range items {
		i := i
		item := item
		go func() {
			index.Posts[i] = newPost(&index, folder, item, url)
			taskDone <- 1
		}()
		taskCount++
	}

	for i := 0; i < taskCount; i++ {
		<-taskDone
	}

	for i := 1; i < len(index.Posts); i++ {
		index.Posts[i].NextPost = index.Posts[i-1]
		index.Posts[i-1].PreviousPost = index.Posts[i]
	}

	result <- &index
}

func (pi *postIndex) GeneratedISOTime() string {
	return pi.GeneratedTime.Format(time.RFC3339)
}

func (pi *postIndex) LatestPosts(limit int) []*post {
	if len(pi.Posts) <= limit {
		return pi.Posts
	}

	return pi.Posts[:limit]
}
