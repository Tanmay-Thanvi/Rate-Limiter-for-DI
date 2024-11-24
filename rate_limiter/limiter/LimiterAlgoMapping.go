package limiter

import algo "RateLimiter/rate_limiter/algorithms"

/*
This map needs to be completed i.e
all algos should be present for all levels
*/
var LimiterAlgoMapping = map[Level]algo.Algorithm{
	Organization: algo.Sliding_Window,
	Resource:     algo.Token_Bucket,
	API:          algo.Fixed_Counter,
	User:         algo.Sliding_Window,
}
