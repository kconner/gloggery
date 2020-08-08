package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type folderItem struct {
	Filename string
	ModifiedTime time.Time
}

func listFolderItemsReverse(folder string) []folderItem {
	fileInfos, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatalf("can't list folder %v", folder)
	}

	items := make([]folderItem, 0, len(fileInfos))

	// ioutil.ReadDir returns results sorted by name,
	// which is chronological order given date-prefixed files.
	// For a glog, we want reverse chronological order.
	for i := len(fileInfos) - 1; 0 <= i; i-- {
		items = append(items, folderItem{
			Filename: fileInfos[i].Name(),
			ModifiedTime: fileInfos[i].ModTime(),
		})
	}

	return items
}

func findModifiedTime(folder, filename string) (time.Time, bool) {
	path := filepath.Join(folder, filename)

	fileInfo, err := os.Stat(path)
	if err != nil {
		return time.Time{}, false
	}

	return fileInfo.ModTime(), true
}

func readFile(folder, filename string) []byte {
	path := filepath.Join(folder, filename)

	result, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("can't read %v", path)
	}

	return result
}

func writeFile(folder, filename string, write func(writer io.Writer)) {
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
