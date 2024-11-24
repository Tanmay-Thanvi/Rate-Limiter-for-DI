package model

import (
	algo "RateLimiter/rate_limiter/algorithms"
	Errors "RateLimiter/rate_limiter/errors"
	"net/http"
)

type RateLimiter struct {
	Config    string
	Algorithm algo.RLAlgorithm
}

func (RL *RateLimiter) EvaluateRequest(request *http.Request) bool {
	isAllowed, err := RL.Algorithm.AllowRequest(request)
	Errors.HandleErr(err, "Unable to evaluate Request")
	return isAllowed
}
