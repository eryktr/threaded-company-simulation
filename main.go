package main

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"

	"github.com/projects/threaded-company-simulation/factory"
)

func main() {
	list := make(chan factory.Job, config.TASKLIST_SIZE)
	warehouse := make(chan int, config.WAREHOUSE_SIZE)
	go factory.Ceo(list)
	for i := 0; i < config.NUM_WORKERS; i++ {
		go factory.Worker(list, warehouse)
	}
	for i := 0; i < config.NUM_CUSTOMERS; i++ {
		go factory.Customer(warehouse)
	}
	fmt.Scanln()

}
