package main

import (
	ratelimiter "RateLimiter/rate_limiter"
	Errors "RateLimiter/rate_limiter/errors"
	"RateLimiter/rate_limiter/utils"
	"fmt"
	"net/http"
)

// Handler function to process requests if the Decider allows it
func handler(w http.ResponseWriter, r *http.Request) {
	if ratelimiter.Decider(r) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Request successfully processed!")
	} else {
		/*
			Log request to kibana for network forensics
			& send 429 status code to client
		*/
		utils.Log(r)
		http.Error(w, "Request denied", http.StatusTooManyRequests)
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