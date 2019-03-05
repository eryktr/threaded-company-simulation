package factory

import (
	"math/rand"
	"time"
)

func customer_sleep() {
	time.Sleep(randomSleepDuration(PT_CUSTOMER) * time.Second)
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
	print_product(product, "collected")
}

func Customer(warehouse chan int) {
	for {
		if len(warehouse) > 0 {
			lock_warehouse()
			if len(warehouse) <= 0 {
				unlock_warehouse()
				continue
			}
			fetch_product_from_warehouse(warehouse)
			unlock_warehouse()
			customer_sleep()
		} else {
			sleep_failure()
		}
	}
}
