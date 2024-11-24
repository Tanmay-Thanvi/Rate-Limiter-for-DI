package Errors

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedAlgo = errors.New("UnSupported Algorithm")
)

func HandleErr(err error, message string) {
	if err != nil {
		fmt.Printf("Message : %v, Err : %v\n", message, err.Error())
		panic(err)
	}
}
