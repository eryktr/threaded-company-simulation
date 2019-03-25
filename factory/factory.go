package factory

import (
	"github.com/projects/threaded-company-simulation/agents"
	"github.com/projects/threaded-company-simulation/config"
)

func RunLists() {
	go agents.SynchronizeWarehouse()
	go agents.SynchronizeTaskList()
}

func RunLogger() {
	go agents.RunLogger()
}

func RunBoss() {
	Boss := agents.Boss{1, agents.TaskListWrite, agents.LogChannel}
	go Boss.Run()
}

func RunWorkers() {
	for i := 0; i < config.NUM_WORKERS; i++ {
		w := agents.Worker{i, agents.TaskListRead, agents.WarehouseWrite, agents.LogChannel}
		go w.Run()
	}
}

func RunCustomers() {
	for i := 0; i < config.NUM_CUSTOMERS; i++ {
		cust := agents.Customer{i, agents.WarehouseRead, agents.LogChannel}
		go cust.Run()
	}
}

func RunInputListener() {
	for {
		PrintMenu()
		choice := GetChoice()
		ProcessChoice(choice)
	}
}
