package main

import (
	"testing"
)

func TestNetworkType(t *testing.T) {
	msg := getNetworkType()
	if msg != "tcp" {
		t.Fatalf(`getNetworkType("") = %q, want ""`, msg)
	}
}
