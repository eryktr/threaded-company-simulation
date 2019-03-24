package agents

import (
	"github.com/projects/threaded-company-simulation/config"
)

var MaxSize = config.TASKLIST_SIZE
var TaskList []Job = make([]Job, 0, MaxSize)

func SynchronizeTaskList() {
	for {
		select {
		case write := <-TaskListWrite:
			if len(TaskList) < MaxSize {
				TaskList = append(TaskList, write.Task)
				write.Success <- true
			} else {
				write.Success <- false
			}
		case read := <-TaskListRead:
			if len(TaskList) > 0 {
				task := TaskList[0]
				TaskList = TaskList[1:]
				read.Task <- task
				read.Success <- true
			} else {
				read.Success <- false
			}
		}
	}
}
