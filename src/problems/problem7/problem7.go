package problem7

import "problems"

func isPrime(factors []int, number int) bool {
	for _, fact := range factors{
		if number % fact == 0 {
			return false
		}
	}
	return true
}

func findPrime(primeNr int) int {
	factors := make([]int, 0)
	length := 0
	for start := 2; length < primeNr; start += 1{
		if isPrime(factors, start) {
			factors = append(factors, start)
			length += 1
		}
	}
	return factors[primeNr - 1]
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{}) int {
		return findPrime(args[0].(int))
	}}
}
