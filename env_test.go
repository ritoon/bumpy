package main

import "testing"

func TestSetEnvFilesVersion(t *testing.T) {
	setEnvFilesVersion("./.env.prod", "v1.2.3")
}
