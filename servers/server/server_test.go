package main

import (
	"testing"
)

func TestGetServerHost(t *testing.T) {
	msg := getNetworkType()
	if msg != "tcp" {
		t.Fatalf(`getNetworkType("") = %q, want ""`, msg)
	}
}
