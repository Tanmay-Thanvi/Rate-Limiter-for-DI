package ratelimiter

import (
	"RateLimiter/rate_limiter/limiter"
	"net/http"
)

func Decider(request *http.Request) {
	for _, RateLimiter := range limiter.GetAllRateLimiters() {
		go RateLimiter.EvaluateRequest(request)
	}
}
