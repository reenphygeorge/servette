package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/reenphygeorge/light-server/internal/logger"
)

/*
	Global port.
	Will hold port number read from config.
	Used to avoid argument passing through out the functions.
*/
var globalPort int

// Intercept http request and call modifyHTML function.
func fileInterceptorHandler(fs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if !strings.Contains(url, ".") {
			url = url + "index.html"
		}
		if strings.HasSuffix(url, ".html") {
			htmlContent, err := readFile("." + url)
			if err != nil {
				http.Error(w, "Page not found!", http.StatusNotFound)
				return
			}
			modifiedHTML := modifyHTML(htmlContent)
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, modifiedHTML)
			return
		} else if strings.HasSuffix(url, ".css") {
			w.Header().Set("Content-Type", "text/css")
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")
		}
		fs.ServeHTTP(w, r)
	})
}

// Read and return html files.
func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Modify html files and add inject socket connection code.
func modifyHTML(htmlContent string) string {
	injectedScript := fmt.Sprintf(`<script>
	const socket = new WebSocket("ws://localhost:%s/ws");
	socket.addEventListener("open", (event) => {
	  console.log("Welcome to Light Server!");
	});
	socket.addEventListener("message", (event) => {
	  if(event.data === 'Reload'){
		location.reload();
	  }
	});</script>`, strconv.Itoa(globalPort))
	modifiedContent := strings.Replace(string(htmlContent), "</body>", injectedScript+"</body>", 1)
	return modifiedContent
}

// Server main functions.
func Server(port int, htmlFiles *[]string) {
	globalPort = port
	fileServer := http.FileServer(http.Dir("."))
	interceptor := fileInterceptorHandler(fileServer)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := SocketUpgrader(w, r)
		if err != nil {
			return
		} else {
			logger.StartAndReload(strconv.Itoa(globalPort), htmlFiles)
		}
		defer conn.Close()
		HandleMessage(conn)
	})
	http.Handle("/", interceptor)
	serve(htmlFiles)
}

// Starting server at available port
func serve(htmlFiles *[]string) {
	time.Sleep(time.Second / 3)
	logger.Visit(strconv.Itoa(globalPort),htmlFiles)
	err := http.ListenAndServe(":"+strconv.Itoa(globalPort), nil)
	if err != nil {
		globalPort++
		logger.Visit(strconv.Itoa(globalPort),htmlFiles)
		logger.StartAndReload(strconv.Itoa(globalPort), htmlFiles)
		serve(htmlFiles)
	}
}
