package path

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func explorePaths(path string, file fs.DirEntry, err error) (error, []string) {
	var pathList []string
	if err != nil {
		log.Fatal(err)
		return err,nil
	}
	
	if strings.Contains(path,"git") {
		return filepath.SkipDir,nil
	}
	if !file.IsDir() {
		pathList = append(pathList, path)
	}
	return nil,pathList
}

func GetFilePaths() []string {
	root := "../"

	var visitedPaths []string

	err := filepath.WalkDir(root, func(path string, file os.DirEntry, err error) error {
		err, paths := explorePaths(path, file, err)
		if err != nil {
			// log.Println(err)
		}
		visitedPaths = append(visitedPaths, paths...)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return visitedPaths
}
