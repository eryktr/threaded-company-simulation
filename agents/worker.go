package agents

import (
	"fmt"
	"time"
)

type Worker struct {
	Id        int
	TaskList  chan TaskListReadOperation
	Warehouse chan WarehouseWriteOperation
}

func (worker *Worker) Run() {
	for {
		job := worker.FetchJob()
		worker.Sleep()
		product := worker.CreateProduct(job)
		worker.StoreProduct(product)
		worker.Sleep()
	}
}
func (worker *Worker) CreateProduct(job Job) int {
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

func (worker *Worker) FetchJob() Job {
	accepted := false
	response := make(chan Job, 1)
	for !accepted {
		success := make(chan bool, 1)
		request := TaskListReadOperation{response, success}
		worker.TaskList <- request
		accepted = <-success
		time.Sleep(100 * time.Millisecond)
	}
	job := <-response
	fmt.Printf("Worker %d: FETCHED %s\n", worker.Id, job.ToString())
	return job
}

func (worker *Worker) Sleep() {
	time.Sleep(RandomSleepDuration(PT_WORKER) * time.Second)
}

func (worker *Worker) StoreProduct(product int) {
	accepted := false

	for !accepted {
		success := make(chan bool, 1)
		operation := WarehouseWriteOperation{product, success}
		worker.Warehouse <- operation
		accepted = <-success
	}
	fmt.Printf("Worker %d: PRODUCT %d STORED IN THE WAREHOUSE\n", worker.Id, product)
}
