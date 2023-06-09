package watch

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func WatchFiles(pathList []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Add a directory or file to the watcher
	for _, path := range pathList {
		err = watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Start the event loop to receive file system notifications
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
			// fileState.SetFileState(false)
			return
			}
			fmt.Println("Event:", event.Name, event.Op)

			// Handle file saves (modifications)
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("File saved:", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}
}