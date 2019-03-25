package factory

import (
	"fmt"
	"strconv"

	"github.com/projects/threaded-company-simulation/agents"
)

func PrintMenu() {
	fmt.Println("[1] Show tasklist")
	fmt.Println("[2] Show warehouse")
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
	default:
		fmt.Println("Invalid choice:-(")
	}
}
