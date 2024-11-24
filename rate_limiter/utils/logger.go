package utils

import (
	"fmt"
	"net/http"
)

func Log(request *http.Request) {
	fmt.Println("Request : & Reject Reason need to be handled !", request)
}
