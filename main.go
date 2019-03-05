package main

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"

	"github.com/projects/threaded-company-simulation/factory"
)

func main() {

	go factory.Ceo(factory.List)
	for i := 0; i < config.NUM_WORKERS; i++ {
		go factory.Worker(factory.List, factory.Warehouse)
	}
	for i := 0; i < config.NUM_CUSTOMERS; i++ {
		go factory.Customer(factory.Warehouse)
	}
	fmt.Scanln()

}
