package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func listFilenamesOrderedReverse(folder string) []string {
	folderItems, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatalf("can't list folder %v", folder)
	}

	result := make([]string, 0, len(folderItems))

	// ioutil.ReadDir returns results sorted by name,
	// which is chronological order for date-prefixed files.
	// We want reverse chronological order.
	for i := len(folderItems) - 1; 0 <= i; i-- {
		result = append(result, folderItems[i].Name())
	}

	return result
}

func readFile(path string) []byte {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("can't read %v", path)
	}

	return result
}

func writeToFile(folder, filename string, write func(writer io.Writer)) {
	path := filepath.Join(folder, filename)

	mode := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(path, mode, 0644)
	if err != nil {
		log.Fatalf("can't write file %v", file)
	}
	defer file.Close()

	write(file)

	fmt.Printf("wrote %v\n", filename)
}
