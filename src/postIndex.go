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
	filenames := listFolderItemsReverse(folder)

	if len(filenames) == 0 {
		fmt.Printf("no posts in %v\n", folder)
		os.Exit(0)
	}

	posts := make([]*post, 0, len(filenames))
	for _, filename := range filenames {
		posts = append(posts, newPost(folder, filename, url))
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

	return pi.Posts[0].ISODate
}

func (pi *postIndex) LatestPosts(limit int) []*post {
	if len(pi.Posts) <= limit {
		return pi.Posts
	}

	return pi.Posts[:limit]
}
