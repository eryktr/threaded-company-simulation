package agents

type TaskListWriteOperation struct {
	Task    Job
	Success chan bool
}

type TaskListReadOperation struct {
	JobChannel chan Job
	Success    chan bool
}

type WarehouseWriteOperation struct {
	product int
	Success chan bool
}

type WarehouseReadOperation struct {
	product chan int
	Success chan bool
}

var TaskListWrite = make(chan TaskListWriteOperation)
var TaskListRead = make(chan TaskListReadOperation)
var WarehouseWrite = make(chan WarehouseWriteOperation)
var WarehouseRead = make(chan WarehouseReadOperation)
