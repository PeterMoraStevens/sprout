package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sprout/packages/engine"
)

// this server needs to handle recieving a request to start a test run
// we'll be provided some json map data
// this data will tell us what the user profile will look like
// we'll then be able to pass this on to our engine and begin making runs
// we'll use this continiously running server to track the number of process we're running
// until we reach our goal of ___ num of digital users

// start run should extract much of the intformation which we require
// such as the runs type (we should expand this in iterations)

type RunData struct {
	NumUsers int16              `json:"numusers"`
	Profile  engine.UserProfile `json:"profile"`
}

func start(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var data RunData

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Starting Run with %d users, %+v", data.NumUsers, data.Profile)
}

func main() {
  const port = 8090
	portString := fmt.Sprintf(":%d", port)

	http.HandleFunc("/start", start)

	fmt.Printf("Listening on port %d\n", port)
	http.ListenAndServe(portString, nil)
}
