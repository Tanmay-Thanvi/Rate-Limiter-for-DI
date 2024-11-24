package logs

import (
	"log"
	"net/http"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func LogRequest(request *http.Request) {
	Info.Println("Request : & Reject Reason need to be handled !", request)
}
