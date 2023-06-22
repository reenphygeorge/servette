package path

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/reenphygeorge/light-server/internal/logger"
)

/*
	Filter path from directories like .git and node_modules.
	mode = 0: Path to folders are provided for watching.
	mode = 1: HTML files are provided for logging.
*/
func filterPaths(path string, skipDirs []string, file fs.DirEntry, err error, mode int) (error, []string) {
	var dirList []string

	if err != nil {
		logger.Error()
		return err, nil
	}

	for _, skipDir := range skipDirs {
		if strings.Contains(path, skipDir) {
			return filepath.SkipDir, nil
		}
	}

	if mode == 0 && file.IsDir() {
		dirList = append(dirList, path)
	} else if mode == 1 && !file.IsDir() && strings.HasSuffix(path,".html") {
		dirList = append(dirList, path)
	}

	return nil, dirList
}

/*
	Walks recursively through all folders in the provided rootPath
	calls the filter function and returns the paths to be traversed.
	Mode is specified and used in filterPaths()
*/
func GetFilePaths(rootPath string, skipDirs []string, mode int) []string {

	var visitedPaths []string
	if mode == 0 {
		visitedPaths = append(visitedPaths, rootPath)
	}

	err := filepath.WalkDir(rootPath, func(path string, file os.DirEntry, err error) error {
		err, paths := filterPaths(path, skipDirs, file, err, mode)
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
