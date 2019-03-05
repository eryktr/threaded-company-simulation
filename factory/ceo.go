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

func Ceo(list chan Job) {
	for {
		if len(list) < 30 {
			list_mutex.Lock()
			if len(list) >= 30 {
				list_mutex.Unlock()
				continue
			}
			a, b := randomArguments()
			c := randomOperator()
			job := Job{a, b, c}
			list <- job
			fmt.Println("Task", job.first, job.operation, job.second, "added to the list.")
			list_mutex.Unlock()
			time.Sleep(randomSleepDuration(PT_CEO) * time.Second)
		}
	}
}
