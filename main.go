package main

import (
	"fmt"                                       // golang formatting library -> for printing out to stdout
	"github.com/Robbie08/SmartSafe/pkg/piUtils" // contains our code to control pi
	"github.com/cgxeiji/servo"
	log "github.com/sirupsen/logrus" // library that helps with loging and monitoring
	"net/http"                       // library that provides us with code for creating HTTP server and request response logic
	"os"                             // gives us access to system calls
)

var password string = ""

func main() {
	// we will leave this close function in case ListenAndServe() unexpectedly stops
	defer servo.Close() // close out any connections with servos and pi-blaster
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting server...")
	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/pyclient", pyclient)
	http.ListenAndServe(":8080", nil)
}

func pyclient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseMultipartForm(0)
		message := r.FormValue("message")
		password = message /* Store our password as a global var so validation can use it*/
		fmt.Println("\n--------------- Received Python Client Message ---------------")
		fmt.Println("Password from pClient: ", message)
	default:
		fmt.Println("This handle only handles POST requests")
	}
}

// Default handle for when server is spun up. You can think of this
// as the home page
func defaultPage(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		/* Server our clients website */
		path := r.URL.Path
		if path == "/" {
			path = "./client/index.html"
		} else {
			path = "." + path
		}
		http.ServeFile(w, r, path)

	case "POST":
		/* Receive and unpackage payload sent via HTTP */
		r.ParseMultipartForm(0)
		message := r.FormValue("message")

		fmt.Println("\n--------------------- Authenticating -------------------------")
		fmt.Println("Password from SmartSafe client: ", message)

		/* Authenticate if the user entered the correct password */
		if password == "" {
			fmt.Println("Error: no Authentication provided by Python client")
		} else {
			if password == message {
				piUtils.UnlockSafe()
				fmt.Println("Safe Unlocked")
				password = "" /* make sure to reset password after successful auth*/
			} else {
				fmt.Println("Error: passwords do not match")
			}
		}
	default:
		fmt.Println("This service only supports GET and POST requests")
	}
}

// This function is in charge of gracefully shutting down
// our HTTP Server to prevent any external access to the pi
func shutdown(w http.ResponseWriter, r *http.Request) {
	log.Info("Shutting server down...")
	servo.Close() // close out any connections with servos and pi-blaster
	os.Exit(0)    // our system exited without any errors
}
