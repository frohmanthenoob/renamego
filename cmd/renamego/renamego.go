package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filepath.Walk(currentPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return err
		}

		if currentPath == path {
			return err
		}

		if info.IsDir() {
			return filepath.SkipDir
		}

		pathDir := filepath.Dir(path)
		fileExt := filepath.Ext(path)
		fileName := filepath.Base(path)
		fileName = strings.TrimSuffix(fileName, fileExt)

		r := regexp.MustCompile(`[\s,，.︰:–!-]+`)
		newFileName := r.ReplaceAllString(fileName, "_")

		for strings.Index(newFileName, "__") >= 0 {
			n := strings.Index(newFileName, "__")
			newFileName = newFileName[:n] + "_" + newFileName[n+2:]
		}

		fileName = fileName + fileExt
		newFileName = newFileName + fileExt
		newPath := pathDir + "/" + newFileName

		if path != newPath {
			fmt.Println(fileName)
			fmt.Println(newFileName)
			os.Rename(path, newPath)
		}

		return err
	})

}

// GOBIN=$GOPATH/bin go install cmd/renamego.go
