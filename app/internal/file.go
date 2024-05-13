package internal

import (
	"log"
	"os"
)

func LoadImages(filePath string) []string {
	files, err := os.ReadDir(filePath)
	if err != nil {
		log.SetPrefix(`loadimages func dead`)
		log.Fatal(err)
	}

	var fileSlice []string

	for _, fileName := range files {
		fileSlice = append(fileSlice, fileName.Name())
	}
	return fileSlice
}
