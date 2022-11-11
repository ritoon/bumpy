package main

import (
	"errors"
	"strconv"
	"strings"
)

var currentVersion string

const (
	versionFix   = 2
	versionMinor = 1
	versionMajor = 0
)

func VersionFix(current string) (new string, err error) {
	return versionUp(current, versionFix)
}

func VersionMinor(current string) (new string, err error) {
	return versionUp(current, versionMinor)
}

func VersionMajor(current string) (new string, err error) {
	return versionUp(current, versionMajor)
}

func versionUp(current string, lvlVersion int) (new string, err error) {
	if len(current) == 0 {
		return "v0.0.0", nil
	}
	slitOfCurrent := strings.Split(current, ".")
	if len(slitOfCurrent) == 0 {
		return "", errors.New("bumpy: current version is not compatible.")
	}
	lastDigit := slitOfCurrent[lvlVersion]
	if lvlVersion == versionMajor {
		if lastDigit[0] == 'v' {
			lastDigit = lastDigit[1:]
		}
	}
	lastDigitInt, err := strconv.Atoi(lastDigit)
	if err != nil {
		return "", err
	}
	lastDigitInt++
	slitOfCurrent[lvlVersion] = strconv.Itoa(lastDigitInt)

	for k := range slitOfCurrent {
		new += slitOfCurrent[k] + "."
	}
	if lvlVersion == versionMajor {
		new = "v" + new
	}
	return new[:len(new)-1], nil
}
