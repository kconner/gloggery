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
	atomTemplateName = "atom.tmpl"
)

var templateNames = []string{
	postTemplateName,
	indexTemplateName,
	atomTemplateName,
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

func (b *builder) buildGlog(folder string, postIndex *postIndex) {
	fmt.Printf("in %v:\n", folder)

	// Start each build task and collect their completion signals
	taskDone := make(chan int)

	for _, post := range postIndex.Posts {
		go b.buildFile(folder, post.Filename, postTemplateName, post, taskDone)
	}

	go b.buildFile(folder, "index.gmi", indexTemplateName, postIndex, taskDone)

	go b.buildFile(folder, "atom.xml", atomTemplateName, postIndex, taskDone)

	taskCount := 2 + len(postIndex.Posts)
	for i := 0; i < taskCount; i++ {
		<-taskDone
	}
}

func (b *builder) buildFile(folder, filename, templateName string, model interface{}, taskDone chan int) {
	writeFile(folder, filename, func(writer io.Writer) {
		b.ExecuteTemplate(writer, templateName, model)
	})

	taskDone <- 1
}
