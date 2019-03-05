package factory

import (
	"time"
)

func customer_sleep() {
	time.Sleep(randomSleepDuration(PT_CUSTOMER) * time.Second)
}

func fetch_product_from_warehouse(warehouse chan int) {
	product := <-warehouse
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
