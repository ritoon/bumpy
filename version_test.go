package main

import "testing"

type testDataVersion struct {
	title    string
	valueIn  string
	valueOut string
	err      error
}

func TestVersionFix(t *testing.T) {
	testData := []testDataVersion{
		{"A", "", "v0.0.0", nil},
		{"B", "v0.0.0", "v0.0.1", nil},
	}

	for _, d := range testData {
		res, _ := VersionFix(d.valueIn)
		if res != d.valueOut {
			t.Errorf("for test %v got %v but wait for %v", d.title, res, d.valueOut)
		}
	}
}

func TestVersionMinor(t *testing.T) {
	testData := []testDataVersion{
		{"A", "", "v0.0.0", nil},
		{"B", "v0.0.0", "v0.1.0", nil},
	}

	for _, d := range testData {
		res, _ := VersionMinor(d.valueIn)
		if res != d.valueOut {
			t.Errorf("for test %v got %v but wait for %v", d.title, res, d.valueOut)
		}
	}
}

func TestVersionMajor(t *testing.T) {
	testData := []testDataVersion{
		{"A", "", "v0.0.0", nil},
		{"B", "v0.0.0", "v1.0.0", nil},
		{"C", "0.0.0", "v1.0.0", nil},
	}

	for _, d := range testData {
		res, _ := VersionMajor(d.valueIn)
		if res != d.valueOut {
			t.Errorf("for test %v got %v but wait for %v", d.title, res, d.valueOut)
		}
	}
}
