package algorithms

import (
	impl "RateLimiter/rate_limiter/algorithms/algos"
	Errors "RateLimiter/rate_limiter/errors"
)

func NewRLAlgorithm(algoType Algorithm /* & Config Values for algo */) (RLAlgorithm, error) {
	switch algoType {

	case Fixed_Counter, Sliding_Window, Token_Bucket: // for now
		return impl.NewSlidingWindow()
	default:
		return nil, Errors.ErrUnsupportedAlgo
	}

}
