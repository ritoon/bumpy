package main

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := getEnvFiles()
	if err != nil {
		panic(err)
	}
	for k := range files {
		log.Println(files[k].Name())
	}
}

func getEnvFiles() ([]fs.DirEntry, error) {
	files, err := os.ReadDir("./")
	if err != nil {
		return nil, err
	}

	res := make([]fs.DirEntry, 0)

	for k := range files {
		if strings.Contains(files[k].Name(), ".env") {
			res = append(res, files[k])
		}
	}
	return res, nil
}
