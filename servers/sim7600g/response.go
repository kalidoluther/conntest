package main

import "fmt"

type CommandResponse struct {
	header  string
	result  string
	trailer string
}

func NewCommandResponseFromData(data string) *CommandResponse {
	fmt.Printf("Constructing CommandResponse from data:\n'%v'\n", data)
	return &CommandResponse{}
}

func (r *CommandResponse) Result() string {
	return r.result
}

//+27841000000
