package factory

import (
	"github.com/projects/threaded-company-simulation/config"
)

var List = make(chan Job, config.TASKLIST_SIZE)
var Warehouse = make(chan int, config.WAREHOUSE_SIZE)
