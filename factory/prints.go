package factory

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"
)

var mode = config.MODE

func print_job(j Job, m string) {
	if mode == 1 || (m != "added" && m != "assigned") {
		return
	}
	if m == "added" {
		fmt.Println("Task", j.first, j.operation, j.second, "added to the list.")
	} else {
		fmt.Println("Job", j.first, j.operation, j.second, "assigned to a worker")
	}
}

func print_product(p int, m string) {
	if mode == 1 || (m != "stored" && m != "collected") {
		return
	}
	if m == "stored" {
		fmt.Println("Product", p, "Stored in the warehouse.")
	} else {
		fmt.Println("Product", p, "collected by a customer")
	}
}

func print_all_jobs(list chan Job) {
	lock_list()
	lock_warehouse()
	size := len(list)
	for i := 0; i < size; i++ {
		j := <-list
		print_job(j, "added")
		list <- j
	}
	unlock_list()
	unlock_warehouse()
}
