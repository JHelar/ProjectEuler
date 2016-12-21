package problem1

import (
	"problems"
)

func treeFive(start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	return sum
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{})int{
		theNumber := args[0].(int)
		return treeFive(1, theNumber - 1)
	}}
}

