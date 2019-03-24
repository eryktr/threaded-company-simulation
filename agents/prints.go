package agents

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"
)

func Operation_to_ascii(operation Operator) string {
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

func Print_job_added(j Job) {
	if mode == 1 {
		return
	}
	fmt.Println("Task", j.first, Operation_to_ascii(j.operation), j.second, "added to the list.")

}

func Print_job_fetched(j Job) {
	if mode == 1 {
		return
	}
	fmt.Println("Job", j.first, Operation_to_ascii(j.operation), j.second, "assigned to a worker")
}

func Print_job_diagnostic(j Job) {
	fmt.Println("Job", j.first, Operation_to_ascii(j.operation), j.second)
}

func Print_product_added(product int) {
	if mode == 1 {
		return
	}
	fmt.Println("Product", product, "Stored in the warehouse.")
}

func Print_product_collected(product int) {
	if mode == 1 {
		return
	}
	fmt.Println("Product", product, "collected by a customer")
}

func Print_product_diagnostic(product int) {
	fmt.Println("Product", product)
}

func Print_all_jobs(list chan Job) {
	size := len(list)
	if size == 0 {
		println("No jobs at this moment.")
	} else {
		for i := 0; i < size; i++ {
			j := <-list
			Print_job_diagnostic(j)
			list <- j
		}
	}
}

func Print_all_products(list chan int) {

	size := len(list)
	if size == 0 {
		println("No products at this moment.")
	} else {
		for i := 0; i < size; i++ {
			j := <-list
			Print_product_diagnostic(j)
			list <- j
		}
	}

}

func Print_config() {
	fmt.Println("CEO DELAY:", config.AVERAGE_CEO_DELAY)
	fmt.Println("WORKER DELAY:", config.AVERAGE_WORKER_DELAY)
	fmt.Println("CUSTOMER DELAY:", config.AVERAGE_CUSTOMER_DELAY)
	fmt.Println("TASKLIST SIZE:", config.TASKLIST_SIZE)
	fmt.Println("WAREHOUSE SIZE:", config.WAREHOUSE_SIZE)
}
