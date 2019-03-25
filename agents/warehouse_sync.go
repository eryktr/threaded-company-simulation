package agents

import (
	"math/rand"

	"github.com/projects/threaded-company-simulation/config"
)

var Capacity = config.WAREHOUSE_SIZE
var Warehouse []int = make([]int, 0, Capacity)

func SynchronizeWarehouse() {
	for {
		select {
		case delivery := <-maybeStore(len(Warehouse) < Capacity, WarehouseWrite):
			Warehouse = append(Warehouse, delivery.product)
			delivery.Success <- true

		case visit := <-maybeCollect(len(Warehouse) > 0, WarehouseRead):
			prod := popRandomProduct()
			visit.product <- prod
			visit.Success <- true
			break
		}

	}
}

func popRandomProduct() int {
	i := rand.Intn(len(Warehouse))
	prod := Warehouse[i]
	Warehouse = append(Warehouse[:i], Warehouse[i+1:]...)
	return prod
}

func maybeStore(expression bool, c chan WarehouseWriteOperation) chan WarehouseWriteOperation {
	if expression {
		return c
	} else {
		return nil
	}
}

func maybeCollect(expression bool, c chan WarehouseReadOperation) chan WarehouseReadOperation {
	if expression {
		return c
	} else {
		return nil
	}
}
