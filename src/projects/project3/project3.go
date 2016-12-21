package project3

import (
	"projects"
)

func primeFactors(n int) []int{
	factors := make([]int, 0)
	d := 2
	for n > 1 {
		for n % d == 0 {
			factors = append(factors, d)
			n /= d
		}
		d += 1
	}
	return factors
}

func New() *project.Project{
	return &project.Project{Calculate:func(args...interface{})int {
		biggest := 0
		for _, v := range primeFactors(args[0].(int)) {
			if v > biggest{
				biggest = v
			}
		}
		return biggest
	}}
}
