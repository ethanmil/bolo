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
