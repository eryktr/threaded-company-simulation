package agents

import "fmt"

type Operator int

const (
	PLUS  Operator = 1
	MINUS Operator = 2
	TIMES Operator = 3
)

type Job struct {
	first, second int
	operation     Operator
}

func (job *Job) ToString() string {
	operation := Operation_to_ascii(job.operation)
	return fmt.Sprintf("%d %s %d", job.first, operation, job.second)
}

func Operation_to_ascii(operation Operator) string {
	switch operation {
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	default:
		return "*"
	}
}
