package agents

import (
	"fmt"
	"time"
)

type MachineWriteOp struct {
	InputJob Job
	Result   chan Job
}

type AdditionMachine struct {
	Id     int
	Input  chan MachineWriteOp
	Logger chan string
}

type MultiplicationMachine struct {
	Id     int
	Input  chan MachineWriteOp
	Logger chan string
}

func (am *AdditionMachine) SolveAddition(job Job, result chan Job) {
	am.Logger <- fmt.Sprintf("Addition Machine %d: Executing Job %s", am.Id, job.ToString())
	time.Sleep(RandomSleepDuration(PT_MACHINE))
	job.Result = job.First + job.Second
	am.Logger <- fmt.Sprintf("Addition Machine %d: Job %s = %d Solved!", am.Id, job.ToString(), job.Result)
	result <- job
}

func (mm *MultiplicationMachine) SolveMultiplication(job Job, result chan Job) {
	mm.Logger <- fmt.Sprintf("Multiplication Machine %d: Executing Job %s", mm.Id, job.ToString())
	time.Sleep(RandomSleepDuration(PT_MACHINE))
	job.Result = job.First * job.Second
	mm.Logger <- fmt.Sprintf("Multiplication Machine %d: Job %s = %d Solved!", mm.Id, job.ToString(), job.Result)
	result <- job
}

func (am *AdditionMachine) RunAdditionMachine() {
	for {
		select {
		case task := <-am.Input:
			am.SolveAddition(task.InputJob, task.Result)
		}
	}
}

func (mm *MultiplicationMachine) RunMultiplicationMachine() {
	for {
		select {
		case task := <-mm.Input:
			mm.SolveMultiplication(task.InputJob, task.Result)
		}
	}
}
