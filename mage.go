// +build mage

package main

import (
	magesh "github.com/magefile/mage/sh"
)

func Proto() error {
	args := []string{
		"-I", ".",
		"./guide/bolo.proto",
		"--go_opt=paths=source_relative",
		"--go_out=plugins=grpc:.",
	}

	return magesh.RunV("protoc", args...)
}

func C() error {
	args := []string{
		"run",
		"client/client.go",
	}

	return magesh.RunV("go", args...)
}

func S() error {
	args := []string{
		"run",
		"server/server.go",
	}

	return magesh.RunV("go", args...)
}
