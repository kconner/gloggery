package main

import (
	"fmt"
	"io"
	"log"
	"path/filepath"
	"text/template"
)

const (
	postTemplateName = "post.tmpl"
	indexTemplateName = "index.tmpl"
)

var templateNames = []string{
	postTemplateName,
	indexTemplateName,
}

type builder struct {
	*template.Template
}

func newBuilder(folder string) *builder {
	paths := make([]string, 0, len(templateNames))
	for _, filename := range templateNames {
		paths = append(paths, filepath.Join(folder, filename))
	}

	template, err := template.ParseFiles(paths...)
	if err != nil {
		log.Fatalf("can't parse templates %v", paths)
	}

	return &builder{template}
}

func (b *builder) buildGlog(folder string, posts []*post) {
	fmt.Printf("in %v:\n", folder)

	// While writing post pages, collect filenames to link from the index
	filenames := make([]string, 0, len(posts))
	for _, post := range posts {
		filename := fmt.Sprintf("%v.gmi", post.Filename)
		filenames = append(filenames, filename)

		b.buildFile(folder, filename, postTemplateName, *post)
	}

	b.buildFile(folder, "index.gmi", indexTemplateName, filenames)
}

func (b *builder) buildFile(folder, filename, templateName string, model interface{}) {
	writeFile(folder, filename, func(writer io.Writer) {
		b.ExecuteTemplate(writer, templateName, model)
	})
}
