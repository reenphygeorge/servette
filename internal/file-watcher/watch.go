package watch

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fsnotify/fsnotify"
	"github.com/reenphygeorge/servette/internal/config-handle"
	"github.com/reenphygeorge/servette/internal/logger"
	"github.com/reenphygeorge/servette/internal/path"
	"github.com/reenphygeorge/servette/internal/server"
)

/*
	Watch all files in the provided path for changes.
	If a new directory is created then it's path is added to watcher.
*/
func WatchFiles(pathList []string, htmlFiles *[]string, configObject config.Config) {
	for _,htmlFile := range(*htmlFiles) {
		fmt.Println("\033[33m\t",htmlFile,"\n\033[0m")
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error("")
	}
	defer watcher.Close()
	for _, path := range pathList {
		err = watcher.Add(path)
		if err != nil {
			logger.Error("")
		}
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			*htmlFiles = path.GetFilePaths(configObject.RootPath, configObject.SkipDirectories,1)
			if event.Op&fsnotify.Create == fsnotify.Create {
				isDir, _ := isDirectory(event.Name)
				if isDir == true {
					watcher.Add(event.Name)
				}
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				isDir, _ := isDirectory(event.Name)
				if isDir == true {
					watcher.Remove(event.Name)
				}
			}
			logger.Reloading()
			if server.GlobalConn != nil {
				server.ReloadRequest()
			} else {
				logger.StartAndReload(strconv.Itoa(configObject.Port),htmlFiles)
			}
		case _, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Error("")
		}
	}
}

// Checks whether the newly created item is a directory or a file.
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}
