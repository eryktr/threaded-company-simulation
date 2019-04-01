package main

import (
	"fmt"
	"strconv"

	"github.com/projects/threaded-company-simulation/agents"
)

func PrintMenu() {
	fmt.Println("[1] Show tasklist")
	fmt.Println("[2] Show warehouse")
	fmt.Println("[3] Show workers information")
	fmt.Println("Choice: ")
}

func GetChoice() int {
	var choice string
	fmt.Scanln(&choice)
	res, _ := strconv.Atoi(choice)
	return res
}

func ProcessChoice(choice int) {
	switch choice {
	case 1:
		agents.PrintTasklist()
	case 2:
		agents.PrintWarehouse()
	case 3:
		PrintWorkers()
	default:
		fmt.Println("Invalid choice:-(")
	}
}

func PrintWorkers() {
	for i := 0; i < len(workers); i++ {
		w := workers[i]
		fmt.Printf("Id: %d, Tasks done: %d, Is patient: %v\n", w.Id, w.CompletedTasks, w.IsPatient)
	}
}
