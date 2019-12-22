package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var homePath = os.Getenv("HOME")
var desktopPath = homePath + "/Desktop"
var bufferPath = desktopPath + "/buffer"

func stashAllContentsOfDir(dirPath string) error {
	fmt.Println(dirPath + " is cleaning now...")
	fileInfoList, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, fileInfo := range fileInfoList {
		absPath := dirPath + "/" + fileInfo.Name()
		if absPath == bufferPath {
			continue
		}

		err := os.Rename(absPath, bufferPath+"/"+fileInfo.Name())
		if err != nil {
			return err
		}
	}
	fmt.Println(dirPath + " is cleanup!")

	return nil
}

func main() {
	_, err := os.Stat(bufferPath)
	if err != nil {
		if err := os.Mkdir(bufferPath, 0777); err != nil {
			fmt.Println(err)
			return
		}
	}

	if err := stashAllContentsOfDir(desktopPath); err != nil {
		fmt.Println(err)
		return
	}

	if err := stashAllContentsOfDir(homePath + "/Downloads"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("cleanup is completed!!!")
}
