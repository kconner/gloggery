package main

import (
	"fmt"
	"os"
)

type postIndex struct {
	URL   string
	Title string
	Posts []*post
}

func loadPostIndex(folder, url, title string, result chan *postIndex) {
	items := listFolderItemsReverse(folder)

	if len(items) == 0 {
		fmt.Printf("no posts in %v\n", folder)
		os.Exit(0)
	}

	posts := make([]*post, 0, len(items))
	for _, item := range items {
		posts = append(posts, newPost(folder, item, url))
	}

	result <- &postIndex{
		Title: title,
		URL:   url,
		Posts: posts,
	}
}

func (pi *postIndex) LatestPostISODate() string {
	if len(pi.Posts) == 0 {
		return ""
	}

	return pi.Posts[0].ISODate()
}

func (pi *postIndex) LatestPosts(limit int) []*post {
	if len(pi.Posts) <= limit {
		return pi.Posts
	}

	return pi.Posts[:limit]
}
