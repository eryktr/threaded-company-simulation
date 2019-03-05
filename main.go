package main

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/factory"
)

func main() {
	list := make(chan factory.Job, 200)
	warehouse := make(chan int, 500)
	go factory.Ceo(list)
	for i := 0; i < 1; i++ {
		go factory.Worker(list, warehouse)
	}
	for i := 0; i < 1; i++ {
		go factory.Customer(warehouse)
	}
	fmt.Scanln()

}
