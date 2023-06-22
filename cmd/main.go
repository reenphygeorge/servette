package main

import (
	"github.com/reenphygeorge/light-server/internal/config-handle"
	watch "github.com/reenphygeorge/light-server/internal/file-watcher"
	"github.com/reenphygeorge/light-server/internal/logger"
	"github.com/reenphygeorge/light-server/internal/path"
	"github.com/reenphygeorge/light-server/internal/server"
)

func main() {
	var configObject config.Config
	config.GetValues(&configObject)
	logger.StaringServer()
	vistedPath := path.GetFilePaths(configObject.RootPath, configObject.SkipDirectories)
	go watch.WatchFiles(vistedPath)
	server.Server(configObject.Port)
}
