package limiter

import (
	algo "RateLimiter/rate_limiter/algorithms"
)

type level int
type levelAlgoMapping map[level]algo.Algorithm

const (
	Organization = iota
	Resource
	API
	User
)

var limiterMap levelAlgoMapping = levelAlgoMapping{
	Organization: algo.Sliding_Window,
	Resource:     algo.Token_Bucket,
}

func GetLimiters() map[level]RLimiter {
	var limiters map[level]RLimiter = make(map[level]RLimiter)
	for rLevel, algoType := range limiterMap {
		algorithm, err := algo.NewRLAlgorithm(algoType)
		if err != nil {
			panic(err)
		}
		limiters[rLevel] = NewLimiter(rLevel, algorithm)
	}
	return limiters
}
