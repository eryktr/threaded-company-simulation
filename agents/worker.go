package agents

import (
	"fmt"
	"time"
)

type Worker struct {
	Id        int
	TaskList  chan TaskListReadOperation
	Warehouse chan WarehouseWriteOperation
	Logger    chan string
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
	response := make(chan Job, 1)
	success := make(chan bool, 1)
	request := TaskListReadOperation{response, success}
	worker.TaskList <- request
	job := <-response
	worker.Logger <- fmt.Sprintf("Worker %d: FETCHED %s\n", worker.Id, job.ToString())
	return job
}

func (worker *Worker) Sleep() {
	time.Sleep(RandomSleepDuration(PT_WORKER) * time.Second)
}

func (worker *Worker) StoreProduct(product int) {

	success := make(chan bool, 1)
	operation := WarehouseWriteOperation{product, success}
	worker.Warehouse <- operation
	worker.Logger <- fmt.Sprintf("Worker %d: PRODUCT %d STORED IN THE WAREHOUSE\n", worker.Id, product)
}
