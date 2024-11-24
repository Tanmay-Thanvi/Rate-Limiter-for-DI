package utils

import (
	"log"
	"net/http"
)

func Log(request *http.Request) {
	log.Println("Request : & Reject Reason need to be handled !", request)
}
