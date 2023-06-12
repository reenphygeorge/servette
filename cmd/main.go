package main

import (
	watch "github.com/reenphygeorge/light-server/internal/file-watcher"
	logger "github.com/reenphygeorge/light-server/internal/logger"
	path "github.com/reenphygeorge/light-server/internal/path"
	ports "github.com/reenphygeorge/light-server/internal/port-checker"
	server "github.com/reenphygeorge/light-server/internal/server"
)

func main() {
	logger.StaringServer()
	vistedPath := path.GetFilePaths()
	startPort := 5500
	endPort := 5510
	port,_ := ports.GetFreePort(startPort,endPort)
	go watch.WatchFiles(vistedPath)
	server.Server(port)
}