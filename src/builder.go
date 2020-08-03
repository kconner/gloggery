package main

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"text/template"
)

type builder struct {
	*template.Template
}

const (
	postTemplateName  = "post.tmpl"
	indexTemplateName = "index.tmpl"
)

var templateNames = []string{
	postTemplateName,
	indexTemplateName,
}

func loadBuilder(folder string, result chan *builder) {
	paths := make([]string, 0, len(templateNames))
	for _, filename := range templateNames {
		paths = append(paths, filepath.Join(folder, filename))
	}

	template, err := template.ParseFiles(paths...)
	if err != nil {
		log.Fatalf("can't parse templates %v", paths)
	}

	result <- &builder{template}
}

func (b *builder) buildGlog(folder string, posts []*post) {
	fmt.Printf("in %v:\n", folder)

	// Start each build task and collect their completion signals
	taskDone := make(chan int)

	// While writing post pages, collect their filenames to link from the index
	filenames := make([]string, 0, len(posts))

	for _, post := range posts {
		filename := fmt.Sprintf("%v.gmi", post.Filename)
		filenames = append(filenames, filename)

		go b.buildPost(folder, filename, post, taskDone)
	}

	go b.buildFile(folder, "index.gmi", indexTemplateName, filenames, taskDone)

	taskCount := 1 + len(posts)
	for i := 0; i < taskCount; i++ {
		<-taskDone
	}
}

func (b *builder) buildPost(folder, filename string, post *post, taskDone chan int) {
	content := post.Read()

	b.buildFile(folder, filename, postTemplateName, content, taskDone)
}

func (b *builder) buildFile(folder, filename, templateName string, model interface{}, taskDone chan int) {
	writeFile(folder, filename, func(writer io.Writer) {
		b.ExecuteTemplate(writer, templateName, model)
	})

	taskDone <- 1
}
