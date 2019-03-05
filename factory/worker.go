package factory

import (
	"fmt"
	"time"
)

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
		if len(list) > 0 {
			list_mutex.Lock()
			if len(list) == 0 {
				list_mutex.Unlock()
				continue
			}
			job := <-list
			fmt.Println("Job", job.first, job.operation, job.second, "assigned to a worker")
			list_mutex.Unlock()
			time.Sleep(randomSleepDuration(PT_WORKER) * time.Second)
			product := createProduct(job)
			for {
				if len(warehouse) < 30 {
					warehouse_mutex.Lock()
					if len(warehouse) >= 30 {
						warehouse_mutex.Unlock()
						continue
					}
					warehouse <- product
					fmt.Println("Product", product, "Stored in the warehouse.")
					warehouse_mutex.Unlock()
					break
				}
			}
		}

	}
}
