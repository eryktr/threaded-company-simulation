package factory

import (
	"math/rand"
	"time"

	"github.com/projects/threaded-company-simulation/config"
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

func randomJob() Job {
	a, b := randomArguments()
	c := randomOperator()
	return Job{a, b, c}
}

func addJobToList(j Job, l chan Job) {
	l <- j
	print_job(j, "added")
}

func sleep() {
	time.Sleep(randomSleepDuration(PT_CEO) * time.Second)
}

func Ceo(list chan Job) {
	max_list_size := config.TASKLIST_SIZE
	for {
		if len(list) < max_list_size {
			lock_list()
			if len(list) >= max_list_size {
				unlock_list()
				continue
			}
			job := randomJob()
			addJobToList(job, list)
			unlock_list()
			sleep()
		} else {
			sleep_failure()
		}
	}
}
