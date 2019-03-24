package agents

import (
	"time"
)

type Worker struct {
	id        int
	TaskList  TaskListReadOperation
	Warehouse WarehouseWriteOperation
}

func (worker *Worker) run() {
	for {
		// if len(list) > 0 {
		// 	lock_list()
		// 	if len(list) == 0 {
		// 		unlock_list()
		// 		continue
		// 	}
		// 	job := fetch_job(list)
		// 	unlock_list()
		// 	worker_sleep()
		// 	product := createProduct(job)
		// 	store_in_warehouse(product, warehouse)
		// } else {
		// 	sleep_failure()
		// }
	}
}
func (worker *Worker) createProduct(job Job) int {
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

func (worker *Worker) fetch_job(list chan Job) Job {
	job := <-list
	Print_job_fetched(job)
	return job
}

func (worker *Worker) sleep() {
	time.Sleep(RandomSleepDuration(PT_WORKER) * time.Second)
}

func (worker *Worker) store_product(product int, warehouse chan int) {
	// max_warehouse_size := config.WAREHOUSE_SIZE
	// for {
	// 	if len(warehouse) < max_warehouse_size {
	// 		warehouse_mutex.Lock()
	// 		if len(warehouse) >= max_warehouse_size {
	// 			warehouse_mutex.Unlock()
	// 			continue
	// 		}
	// 		warehouse <- product
	// 		print_product_added(product)
	// 		warehouse_mutex.Unlock()
	// 		break
	// 	} else {
	// 		sleep_failure()
	// 	}
	// }
}
