package main

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/agents"
	"github.com/projects/threaded-company-simulation/config"
	"github.com/projects/threaded-company-simulation/factory"
)

var workers []*agents.Worker
var serviceWorkers []*agents.ServiceWorker

func main() {
	factory.RunLists()
	factory.RunLogger()
	factory.RunBoss()
	factory.RunService()
	workers = factory.RunWorkers()
	serviceWorkers = factory.RunServiceWorkers(workers[0].MulltMachines, workers[0].AddMachines)
	factory.RunCustomers()
	if config.MODE == 1 {
		RunInputListener()
	} else {
		fmt.Scanln()
	}

}

func RunInputListener() {
	for {
		PrintMenu()
		choice := GetChoice()
		ProcessChoice(choice)
	}
}
