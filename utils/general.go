package utils

// Find the path name for the current directory.

import (
	"log"
	"os"
)

func CurrentDirectoryPathName() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}
