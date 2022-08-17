package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sim7600g", func() {

	Describe("NewSim7600GClient", func() {
		It("should create a new Sim7600GClient struct from port and baud", func() {
			Expect(*NewSim7600GClient("COM34", 115200)).To(BeEquivalentTo(Sim7600GClient{
				SerialPort: "COM34",
				SerialBaud: 115200,
			}))
		})
	})

})
