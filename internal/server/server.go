package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func fileInterceptorHandler(fs http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Intercept requests for HTML files only
		url := r.URL.Path
		if(!strings.Contains(url,".")) {
			url = url+"index.html"
		}
		if strings.HasSuffix(url, ".html") {
			// Read the HTML file
			htmlContent, err := readFile("."+url)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Modify the HTML content
			modifiedHTML := modifyHTML(htmlContent)

			// Set the appropriate Content-Type header
			w.Header().Set("Content-Type", "text/html")

			// Write the modified HTML content to the response
			fmt.Fprint(w, modifiedHTML)
			return
		}

		// Serve other files using the default file server
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
	// Modify the HTML content as per your requirements
	// Replace this with your own implementation
	injectedScript := `<script>const socket = new WebSocket("ws://localhost:8001/ws");

	// Connection opened
	socket.addEventListener("open", (event) => {
	  socket.send("Hello Server!");
	});

	// Listen for messages
	socket.addEventListener("message", (event) => {
	  console.log("Message from server ", event.data);
	});</script>`
	modifiedContent := strings.Replace(string(htmlContent), "</body>", injectedScript+"</body>", 1)
	return modifiedContent
}

func Server() {	
	fileServer := http.FileServer(http.Dir("."))
	
	interceptor := fileInterceptorHandler(fileServer)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn,err := SocketUpgrader(w,r)
        if err != nil {
            log.Println("Failed to upgrade to WebSocket:", err)
            return
        } else {
			log.Print("Connection Success!")
		}
		defer conn.Close()
		HandleMessage(conn)
	})

	http.Handle("/", interceptor)
	
	fmt.Printf("Starting server at port 8000\n")

	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}