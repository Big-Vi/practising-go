package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"math/rand"
)

func main() {
	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a channel for signaling server shutdown
	var ch = make(chan bool)

	// Start the HTTP server in a goroutine
	go startServer(ctx, ch)

	// Capture OS signals (SIGINT and SIGTERM)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sig:
		// Handle OS signal to gracefully shut down the server
		fmt.Println("Shutting down due to interrupt")
		cancel() // Cancel the context to signal server shutdown
		os.Exit(0)
	}

	// Wait for the server to complete its shutdown
	<-ch
}

// startServer starts the HTTP server and handles incoming requests.
func startServer(ctx context.Context, ch chan bool) {
	// Define the homepage handler function
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Simulate network response - Set random sleep time
		// rand.Seed(time.Now().UnixNano())
		minDuration := 1
		maxDuration := 6

		sleepDuration := time.Duration(rand.Intn(maxDuration - minDuration + 1)+ minDuration)  * time.Second
		fmt.Println(sleepDuration)
		time.Sleep(sleepDuration)

		select {
		case <-ctx.Done():
			// Handle server errors when the context is canceled
			err := ctx.Err()
			fmt.Println("Server error:", err)
			ch <- true
			// http.Error(w, err.Error(), http.StatusInternalServerError)
		default:
			// Continue processing the request
			fmt.Println("Received signal.")
			fmt.Fprintf(w, "home")
		}
	})

	// Start the HTTP server on port 8090
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Server error:", err)
		ch <- true
	}
}

