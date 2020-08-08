package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	inputFolder, outputFolder, url, title, shouldRebuild := parseFlags()

	postsFolder := filepath.Join(inputFolder, "posts")
	templatesFolder := filepath.Join(inputFolder, "templates")

	postIndex := make(chan *postIndex)
	go loadPostIndex(postsFolder, url, title, postIndex)

	builder := make(chan *builder)
	go loadBuilder(templatesFolder, builder)

	(<-builder).buildGlog(outputFolder, <-postIndex, shouldRebuild)
}

func parseFlags() (inputFolder, outputFolder, url, title string, shouldRebuild bool) {
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
	flag.StringVar(&url, "url", fmt.Sprintf("gemini://%v/~%v/glog", host, user), "gemini:// url equivalent to the output folder")
	flag.StringVar(&title, "title", fmt.Sprintf("~%v", user), "reader-facing glog title")
	flag.BoolVar(&shouldRebuild, "rebuild", false, "rebuild posts even if unchanged")

	flag.Parse()
	return
}
