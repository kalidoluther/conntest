package main

import (
	"fmt"
	"strings"
)

type CommandLine struct {
	prefix      string
	body        string
	termination string
}

func NewCommandLineFromBody(body string) *CommandLine {
	const PREFIX = "AT"
	const TERMINATION = string('\r')

	return &CommandLine{
		prefix:      PREFIX,
		body:        body,
		termination: TERMINATION,
	}
}

func (c *CommandLine) Command() string {
	return fmt.Sprintf("%s%s%s", c.prefix, c.body, c.termination)
}

func HasResultCode(resp string) bool {
	// check for OK (there are more to check)
	return strings.Contains(resp, "OK")
}
