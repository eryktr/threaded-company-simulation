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
		go machine.RunMultiplicationMachine()
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
		go machine.RunAdditionMachine()
		machines = append(machines, machine)
	}
	return machines
}

func RunWorkers() {
	additionMachines := CreateAdditionMachines()
	multiplicationMachines := CreateMultMachines()
	for i := 0; i < config.NUM_WORKERS; i++ {
		//outcome := rand.Intn(100)
		// var isPatient bool
		// if outcome < 50 {
		// 	isPatient = true
		// } else {
		// 	isPatient = false
		// }
		w := agents.Worker{
			Id:             i,
			TaskList:       agents.TaskListRead,
			Warehouse:      agents.WarehouseWrite,
			Logger:         agents.LogChannel,
			MulltMachines:  multiplicationMachines,
			AddMachines:    additionMachines,
			CompletedTasks: 0,
			IsPatient:      true}
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
