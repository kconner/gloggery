package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// TODO: Decide paths with command-line options
	homeFolder := os.Getenv("HOME")
	postsFolder := filepath.Join(homeFolder, "gloggery/posts")
	templatesFolder := filepath.Join(homeFolder, "gloggery/templates")
	glogFolder := filepath.Join(homeFolder, "public_gemini/glog")

	postFilenames := listFilenamesOrderedReverse(postsFolder)

	if len(postFilenames) == 0 {
		fmt.Printf("no posts in %v\n", postsFolder)
		os.Exit(0)
	}

	templates := parseTemplates(templatesFolder, "post.tmpl", "index.tmpl")

	posts := make([]*Post, 0, len(postFilenames))
	for _, postFilename := range postFilenames {
		posts = append(posts, readPost(postsFolder, postFilename))
	}

	writeGlog(templates, glogFolder, posts)
}
