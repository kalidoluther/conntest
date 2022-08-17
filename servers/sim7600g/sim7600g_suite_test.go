package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSim7600g(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sim7600g Suite")
}
