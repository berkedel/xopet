package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func createDir(filePath string) {
	parentDir := filepath.Dir(filePath)
	_, err := os.Stat(parentDir)
	if os.IsNotExist(err) {
		e := os.MkdirAll(parentDir, os.ModePerm)
		if e != nil {
			fmt.Printf("> Error: Failed to create a directory: %s\n", parentDir)
		}
	}
}

func dumpFile(filePath string, content []byte) {
	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Printf("> Error: Could not dump the file: %s\n", filePath)
	}
}
