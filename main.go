package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		files, err := getFilesRec(arg)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
}

func getFilesRec(dirName string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			childrenFiles, err := getFilesRec(dirName + string(os.PathSeparator) + file.Name())
			if err != nil {
				return nil, err
			}
			files = append(files, childrenFiles...)
		}
	}
	return files, nil
}
