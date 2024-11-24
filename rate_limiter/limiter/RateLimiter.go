package limiter

import (
	algo "RateLimiter/rate_limiter/algorithms"
	"net/http"
)

type RateLimiter struct {
	Config    string
	Algorithm algo.RLAlgorithm
}

func (RL *RateLimiter) EvaluateRequest(request *http.Request) bool {
	isAllowed, err := RL.Algorithm.AllowRequest(request)
	if err != nil {
		panic(err)
	}
	return isAllowed
}
