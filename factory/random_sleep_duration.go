package factory

import (
	"math/rand"
	"time"
)

type PersonType int

const (
	PT_CEO      PersonType = 1
	PT_WORKER   PersonType = 2
	PT_CUSTOMER PersonType = 3
)

func randomSleepDuration(pt PersonType) float64 {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	return
}
