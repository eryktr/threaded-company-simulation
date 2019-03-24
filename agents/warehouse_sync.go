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
		case delivery := <-WarehouseWrite:
			if len(Warehouse) < Capacity {
				Warehouse = append(Warehouse, delivery.product)
				delivery.Success <- true
			} else {
				delivery.Success <- false
			}
		case visit := <-WarehouseRead:
			if len(Warehouse) > 0 {

			}
		}

	}
}

func popRandomProduct() int {
	i := rand.Intn(len(Warehouse))
	prod := Warehouse[i]
	Warehouse = append(Warehouse[:i], Warehouse[i+1:]...)
	return prod
}
