package factory

import (
	"fmt"
	"time"
)

func Customer(warehouse chan int) {
	for {
		if len(warehouse) > 0 {
			warehouse_mutex.Lock()
			if len(warehouse) <= 0 {
				warehouse_mutex.Unlock()
				continue
			}
			product := <-warehouse
			fmt.Println("Product", product, "collected by a customer")
			warehouse_mutex.Unlock()
			time.Sleep(randomSleepDuration(PT_CUSTOMER) * time.Second)
		}
	}
}
