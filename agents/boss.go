package agents

import (
	"fmt"
	"math/rand"
	"time"
)

type Boss struct {
	Id       int
	TaskList chan TaskListWriteOperation
	Logger   chan string
}

func (boss *Boss) Run() {
	for {
		job := boss.GenerateRandomJob()
		responseChannel := make(chan bool, 1)
		writeOperation := TaskListWriteOperation{job, responseChannel}
		boss.TaskList <- writeOperation
		time.Sleep(100 * time.Millisecond)
		boss.Logger <- fmt.Sprintln("BOSS: Task ", job.ToString(), "added to the list.")
		boss.Sleep()
	}
}

func (boss *Boss) GenerateRandomJob() Job {
	a, b := randomArguments()
	c := randomOperator()
	return Job{a, b, c}
}

func (boss *Boss) Sleep() {
	time.Sleep(RandomSleepDuration(PT_CEO) * time.Second)
}

func randomArguments() (int, int) {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	arg1, arg2 := r.Intn(1000), r.Intn(1000)
	return arg1, arg2
}

func randomOperator() Operator {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	op := r.Intn(3) + 1
	return Operator(op)
}
