package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/reenphygeorge/servette/internal/config-handle"
	watch "github.com/reenphygeorge/servette/internal/file-watcher"
	"github.com/reenphygeorge/servette/internal/logger"
	"github.com/reenphygeorge/servette/internal/path"
	"github.com/reenphygeorge/servette/internal/server"
)

func main() {
	args := os.Args

	if len(args) == 1 { 
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-interrupt
			logger.Exit()
			time.Sleep(500 * time.Millisecond)
			os.Exit(0)
		}()

		rootPath := "."
		var configObject config.Config
		
		config.GetValues(&configObject)
		logger.StaringServer()
		vistedPath := path.GetFilePaths(rootPath, configObject.SkipDirectories,0)
		htmlFiles := path.GetFilePaths(rootPath, configObject.SkipDirectories,1)
		go watch.WatchFiles(vistedPath, &htmlFiles, configObject, rootPath)
		server.Server(configObject.Port, &htmlFiles)	
	} else {
		for i := 1; i < len(args); i++ {
			arg := args[i]
			if arg == "-v" || arg == "--version" {
				fmt.Println("Servette version 1.0.0")
			} else {
				fmt.Println("unknown option: ", arg)
			}
		}
	}
}
