package logs

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	_info  *log.Logger
	_error *log.Logger
)

func init() {
	_info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	_error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// A helper function to check if a string contains a substring
func contains(str, substr string) bool {
	return len(str) >= len(substr) && len(substr) > 0 && fmt.Sprintf("%v", str[:len(substr)]) == substr
}

// Helper function to detect format specifiers in a string
func containsFormatSpecifiers(input string) bool {
	// A simple check for common format specifiers like %d, %s, %v, %f, %t, %x
	formatSpecifiers := []string{"%d", "%s", "%v", "%f", "%t", "%x"}
	for _, spec := range formatSpecifiers {
		if contains(input, spec) {
			return true
		}
	}
	return false
}

func common(logLevel *log.Logger, params ...interface{}) {
	if len(params) == 0 {
		return
	}

	firstArg, ok := params[0].(string)
	if ok && containsFormatSpecifiers(firstArg) {
		// Use fmt.Printf
		logLevel.Printf(firstArg, params[1:]...)
		logLevel.Println() // Add a newline since Printf does not append one
	} else {
		// Use fmt.Println
		logLevel.Println(params...)
	}
}

func Info(params ...interface{}) {
	common(_info, params...)
}

func Error(params ...interface{}) {
	common(_error, params...)
}

func LogRequest(request *http.Request) {
	Info("Request : & Reject Reason need to be handled !", request)
}
