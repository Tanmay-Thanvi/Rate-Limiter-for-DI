package implementation

import (
	"log"
	"net/http"
	"time"
)

type slidingWindow struct {
	WindowSize   time.Duration
	RequestLimit int
}

func NewSlidingWindow( /* Rate Limiter Configs */ ) (*slidingWindow, error) {
	// Get this values from RateLimiter Config
	algo := &slidingWindow{
		WindowSize:   1 * time.Second,
		RequestLimit: 5,
	}
	return algo, nil
}

func (*slidingWindow) AllowRequest(Request *http.Request) (bool, error) {
	log.Println("Sliding Window Rate Limiting Strategy implemented !")
	return true, nil
}
