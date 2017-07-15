package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*
 * This file contains all of the necessary uitility functions that the package will use and share amongst each other
 */

func findParentDir(toFind, curDir string) (dir string, found bool) {
	log.Printf("Current dir: %v\n", curDir)
	if curDir == "/" {
		return "", false
	}

	if toFind == filepath.Base(curDir) {
		return toFind, true
	}

	dir, found = findParentDir(toFind, filepath.Dir(curDir))
	return
}

func writeDirectory(toWrite, curDir string) error {
	log.Printf("Current dir: %v\n", curDir)

	files, err := ioutil.ReadDir(curDir)

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if found := switchFile(file); found {
			return nil
		}
	}
	return fmt.Errorf("Failed to write file %v", toWrite)
}

func switchFile(file os.FileInfo) bool {
	if file.IsDir() {

	}
	return false
}
