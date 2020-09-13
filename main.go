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
		oldPath := dirPath + "/" + fileInfo.Name()
		if oldPath == bufferPath {
			continue // デスクトップに置いているbufferディレクトリは残す
		}

		newPath, err := newPath(fileInfo)
		if err != nil {
			return err
		}

		err = os.Rename(oldPath, *newPath)
		if err != nil {
			return err
		}
	}
	fmt.Println(dirPath + " is cleanup!")

	return nil
}

func newPath(info os.FileInfo) (*string, error) {
	var newPath string
	existFile, err := os.Stat(bufferPath + "/" + info.Name())
	if existFile != nil {
		newPath = bufferPath + "/" + info.Name() + "_1"
	}
	if err != nil {
		if os.IsNotExist(err) {
			newPath = bufferPath + "/" + info.Name()
		} else {
			return nil, err
		}
	}

	return &newPath, nil
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
