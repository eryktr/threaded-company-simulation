package factory

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"
)

func operation_to_ascii(operation Operator) string {
	switch operation {
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	default:
		return "*"
	}
}

var mode = config.MODE

func print_job(j Job, m string) {

	if (mode == 1 && m != "diagnostic") || (m != "added" && m != "assigned" && m != "diagnostic") {
		return
	}
	if m == "added" || m == "diagnostic" {
		fmt.Println("Task", j.first, operation_to_ascii(j.operation), j.second, "added to the list.")
	} else {
		fmt.Println("Job", j.first, operation_to_ascii(j.operation), j.second, "assigned to a worker")
	}
}

func print_product(p int, m string) {
	if (mode == 1 && m != "diagnostic") || (m != "stored" && m != "collected" && m != "diagnostic") {
		return
	}
	if m == "stored" || m == "diagnostic" {
		fmt.Println("Product", p, "Stored in the warehouse.")
	} else {
		fmt.Println("Product", p, "collected by a customer")
	}
}

func print_all_jobs(list chan Job) {
	lock_list()
	lock_warehouse()
	size := len(list)
	if size == 0 {
		println("No jobs at this moment.")
	} else {
		for i := 0; i < size; i++ {
			j := <-list
			//print("printing job")
			print_job(j, "diagnostic")
			list <- j
		}
	}
	unlock_list()
	unlock_warehouse()
}

func print_all_products(list chan int) {
	lock_list()
	lock_warehouse()

	size := len(list)
	if size == 0 {
		println("No products at this moment.")
	} else {
		for i := 0; i < size; i++ {
			j := <-list
			//print("printing job")
			print_product(j, "diagnostic")
			list <- j
		}
	}
	unlock_list()
	unlock_warehouse()
}

func print_config() {
	fmt.Println("CEO DELAY:", config.AVERAGE_CEO_DELAY)
	fmt.Println("WORKER DELAY:", config.AVERAGE_WORKER_DELAY)
	fmt.Println("CUSTOMER DELAY:", config.AVERAGE_CUSTOMER_DELAY)
	fmt.Println("TASKLIST SIZE:", config.TASKLIST_SIZE)
	fmt.Println("WAREHOUSE SIZE:", config.WAREHOUSE_SIZE)
}
