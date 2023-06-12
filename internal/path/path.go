package path

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/reenphygeorge/light-server/internal/logger"
)

func explorePaths(path string, file fs.DirEntry, err error) (error, []string) {
	var pathList []string
	
	if err != nil {
		logger.Error()
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
	root := "."

	var visitedPaths []string

	err := filepath.WalkDir(root, func(path string, file os.DirEntry, err error) error {
		err, paths := explorePaths(path, file, err)
		if err != nil {
		}
		visitedPaths = append(visitedPaths, paths...)
		return nil
	})

	if err != nil {
		logger.Error()
	}

	return visitedPaths
}
