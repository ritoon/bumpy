package main

import (
	"log"
	"os/exec"
)

func main() {
	files, err := getEnvFiles()
	if err != nil {
		panic(err)
	}

	for k := range files {
		log.Println(files[k].Name())
	}

	err = execShell("test.sh", "hello")
	if err != nil {
		log.Println(err)
	}
}

func execShell(filename, command string) error {
	cmd := exec.Command("./"+filename, command)
	res, err := cmd.Output()
	if err != nil {
		return err
	}
	log.Println(string(res))
	return nil
}
