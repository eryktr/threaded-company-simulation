package factory

import (
	"fmt"
	"math/rand"
	"time"
)

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

func Ceo(sync chan Job) {
	for i := 0; i < 10; i++ {
		a, b := randomArguments()
		c := randomOperator()
		job := Job{a, b, c}
		sync <- job
		fmt.Println("Task", job.first, job.operation, job.second, "added to the list.")
		time.Sleep()
	}
}
