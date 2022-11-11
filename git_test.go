package main

import (
	"testing"
)

func TestGetGitTags(t *testing.T) {
	res, _ := getGitTags()
	t.Log(res)

	for _, v := range res {
		for _, w := range v {
			t.Log("loop")
			t.Log(w)
			t.Log(string(w))
		}
	}
}
