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
const ()

func getParentDir(toFind, curDir string) (dir string, found bool) {
	log.Printf("Current dir from get parent dir: %v\n", curDir)
	if curDir == "/" {
		return "", false
	}

	if toFind == filepath.Base(curDir) {
		return curDir, true
	}

	err := os.Chdir(curDir)
	handleError(err, "Error changing directory from get parent dir")

	_, err = os.Stat(".rgConf")
	handleError(err, "from getparentdir os.Stat")
	if err == nil {
		return curDir, true
	}
	if !os.IsNotExist(err) {
		return "Something bad happend", false
	}

	dir, found = getParentDir(toFind, filepath.Dir(curDir))
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

func handleError(e error, tag string) {
	if e != nil {
		fmt.Printf("From Handle Error\n%v: %v\n", tag, e)
	}
}
