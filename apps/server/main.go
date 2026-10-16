package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sprout/packages/engine"
	"time"
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
	NumUsers      int16              `json:"numusers"`
	ExecutionTime int16              `json:"executionTime"`
	Profile       engine.UserProfile `json:"profile"`
}

func start(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var data RunData

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid JSON payload, %s", err), http.StatusBadRequest)
		return
	}

	// will need to wrap this with parallelization
	// goroutines and for loops

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(6))

	for i := 0; i < int(data.NumUsers); i++ {
		go engine.SimulateUser(ctx, data.Profile)
	}

	select {
	case <-ctx.Done():
		log.Println("Got a cancel request")
		return
	case <-time.After(time.Second * time.Duration(data.ExecutionTime)):
		log.Println(data.ExecutionTime)
		log.Println("Timed out")
		cancel()
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
