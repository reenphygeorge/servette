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
	fmt.Println("\033[35mStarting Server...\033[0m")
}

func Reloading() {
	clearScreen()
	fmt.Println("\033[35mReloading...\033[0m")
}

func displayHTMLPaths(htmlFiles *[]string) {
	fmt.Println("\033[36m\nAvailable Paths: \033[0m")
	if len(*htmlFiles) == 0 {
		fmt.Println("\033[31mN/A\033[0m")
	} else {
		for _,htmlFile := range(*htmlFiles) {
			fmt.Printf("\033[36m\t/%s\n\033[0m",htmlFile)
		}
	}
}

func StartAndReload(port string, htmlFiles *[]string) {
	clearScreen()
	fmt.Println("\033[35mServer running at port "+port, "\033[0m")
	fmt.Println("\033[33mReloaded ⚡\033[0m")
	displayHTMLPaths(htmlFiles)
}

func Visit(port string, htmlFiles *[]string) {
	clearScreen()
	fmt.Println("\033[35mServer Started")
	fmt.Println("\033[33mVisit http://localhost:"+port, "\033[0m")
	displayHTMLPaths(htmlFiles)
}

func Error() {
	clearScreen()
	fmt.Println("\033[31mSomething's Wrong ", "\033[0m")
}

func Exit() {
	fmt.Println("\033[35m\nStopping Gracefully...", "\033[0m")
}