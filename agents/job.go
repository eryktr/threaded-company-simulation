package agents

import "fmt"

type Operator int

const (
	PLUS  Operator = 1
	TIMES Operator = 2
)

type Job struct {
	First, Second int
	Operation     Operator
	Result        int
}

func (job *Job) ToString() string {
	operation := Operation_to_ascii(job.Operation)
	return fmt.Sprintf("%d %s %d", job.First, operation, job.Second)
}

func Operation_to_ascii(operation Operator) string {
	switch operation {
	case PLUS:
		return "+"
	default:
		return "*"
	}
}
