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

func CreateMultMachines() []agents.MultiplicationMachine {
	var machines = make([]agents.MultiplicationMachine, config.NUM_MULT_MACHINES)
	for i := 0; i < config.NUM_MULT_MACHINES; i++ {
		machine := agents.MultiplicationMachine{
			Id:     i,
			Input:  make(chan agents.MachineWriteOp, 1),
			Logger: agents.LogChannel}
		machines = append(machines, machine)
	}
	return machines
}

func CreateAdditionMachines() []agents.AdditionMachine {
	var machines = make([]agents.AdditionMachine, config.NUM_MULT_MACHINES)
	for i := 0; i < config.NUM_MULT_MACHINES; i++ {
		machine := agents.AdditionMachine{
			Id:     i,
			Input:  make(chan agents.MachineWriteOp, 1),
			Logger: agents.LogChannel}
		machines = append(machines, machine)
	}
	return machines
}

func RunWorkers() {
	additionMachines := CreateAdditionMachines()
	multiplicationMachines := CreateMultMachines()
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
