package limiter

import "RateLimiter/rate_limiter/algorithms"

type RateLimitersMap map[Level]*RateLimiter

var allRateLimiters = make(RateLimitersMap)

func init() {
	for _, Level := range Levels {

		algorithm, err := algorithms.NewRLAlgorithm(LimiterAlgoMapping[Level])
		if err != nil {
			panic(err)
		}

		allRateLimiters[Level] = &RateLimiter{
			Config:    "abcd",
			Algorithm: algorithm,
		}
	}
}

func GetAllRateLimiters() RateLimitersMap {
	return allRateLimiters
}
