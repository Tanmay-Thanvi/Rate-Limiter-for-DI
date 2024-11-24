package main

import (
	ratelimiter "RateLimiter/rate_limiter"
	Errors "RateLimiter/rate_limiter/errors"
	"RateLimiter/rate_limiter/limiter"
	"RateLimiter/rate_limiter/utils"
	"fmt"
	"log"
	"net/http"
)

func init() {
	limiter.Initialize()
}

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
	log.Println("Server is running on http://localhost:4000/")
	err := http.ListenAndServe(":4000", nil)
	Errors.HandleErr(Errors.Params{Err: err, Message: "Error starting server !", IsBlocking: true})
}
