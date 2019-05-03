package agents

type TaskListWriteOperation struct {
	Task    Job
	Success chan bool
}

type TaskListReadOperation struct {
	Task    chan Job
	Success chan bool
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
var LogChannel = make(chan string)
var ServiceReportWrite = make(chan ReportChanelWriteOp)
var ServiceReportRead = make(chan ReportChannelReadOp)
var ServiceFixWrite = make(chan FixReport)
