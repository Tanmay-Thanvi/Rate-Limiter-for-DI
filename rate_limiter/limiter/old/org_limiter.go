package limiter

import (
	algo "RateLimiter/rate_limiter/algorithms"
	"net/http"
)

type orgLimiter struct {
	algorithm algo.RLAlgorithm
	config    string // for now its string after that will give some struct
}

func (ol *orgLimiter) EvaluateRequest(request *http.Request) bool {
	isAllowed, err := ol.algorithm.AllowRequest(request)
	if err != nil {
		panic(err)
	}
	return isAllowed
}
