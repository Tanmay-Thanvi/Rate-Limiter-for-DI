package limiter

import (
	logs "RateLimiter/rate_limiter/Logs"
	"RateLimiter/rate_limiter/algorithms"
	Errors "RateLimiter/rate_limiter/errors"
	"RateLimiter/rate_limiter/model"
)

type RateLimitersMap map[Level]*model.RateLimiter

var allRateLimiters = make(RateLimitersMap)

func init() {
	logs.Info.Println("Initializing All Rate Limiters !")
	for _, Level := range Levels {
		algorithm, err := algorithms.NewRLAlgorithm(LimiterAlgoMapping[Level])
		Errors.HandleErr(Errors.Params{Err: err, Message: "Algorithm Unavailable !", IsBlocking: true})

		allRateLimiters[Level] = &model.RateLimiter{
			Config:    "abcd",
			Algorithm: algorithm,
		}
	}
}

func GetAllRateLimiters() RateLimitersMap {
	return allRateLimiters
}
