package main

import (
	"fmt"                                       // golang formatting library -> for printing out to stdout
	"github.com/Robbie08/SmartSafe/pkg/piUtils" // contains our code to control pi
	log "github.com/sirupsen/logrus"            // library that helps with loging and monitoring
	"net/http"                                  // library that provides us with code for creating HTTP server and request response logic
	"os"                                        // gives us access to system calls
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting server...")
	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":8080", nil)
}

// Default handle for when server is spun up. You can think of this
// as the home page
func defaultPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		path := r.URL.Path
		fmt.Println(path)

		if path == "/" {
			path = "./client/index.html"
		} else {
			path = "." + path
		}
		http.ServeFile(w, r, path)

	case "POST":
		r.ParseMultipartForm(0)
		message := r.FormValue("message")
		fmt.Println("--------------------------------------------")
		fmt.Println("Msg from client: ", message)
		// we can make the calll to our piUtils.UnlockSafe() here!
		piUtils.UnlockSafe()
	default:
		fmt.Println("This service only supports GET and POST requests")
	}
}

// This function is in charge of gracefully shutting down
// our HTTP Server to prevent any external access to the pi
func shutdown(w http.ResponseWriter, r *http.Request) {
	log.Info("Shutting server down...")
	os.Exit(0)
}
