package algorithms

import "net/http"

// Interface of Rate Limiting Algorithm

type RLAlgorithm interface {
	AllowRequest(Request *http.Request) (bool, error)
}
