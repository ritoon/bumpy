package main

import (
	"errors"
	"os/exec"
	"strings"
)

func setGitTag(value string) error {
	cmd := exec.Command("git", "tag", value)
	return cmd.Run()
}

func deleteGitTag(value string) error {
	cmd := exec.Command("git", "tag", "-d", value)
	return cmd.Run()
}

func getGitTags() ([]string, error) {
	cmd := exec.Command("git", "tag")
	versions, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	if len(versions) == 0 {
		return nil, errors.New("bumpy: ")
	}
	res := strings.Split(string(versions), string(rune(10)))
	if len(res[len(res)-1]) == 0 {
		return res[:len(res)-1], nil
	}
	return res, nil
}
