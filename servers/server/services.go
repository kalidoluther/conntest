package main

import (
	"fmt"
	"log"
)

var services = make([]Service, 0)

type Service interface {
	Init() error
	Start() error
}

func Add(s Service) {
	services = append(services, s)
}

func RunAll() {
	fmt.Println("Initializing services")
	for _, s := range services {
		err := s.Init()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Starting services")
	for _, serv := range services {
		go func(s Service) {
			err := s.Start()
			if err != nil {
				log.Fatal(err)
			}
		}(serv)
	}

	<-make(chan int)
}
