package factory

import (
	"time"

	"github.com/projects/threaded-company-simulation/config"
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

func fetch_job(list chan Job) Job {
	job := <-list
	print_job(job, "assigned")
	return job
}

func worker_sleep() {
	time.Sleep(randomSleepDuration(PT_WORKER) * time.Second)
}

func store_in_warehouse(product int, warehouse chan int) {
	max_warehouse_size := config.WAREHOUSE_SIZE
	for {
		if len(warehouse) < max_warehouse_size {
			warehouse_mutex.Lock()
			if len(warehouse) >= max_warehouse_size {
				warehouse_mutex.Unlock()
				continue
			}
			warehouse <- product
			print_product(product, "stored")
			warehouse_mutex.Unlock()
			break
		} else {
			sleep_failure()
		}
	}
}
func Worker(list chan Job, warehouse chan int) {

	for {
		if len(list) > 0 {
			lock_list()
			if len(list) == 0 {
				unlock_list()
				continue
			}
			job := fetch_job(list)
			unlock_list()
			worker_sleep()
			product := createProduct(job)
			store_in_warehouse(product, warehouse)
		} else {
			sleep_failure()
		}
	}
}
