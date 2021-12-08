package main

import (
	"fmt"      // golang formatting library -> for printing out to stdout
	"net/http" // library that provides us with code for creating HTTP server and request response logic
	"os"       // gives us access to system calls
	"os/exec"
	"time"

	// contains our code to control pi
	"github.com/cgxeiji/servo"
	log "github.com/sirupsen/logrus" // library that helps with loging and monitoring
)

var password string = "" // python client writes to this, authenticator reads from this
var safeId string = ""   // python client writes to this, authenticator reads from this
var isValid string = "false"

func main() {
	// we will leave this close function in case ListenAndServe() unexpectedly stops
	defer servo.Close() // close out any connections with servos and pi-blaster
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting server...")
	http.HandleFunc("/", defaultPage)
	http.HandleFunc("/shutdown", shutdown)
	http.HandleFunc("/pyclient", pyclient)
	http.HandleFunc("/rfidClient", rfidClient)
	http.ListenAndServe(":8080", nil)
}

// This function will handle any incoming requests from the python client.
// Inside it should unpackage the data and save our password.
func rfidClient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Unpackage and store our password and safeId sent by python client
		r.ParseMultipartForm(0)

		fmt.Println("\n--------------- Received RFID Client Message ---------------")

		cmd := exec.Command("python3", "./pyClient.py")
		fmt.Println("Running python program")
		err := cmd.Run()
		if err != nil {
			fmt.Println("Finished :", err)
		}
		time.Sleep(45 * time.Second)
		// clients want to hit this endpoint http://localhost:8080/rfidClient
		w.Write([]byte(isValid))
		isValid = "false"
	default:
		fmt.Println("This handle only handles POST requests")
	}
}

// This function will handle any incoming requests from the python client.
// Inside it should unpackage the data and save our password.
func pyclient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Unpackage and store our password and safeId sent by python client
		r.ParseMultipartForm(0)
		password = r.FormValue("message")
		safeId = r.FormValue("id")

		fmt.Println("\n--------------- Received Python Client Message ---------------")
		fmt.Println("Password from pClient: ", password)
		fmt.Println("Id from pClient: ", safeId)
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
		passwdSFClient := r.FormValue("message") // The password typed in by user to SmartSafe Client
		idSFClient := "003349"                   // Dummy for now but This should be the value passed by our SmartSafe Client(typed in by user)

		fmt.Println("\n--------------------- Authenticating -------------------------")
		fmt.Println("Password from SmartSafe client	: ", passwdSFClient)
		fmt.Println("Password from Python Client	: ", password)
		fmt.Println("Id from SmartSafe client: ", idSFClient)

		/* Authenticate if the user entered the correct password */
		if password == "" || safeId == "" {
			fmt.Println("Error: no Authentication provided by Python client")
		} else {
			if password == passwdSFClient && safeId == idSFClient {
				//piUtils.UnlockSafe() // contact our hardware program that unlocks safe
				fmt.Println("Authenticated.... will send response back with authentication")
				isValid = "true"
				password = "" /* make sure to reset password after successful auth*/
			} else {
				fmt.Println("Error: incorrect auth credentials")
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
