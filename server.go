package main

import (
	ratelimiter "RateLimiter/rate_limiter"
	logger "RateLimiter/rate_limiter/Logs"
	Errors "RateLimiter/rate_limiter/errors"
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
		logger.LogRequest(r)
		http.Error(w, "Request denied", http.StatusTooManyRequests)
	}
}

func main() {
	// Set up the server
	http.HandleFunc("/", handler)

	// Start the server on port 4000
	logger.Info("Server is running on http://localhost:4000/")
	err := http.ListenAndServe(":4000", nil)
	Errors.HandleErr(Errors.Params{Err: err, Message: "Error starting server !", IsBlocking: true})
}
