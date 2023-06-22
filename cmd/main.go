package main

import (
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
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interrupt
		logger.Exit()
		time.Sleep(500 * time.Millisecond)
		os.Exit(0)
	}()
	var configObject config.Config
	config.GetValues(&configObject)
	logger.StaringServer()
	vistedPath := path.GetFilePaths(configObject.RootPath, configObject.SkipDirectories,0)
	htmlFiles := path.GetFilePaths(configObject.RootPath, configObject.SkipDirectories,1)
	go watch.WatchFiles(vistedPath, &htmlFiles, configObject)
	server.Server(configObject.Port, &htmlFiles)
}
