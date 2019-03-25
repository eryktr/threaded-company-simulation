package agents

import (
	"fmt"

	"github.com/projects/threaded-company-simulation/config"
)

var MaxSize = config.TASKLIST_SIZE
var TaskList []Job = make([]Job, 0, MaxSize)

func SynchronizeTaskList() {
	for {
		select {
		case write := <-maybeWrite(len(TaskList) < MaxSize, TaskListWrite):
			TaskList = append(TaskList, write.Task)
			write.Success <- true

		case read := <-maybeRead(len(TaskList) > 0, TaskListRead):
			task := TaskList[0]
			TaskList = TaskList[1:]
			read.Task <- task
			read.Success <- true
		}
	}
}

func maybeRead(expression bool, c chan TaskListReadOperation) chan TaskListReadOperation {
	if expression {
		return c
	} else {
		return nil
	}
}

func maybeWrite(expression bool, c chan TaskListWriteOperation) chan TaskListWriteOperation {
	if expression {
		return c
	} else {
		return nil
	}
}

func PrintTasklist() {
	for _, element := range TaskList {
		fmt.Println("TASK", element.ToString())
	}
}
