package main

import (
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

func parseTemplates(folder string) *template.Template {
	paths := make([]string, 0, len(templateNames))
	for _, filename := range templateNames {
		paths = append(paths, filepath.Join(folder, filename))
	}

	result, err := template.ParseFiles(paths...)
	if err != nil {
		log.Fatalf("can't parse templates %v", paths)
	}

	return result
}

func renderTemplate(templates *template.Template, folder, filename, templateName string, model interface{}) {
	writeFile(folder, filename, func(writer io.Writer) {
		templates.ExecuteTemplate(writer, templateName, model)
	})
}
