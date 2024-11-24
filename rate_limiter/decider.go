package ratelimiter

import (
	"RateLimiter/rate_limiter/limiter"
	"RateLimiter/rate_limiter/model"
	"fmt"
	"net/http"
	"sync"
)

func Decider(request *http.Request) bool {
	allLimiters := limiter.GetAllRateLimiters()
	var wg sync.WaitGroup
	resultChannel := make(chan bool, len(allLimiters))
	trueCount := 0

	// Launch concurrent goroutines for each rate limiter
	for _, limiter := range allLimiters {
		wg.Add(1)
		go func(limiter *model.RateLimiter) {
			defer wg.Done()
			// Call EvaluateRequest for each RateLimiter and send the result to the channel
			result := limiter.EvaluateRequest(request)
			resultChannel <- result
		}(limiter)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the resultChannel after all goroutines are done
	close(resultChannel)

	// Count how many times the result was true
	for result := range resultChannel {
		if result {
			trueCount++
		}
	}

	fmt.Println("True Count : ", trueCount)

	return trueCount > 0
}
