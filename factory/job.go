package factory

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
