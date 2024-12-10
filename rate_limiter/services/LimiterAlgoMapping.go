package service

import algo "RateLimiter/rate_limiter/algorithms"

/*
This map will be pulled from DB on init only i.e
all algos will be mapped for all levels
*/
var LevelRLAlgoMapping = map[Level]algo.Algorithm{
	Organization: algo.Sliding_Window,
	Resource:     algo.Sliding_Window,
	API:          algo.Sliding_Window,
	User:         algo.Sliding_Window,
}
