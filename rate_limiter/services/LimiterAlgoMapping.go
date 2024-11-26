package service

import algo "RateLimiter/rate_limiter/algorithms"

/*
This map needs to be completed i.e
all algos should be present for all levels
*/
var LevelRLAlgoMapping = map[Level]algo.Algorithm{
	Organization: algo.Sliding_Window,
	Resource:     algo.Sliding_Window,
	API:          algo.Sliding_Window,
	User:         algo.Sliding_Window,
}