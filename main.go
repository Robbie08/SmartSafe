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
	http.HandleFunc("/on", unlockSafe)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":8080", nil)
}

// Default handle for when server is spun up. You can think of this
// as the home page
func defaultPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		log.Info("Page not found")
		return
	}
	log.Info("Someone hit the Default Page...")
	http.ServeFile(w, r, "piInterface.html")
}

// This function is in charge of gracefully shutting down
// our HTTP Server to prevent any external access to the pi
func shutdown(w http.ResponseWriter, r *http.Request) {
	log.Info("Shutting server down...")
	os.Exit(0)
}

// This function is in charge of handling request from client
// to open Smart Safe
func unlockSafe(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "piInterface.html")
	fmt.Println("Someone hit the unlock safe feature")
	piUtils.UnlockSafe()
}
