package main

import (
	"bufio"
	"io/fs"
	"log"
	"os"
	"strings"
)

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

func setEnvFilesVersion(envFilePath, version string) error {
	envFile, err := os.Open(envFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)

	// var envVariableFound bool
	res, found := "", false
	for scanner.Scan() {
		log.Println(scanner.Text())
		if strings.Contains(scanner.Text(), "REACT_APP_VERSION") {
			elems := strings.Split(scanner.Text(), "=")
			res += elems[0] + "=" + version + "\n"
			found = true
		} else {
			res += scanner.Text() + "\n"
		}
	}
	if !found {
		res += "REACT_APP_VERSION=" + version
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	newFile, err := os.Create(envFilePath)
	if err != nil {
		return err
	}
	_, err = newFile.WriteString(res)
	if err != nil {
		return err
	}

	return nil
}
