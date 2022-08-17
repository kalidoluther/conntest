package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CommandLine", func() {

	Describe("Command", func() {
		It("should create command from NewCommandLineFromBody", func() {
			Expect(NewCommandLineFromBody("+CSQ").Command()).To(Equal("AT+CSQ\r"))
		})
	})

})
