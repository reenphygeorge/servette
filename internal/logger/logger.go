package logger

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		terminal := os.Getenv("COMSPEC")
		if strings.Contains(strings.ToLower(terminal), "powershell") {
			cmd = exec.Command("cmd", "/c", "cls")
		} else {
			cmd = exec.Command("powershell", "-command", "Clear-Host")
		}
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func StaringServer() {
	clearScreen()
	fmt.Println("\033[35mStarting Server...")
}

func Reloading() {
	clearScreen()
	fmt.Println("\033[35mReloading...")
}

func StartAndReload(port string) {
	clearScreen()
	fmt.Println("\033[35mServer running at port "+port, "\033[0m")
	fmt.Println("\033[33mReloaded ⚡\033[0m")
}

func Visit(port string) {
	clearScreen()
	fmt.Println("\033[35mServer Started")
	fmt.Println("\033[33mVisit http://localhost:"+port, "\033[0m")
}

func Error() {
	clearScreen()
	fmt.Println("\033[31mSomething's Wrong ", "\033[0m")
}
