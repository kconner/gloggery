package main

import (
	"log"
	"path/filepath"
	"text/template"
)

func parseTemplates(folder string, filenames ...string) *template.Template {
	filePaths := make([]string, 0, len(filenames))
	for _, filename := range filenames {
		filePaths = append(filePaths, filepath.Join(folder, filename))
	}

	result, err := template.ParseFiles(filePaths...)
	if err != nil {
		log.Fatalf("can't parse templates %v", filePaths)
	}

	return result
}
