package main

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/agents"
	"github.com/projects/threaded-company-simulation/config"
)

func main() {

	go agents.SynchronizeWarehouse()
	go agents.SynchronizeTaskList()

	Boss := agents.Boss{1, agents.TaskListWrite}
	go Boss.Run()

	for i := 0; i < config.NUM_WORKERS; i++ {
		w := agents.Worker{i, agents.TaskListRead, agents.WarehouseWrite}
		go w.Run()
	}
	for i := 0; i < config.NUM_CUSTOMERS; i++ {
		cust := agents.Customer{i, agents.WarehouseRead}
		go cust.Run()
	}
	if config.MODE == 1 {
		//factory.InputListener()
	} else {
		fmt.Scanln()
	}

}
