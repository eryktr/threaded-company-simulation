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

func Ceo(sync chan Job) {
	for i := 0; i < 10; i++ {
		a, b := randomArguments()
		fmt.Println(a, "and", b)
	}
}
