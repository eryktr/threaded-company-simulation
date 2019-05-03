package agents

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/projects/threaded-company-simulation/config"
)

type MachineWriteOp struct {
	InputJob Job
	Result   chan Job
}

type AdditionMachine struct {
	Id              int
	Input           chan MachineWriteOp
	Logger          chan string
	FixMe           chan bool
	IsBroken        bool
	BreakdownNumber int
}

type MultiplicationMachine struct {
	Id              int
	Input           chan MachineWriteOp
	Logger          chan string
	FixMe           chan bool
	IsBroken        bool
	BreakdownNumber int
}

func (am *AdditionMachine) BreakDown() {
	if am.IsBroken {
		return
	}
	am.IsBroken = true
	am.Logger <- fmt.Sprintf("ADDITION MACHINE %d HAS BROKEN DOWN \n", am.Id)
}

func (mm *MultiplicationMachine) BreakDown() {
	if mm.IsBroken {
		return
	}
	mm.IsBroken = true
	mm.Logger <- fmt.Sprintf("MULTIPLICATION MACHINE %d HAS BROKEN DOWN \n", mm.Id)
}

func (am *AdditionMachine) SolveAddition(job Job, result chan Job) {
	threshold := config.BREAKDOWN_PROBABILITY
	outcome := rand.Intn(100)
	if outcome > threshold {
		am.BreakDown()
	}
	if am.IsBroken {
		job.Result = 0
		am.Logger <- fmt.Sprintf("ADDITION MACHINE %d IS BROKEN: TASK %s NOT SOLVED\n", am.Id, job.ToString())
	} else {
		am.Logger <- fmt.Sprintf("ADDITION MACHINE %d: EXECUTING JOB %s\n", am.Id, job.ToString())
		time.Sleep(RandomSleepDuration(PT_MACHINE) * time.Second)
		job.Result = job.First + job.Second
		am.Logger <- fmt.Sprintf("ADDITION MACHINE %d: Job %s = %d SOLVED\n", am.Id, job.ToString(), job.Result)
	}
	result <- job
}

func (mm *MultiplicationMachine) SolveMultiplication(job Job, result chan Job) {
	threshold := config.BREAKDOWN_PROBABILITY
	outcome := rand.Intn(100)
	if outcome > threshold {
		mm.BreakDown()
	}
	if mm.IsBroken {
		job.Result = 0
		mm.Logger <- fmt.Sprintf("MULTIPLICATION MACHINE %d IS BROKEN: TASK %s NOT SOLVED", mm.Id, job.ToString())
	} else {
		mm.Logger <- fmt.Sprintf("MULTIPLICATION MACHINE %d: EXECUTING JOB %s\n", mm.Id, job.ToString())
		time.Sleep(RandomSleepDuration(PT_MACHINE) * time.Second)
		job.Result = job.First * job.Second
		mm.Logger <- fmt.Sprintf("MULTIPLICATION MACHINE %d: JOB %s = %d SOLVED!\n", mm.Id, job.ToString(), job.Result)
	}
	result <- job
}

func (am *AdditionMachine) RunAdditionMachine() {
	for {
		select {
		case task := <-am.Input:
			am.SolveAddition(task.InputJob, task.Result)
		case <-am.FixMe:
			am.Logger <- fmt.Sprintf("ADDITION MACHINE %d HAS BEEN FIXED\n", am.Id)
			am.IsBroken = true
			am.BreakdownNumber++
		}

	}
}

func (mm *MultiplicationMachine) RunMultiplicationMachine() {
	for {
		select {
		case task := <-mm.Input:
			mm.SolveMultiplication(task.InputJob, task.Result)
		case <-mm.FixMe:
			mm.Logger <- fmt.Sprintf("MULTIPLICATION MACHINE %d HAS BROKEN DOWN \n", mm.Id)
			mm.IsBroken = true
			mm.BreakdownNumber++
		}

	}
}
