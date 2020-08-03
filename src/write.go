package main

import (
	"fmt"
	"io"
	"text/template"
)

func writeGlog(templates *template.Template, folder string, posts []*Post) {
	fmt.Printf("in %v:\n", folder)

	// While writing post pages, collect filenames to link from the index
	filenames := make([]string, 0, len(posts))
	for _, post := range posts {
		filename := fmt.Sprintf("%v.gmi", post.Filename)
		filenames = append(filenames, filename)

		writePostPage(templates, folder, filename, post)
	}

	writeIndexPage(templates, folder, "index.gmi", filenames)
}

func writePostPage(templates *template.Template, folder, filename string, post *Post) {
	writeToFile(folder, filename, func(writer io.Writer) {
		templates.ExecuteTemplate(writer, "post.tmpl", *post)
	})
}

func writeIndexPage(templates *template.Template, folder, filename string, postFilenames []string) {
	writeToFile(folder, filename, func(writer io.Writer) {
		templates.ExecuteTemplate(writer, "index.tmpl", postFilenames)
	})
}
