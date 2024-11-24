package main

import (
	ratelimiter "RateLimiter/rate_limiter"
	Errors "RateLimiter/rate_limiter/errors"
	"fmt"
	"net/http"
)

// Handler function to process requests if the Decider allows it
func handler(w http.ResponseWriter, r *http.Request) {
	// Call the Decider function
	if ratelimiter.Decider(r) {
		// Process the request (here, we're just sending a success message)
		w.WriteHeader(http.StatusOK) // 200 OK
		fmt.Fprintf(w, "Request successfully processed!")
	} else {
		// If Decider returns false, send 429 status code
		http.Error(w, "Request denied", http.StatusTooManyRequests) // 429
	}
}

func main() {
	// Set up the server
	http.HandleFunc("/", handler)

	// Start the server on port 4000
	fmt.Println("Server is running on http://localhost:4000/")
	err := http.ListenAndServe(":4000", nil)
	Errors.HandleErr(err, "Error starting server !")
}
