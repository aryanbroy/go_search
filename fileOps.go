package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) string {
	log.Println("Reading file...")

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Error reading file: ", err)
	}

	return string(data)
}

func ListFiles(dir string) []string {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return files
}
