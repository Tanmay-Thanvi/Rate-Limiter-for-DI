package ratelimiter

import (
	logger "RateLimiter/rate_limiter/Logs"
	"RateLimiter/rate_limiter/model"
	service "RateLimiter/rate_limiter/services"
	"net/http"
	"sync"
)

func Decider(request *http.Request) bool {
	allLimiters := service.GetAllRateLimiters()
	var wg sync.WaitGroup
	resultChannel := make(chan bool, len(allLimiters))
	trueCount := 0

	// Launch concurrent goroutines for each rate limiter
	for _, limiter := range allLimiters {
		wg.Add(1)
		go func(limiter *model.RateLimiter) {
			defer wg.Done()
			// Call EvaluateRequest for each RateLimiter and send the result to the channel
			resultChannel <- limiter.EvaluateRequest(request)
		}(limiter)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the resultChannel after all goroutines are done
	close(resultChannel)

	// Decision parameter
	acceptReq := true

	// Count how many times the result was true
	for result := range resultChannel {
		if !result {
			acceptReq = false
		} else {
			trueCount++
		}
	}

	// Log request to kibana for monitoring
	if !acceptReq {
		logger.Info("True Count : ", trueCount)
	}

	return acceptReq
}
