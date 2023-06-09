package main

import (
	server "github.com/reenphygeorge/light-server/internal/server"
)

func main() {
	// vistedPath := path.GetFilePaths()
	// go watch.WatchFiles(vistedPath)
	server.Server()
}