package factory

import (
	"fmt"
	"strconv"
)

func display_manual() {
	println("[1] Print task list")
	println("[2] Print warehouse")
	println("[3] Print config settings")
}

func get_choice() int {
	var input string
	fmt.Scanln(&input)
	value, _ := strconv.Atoi(input)
	return value
}

func evaluate_choice(choice int) {
	switch choice {
	case 1:
		print_all_jobs(List)
	case 2:
		print_all_products(Warehouse)
	case 3:
		print_config()
	default:
		println("Illegal choice")
	}

}

func InputListener() {
	for {
		display_manual()
		choice := get_choice()
		evaluate_choice(choice)

	}
}
