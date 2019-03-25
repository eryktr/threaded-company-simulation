package agents

import (
	"fmt"
	"time"
)

type Customer struct {
	Id        int
	Warehouse chan WarehouseReadOperation
}

func (customer *Customer) Run() {
	for {
		product := make(chan int, 1)
		success := make(chan bool, 1)
		request := WarehouseReadOperation{product, success}
		customer.Warehouse <- request
		fmt.Printf("Customer %d: PRODUCT %d PICKED FROM THE WAREHOUSE\n", customer.Id, <-product)
		customer.Sleep()
	}
}

func (customer *Customer) Sleep() {
	time.Sleep(RandomSleepDuration(PT_CUSTOMER) * time.Second)
}
