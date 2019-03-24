package agents

import (
	"math/rand"
	"time"
)

type Customer struct {
	id        int
	Warehouse WarehouseReadOperation
}

func (customer *Customer) start(warehouse chan int) {
	// for {
	// 	if len(warehouse) > 0 {
	// 		if len(warehouse) <= 0 {
	// 			continue
	// 		}
	// 		fetch_product_from_warehouse(warehouse)
	// 	} else {
	// 		sleep_failure()
	// 	}
	// }
}

func (customer *Customer) sleep() {
	time.Sleep(RandomSleepDuration(PT_CUSTOMER) * time.Second)
}

func fetch_random_product(warehouse chan int) int {
	index := rand.Intn(len(warehouse))
	for i := 0; i < index; i++ {
		tmp := <-warehouse
		warehouse <- tmp
	}
	product := <-warehouse
	return product
}

func fetch_product_from_warehouse(warehouse chan int) {
	product := fetch_random_product(warehouse)
	Print_product_collected(product)
}
