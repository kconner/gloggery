package main

import (
	"fmt"
	"text/template"
)

func writeGlog(templates *template.Template, folder string, posts []*Post) {
	fmt.Printf("in %v:\n", folder)

	// While writing post pages, collect filenames to link from the index
	filenames := make([]string, 0, len(posts))
	for _, post := range posts {
		filename := fmt.Sprintf("%v.gmi", post.Filename)
		filenames = append(filenames, filename)

		renderTemplate(templates, folder, filename, postTemplateName, *post)
	}

	renderTemplate(templates, folder, "index.gmi", indexTemplateName, filenames)
}
