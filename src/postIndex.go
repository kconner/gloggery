package main

import (
	"fmt"
	"os"
	"time"
)

type postIndex struct {
	URL           string
	Title         string
	Posts         []*post
	GeneratedTime time.Time
}

func loadPostIndex(folder, url, title string, result chan *postIndex) {
	generatedTime := time.Now().In(time.UTC)

	items := listFolderItemsReverse(folder)

	if len(items) == 0 {
		fmt.Printf("no posts in %v\n", folder)
		os.Exit(0)
	}

	taskDone := make(chan int)
	taskCount := 0

	posts := make([]*post, len(items), len(items))
	for index, item := range items {
		index := index
		item := item
		go func() {
			posts[index] = newPost(folder, item, url)
			taskDone <- 1
		}()
		taskCount++
	}

	for i := 0; i < taskCount; i++ {
		<-taskDone
	}

	result <- &postIndex{
		Title:         title,
		URL:           url,
		Posts:         posts,
		GeneratedTime: generatedTime,
	}
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
