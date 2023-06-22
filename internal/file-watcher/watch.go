package watch

import (
	"github.com/fsnotify/fsnotify"
	"github.com/reenphygeorge/light-server/internal/logger"
	"github.com/reenphygeorge/light-server/internal/server"
)

// Watch all files in the provided path for changes.
func WatchFiles(pathList []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error()
	}
	defer watcher.Close()
	for _, path := range pathList {
		err = watcher.Add(path)
		if err != nil {
			logger.Error()
		}
	}
	for {
		select {
		case _, ok := <-watcher.Events:
			if !ok {
				return
			}
			logger.Reloading()
			server.ReloadRequest()
		case _, ok := <-watcher.Errors:
			if !ok {
				return
			}
			logger.Error()
		}
	}
}
