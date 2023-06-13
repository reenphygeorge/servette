package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/reenphygeorge/light-server/internal/logger"
	"github.com/reenphygeorge/light-server/internal/redirect"
)

func fileInterceptorHandler(fs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if(!strings.Contains(url,".")) {
			url = url+"index.html"
		}
		if strings.HasSuffix(url, ".html") {
			htmlContent, err := readFile("."+url)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			modifiedHTML := modifyHTML(htmlContent)
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, modifiedHTML)
			return
		}
		fs.ServeHTTP(w, r)
	})
}

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

func modifyHTML(htmlContent string) string {
	injectedScript := `<script>
	const socket = new WebSocket("ws://localhost:8001/ws");
	socket.addEventListener("open", (event) => {
	  console.log("Welcome to Light Server!");
	});
	socket.addEventListener("message", (event) => {
	  if(event.data === 'Reload'){
		location.reload();
	  }
	});</script>`
	modifiedContent := strings.Replace(string(htmlContent), "</body>", injectedScript+"</body>", 1)
	return modifiedContent
}

func Server(port string) {	
	fileServer := http.FileServer(http.Dir("."))
	interceptor := fileInterceptorHandler(fileServer)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn,err := SocketUpgrader(w,r)
        if err != nil {
            return
        } else {
			logger.StartAndReload(port)
		}
		defer conn.Close()
		HandleMessage(conn)
	})
	http.Handle("/", interceptor)
	redirect.OpenURL("http://localhost:"+port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Error()
	}
}