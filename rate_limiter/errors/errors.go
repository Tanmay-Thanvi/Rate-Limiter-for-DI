package Errors

import (
	"errors"
	"fmt"
)

var (
	ErrUnsupportedAlgo = errors.New("UnSupported Algorithm")
)

type Params struct {
	Err        error
	Message    string
	IsBlocking bool
}

func HandleErr(params Params) {
	if params.Err != nil {
		if params.Message == "" {
			params.Message = "No Message Provided"
		}

		fmt.Printf("Message : %v \nErr : %v\n", params.Message, params.Err.Error())

		if params.IsBlocking {
			panic(params.Err)
		}
	}
}
