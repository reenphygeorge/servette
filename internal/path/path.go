package path

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/reenphygeorge/light-server/internal/logger"
)

/*
Filter path from directories and
skips checking directories like .git and node_modules.
*/
func filterPaths(path string, skipDirs []string, file fs.DirEntry, err error) (error, []string) {
	var pathList []string

	if err != nil {
		logger.Error()
		return err, nil
	}

	for _, dir := range skipDirs {
		if strings.Contains(path, dir) {
			return filepath.SkipDir, nil
		}
	}

	if !file.IsDir() {
		pathList = append(pathList, path)
	}

	return nil, pathList
}

/*
Walks recursively through all folders in the provided rootPath
calls the filter function and returns the paths to be traversed.
*/
func GetFilePaths(rootPath string, skipDirs []string) []string {

	var visitedPaths []string

	err := filepath.WalkDir(rootPath, func(path string, file os.DirEntry, err error) error {
		err, paths := filterPaths(path, skipDirs, file, err)
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
