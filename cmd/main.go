package main

import (
	watch "github.com/reenphygeorge/light-server/internal/file-watcher"
	logger "github.com/reenphygeorge/light-server/internal/logger"
	path "github.com/reenphygeorge/light-server/internal/path"
	server "github.com/reenphygeorge/light-server/internal/server"
)

func main() {
	logger.StaringServer()
	vistedPath := path.GetFilePaths()
	go watch.WatchFiles(vistedPath)
	server.Server()
}