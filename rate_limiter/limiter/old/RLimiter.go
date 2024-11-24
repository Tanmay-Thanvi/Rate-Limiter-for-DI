package limiter

import "net/http"

type RLimiter interface {
	EvaluateRequest(request *http.Request) bool
}
