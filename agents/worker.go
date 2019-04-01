package agents

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	Id             int
	TaskList       chan TaskListReadOperation
	Warehouse      chan WarehouseWriteOperation
	Logger         chan string
	MulltMachines  []MultiplicationMachine
	AddMachines    []AdditionMachine
	CompletedTasks int
	IsPatient      bool
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

	if worker.IsPatient {
		if job.Operation == PLUS {
			machine := worker.randomAdditionMachine()
			result := make(chan Job)
			request := MachineWriteOp{job, result}
			worker.Logger <- fmt.Sprintf("WORKER %d: WAITING FOR ADDITION MACHINE %d\n", worker.Id, machine.Id)
			machine.Input <- request
			res := <-result
			worker.IncreaseCompletedTasks()
			return res.Result
		} else if job.Operation == TIMES {
			machine := worker.randomMultiplicationMachine()
			result := make(chan Job)
			request := MachineWriteOp{job, result}
			worker.Logger <- fmt.Sprintf("WORKER %d: WAITING FOR MULTIPLICATION MACHINE %d\n", worker.Id, machine.Id)
			machine.Input <- request
			res := <-result
			worker.IncreaseCompletedTasks()
			return res.Result
		}
	} else {
		if job.Operation == PLUS {
			machine := worker.randomAdditionMachine()
			result := make(chan Job)
			request := MachineWriteOp{job, result}
			working := false
			worker.Logger <- fmt.Sprintf("WORKER %d: WAITING FOR ADDITION MACHINE %d\n", worker.Id, machine.Id)
			for !working {
				select {
				case machine.Input <- request:
					working = true
					res := <-result
					worker.IncreaseCompletedTasks()
					return res.Result

				default:
					time.Sleep(1 * time.Second)
					worker.Logger <- fmt.Sprintf("WORKER %d:  ADDITION MACHINE %d IS BUSY. SWITCHING\n", worker.Id, machine.Id)
					machine = worker.randomAdditionMachine()
				}

			}
		} else if job.Operation == TIMES {
			machine := worker.randomMultiplicationMachine()
			result := make(chan Job)
			request := MachineWriteOp{job, result}
			working := false
			worker.Logger <- fmt.Sprintf("WORKER %d: WAITING FOR MULTIPLICATION MACHINE %d\n", worker.Id, machine.Id)
			for !working {
				select {
				case machine.Input <- request:
					working = true
					res := <-result
					worker.IncreaseCompletedTasks()
					return res.Result

				default:
					time.Sleep(1 * time.Second)
					worker.Logger <- fmt.Sprintf("WORKER %d:  MULTIPLICATION MACHINE %d IS BUSY. SWITCHING\n", worker.Id, machine.Id)
					machine = worker.randomMultiplicationMachine()
				}
			}
		}
	}

	return 0
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

func (worker *Worker) IncreaseCompletedTasks() {
	worker.CompletedTasks += 1
}

func (worker *Worker) StoreProduct(product int) {
	success := make(chan bool, 1)
	operation := WarehouseWriteOperation{product, success}
	worker.Warehouse <- operation
	worker.Logger <- fmt.Sprintf("Worker %d: PRODUCT %d STORED IN THE WAREHOUSE\n", worker.Id, product)
}

func (worker *Worker) randomMultiplicationMachine() MultiplicationMachine {
	i := rand.Intn(len(worker.MulltMachines))
	return worker.MulltMachines[i]
}

func (worker *Worker) randomAdditionMachine() AdditionMachine {
	i := rand.Intn(len(worker.AddMachines))
	return worker.AddMachines[i]
}
