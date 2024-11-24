package limiter

import algo "RateLimiter/rate_limiter/algorithms"

var LimiterAlgoMapping = map[Level]algo.Algorithm{
	Organization: algo.Sliding_Window,
}
