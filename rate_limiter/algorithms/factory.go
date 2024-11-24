package algorithms

import (
	impl "RateLimiter/rate_limiter/algorithms/algos"
	Errors "RateLimiter/rate_limiter/errors"
)

func NewRLAlgorithm(algoType Algorithm /* & Config Values for algo */) (RLAlgorithm, error) {
	switch algoType {

	case Sliding_Window:
		return impl.NewSlidingWindow()
	default:
		return nil, Errors.ErrUnsupportedAlgo
	}

}
