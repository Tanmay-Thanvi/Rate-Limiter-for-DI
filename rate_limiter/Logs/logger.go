package logs

import (
	"log"
	"net/http"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogRequest(request *http.Request) {
	Info.Println("Request : & Reject Reason need to be handled !", request)
}
