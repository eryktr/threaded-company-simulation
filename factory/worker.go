package factory

import "fmt"

func createProduct(job Job) int {
	first := job.first
	second := job.second
	operation := job.operation
	switch operation {
	case PLUS:
		return first + second
	case MINUS:
		return first - second
	case TIMES:
		return first * second
	default:
		return 0
	}
}

func Worker(list chan Job, warehouse chan int) {
	for {
		job := <-list
		product := createProduct(job)
		warehouse <- product
		fmt.Println("Product", product, "Stored in the warehouse.")
	}
}
