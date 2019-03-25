package main

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"
	"github.com/projects/threaded-company-simulation/factory"
)

func main() {
	factory.RunLists()
	factory.RunLogger()
	factory.RunBoss()
	factory.RunWorkers()
	factory.RunCustomers()
	if config.MODE == 1 {
		factory.RunInputListener()
	} else {
		fmt.Scanln()
	}

}
