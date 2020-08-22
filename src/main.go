package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	inputFolder, outputFolder, url, title, siteTitle, feedTitle, shouldRebuild := parseFlags()

	postsFolder := filepath.Join(inputFolder, "posts")
	templatesFolder := filepath.Join(inputFolder, "templates")

	postIndex := make(chan *postIndex)
	go loadPostIndex(postsFolder, url, title, siteTitle, feedTitle, postIndex)

	builder := make(chan *builder)
	go loadBuilder(templatesFolder, builder)

	(<-builder).buildGlog(outputFolder, <-postIndex, shouldRebuild)
}

func parseFlags() (inputFolder, outputFolder, url, title, siteTitle, feedTitle string, shouldRebuild bool) {
	host, err := os.Hostname()
	if err != nil {
		host = "host"
	}

	user := os.Getenv("USER")
	if user == "" {
		user = "user"
	}

	homeFolder, err := os.UserHomeDir()
	if err != nil {
		homeFolder = filepath.Join("/home", user)
	}

	flag.StringVar(&inputFolder, "input", filepath.Join(homeFolder, ".gloggery"), "folder path containing posts and templates subfolders")
	flag.StringVar(&outputFolder, "output", filepath.Join(homeFolder, "public_gemini/glog"), "folder path to receive Gemini files")
	flag.StringVar(&url, "url", fmt.Sprintf("gemini://%v/~%v/glog", host, user), "gemini:// URL equivalent to the output folder")
	flag.StringVar(&title, "title", "Glog", "reader-facing glog title")
	flag.StringVar(&siteTitle, "site-title", fmt.Sprintf("~%v", user), "reader-facing title of the glog's parent page")
	flag.StringVar(&feedTitle, "feed-title", "", "Atom feed title, if different from -site-title")
	flag.BoolVar(&shouldRebuild, "rebuild", false, "rebuild posts even if unchanged")

	flag.Parse()

	if feedTitle == "" {
		feedTitle = siteTitle
	}

	return
}
