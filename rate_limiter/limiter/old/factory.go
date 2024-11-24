package limiter

import algo "RateLimiter/rate_limiter/algorithms"

func NewLimiter(Level level, algo algo.RLAlgorithm) RLimiter {
	switch Level {
	case Organization:
		return &orgLimiter{
			algorithm: algo,
			config:    "",
		}
	default:
		return nil
	}

}
