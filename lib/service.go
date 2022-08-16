package service

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
	for _, service := range services {
		err := service.Init()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Starting services")
	for _, service := range services {
		go func() {
			err := service.Start()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	<-make(chan int)
}
