package agents

import (
	"fmt"
	"time"
)

type Machine interface {
}

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
	am.Logger <- fmt.Sprintf("ADDITION MACHINE %d: EXECUTING JOB %s\n", am.Id, job.ToString())
	time.Sleep(RandomSleepDuration(PT_MACHINE) * time.Second)
	job.Result = job.First + job.Second
	am.Logger <- fmt.Sprintf("ADDITION MACHINE %d: Job %s = %d SOLVED\n", am.Id, job.ToString(), job.Result)
	result <- job
}

func (mm *MultiplicationMachine) SolveMultiplication(job Job, result chan Job) {
	mm.Logger <- fmt.Sprintf("MULTIPLICATION MACHINE %d: EXECUTING JOB %s\n", mm.Id, job.ToString())
	time.Sleep(RandomSleepDuration(PT_MACHINE) * time.Second)
	job.Result = job.First * job.Second
	mm.Logger <- fmt.Sprintf("MULTIPLICATION MACHINE %d: JOB %s = %d SOLVED!\n", mm.Id, job.ToString(), job.Result)
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
