package main

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/agents"
	"github.com/projects/threaded-company-simulation/config"
	"github.com/projects/threaded-company-simulation/factory"
)

var workers []*agents.Worker

func main() {
	factory.RunService()
	factory.RunLists()
	factory.RunLogger()
	factory.RunBoss()
	workers = factory.RunWorkers()
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
