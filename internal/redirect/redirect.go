package redirect

import (
	"os/exec"
	"runtime"

	"github.com/reenphygeorge/light-server/internal/logger"
)

func OpenURL(url string) {
	var err error

	switch runtime.GOOS {
		case "windows":
			err = exec.Command("cmd", "/c", "start", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = exec.Command("xdg-open", url).Start()
	}

	if err != nil {
		logger.Error()
	}
}