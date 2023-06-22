package main

import (
	"github.com/reenphygeorge/light-server/internal/config-handle"
	watch "github.com/reenphygeorge/light-server/internal/file-watcher"
	logger "github.com/reenphygeorge/light-server/internal/logger"
	path "github.com/reenphygeorge/light-server/internal/path"
	server "github.com/reenphygeorge/light-server/internal/server"
)

func main() {
	var configObject config.Config
	config.GetValues(&configObject)
	logger.StaringServer()
	vistedPath := path.GetFilePaths(configObject.RootPath, configObject.SkipDirectories)
	go watch.WatchFiles(vistedPath)
	server.Server(configObject.Port)
}
