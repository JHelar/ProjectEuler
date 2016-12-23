package problem10

import (
	"problems"
	"math"
)

//Implementation of the ESieve factorisation algorithm
func ESieve(upperLimit int) (factors []int, sum int){
	sieveBound := int((upperLimit - 1) / 2)
	upperSqrt := int(math.Sqrt(float64(upperLimit)) - 1) / 2
	factorBitMap := make(map[int]bool)
	factors = make([]int, 1)
	sum = 2
	factors[0] = 2
	for bit := 1; bit <= sieveBound + 1; bit++{
		factorBitMap[bit] = true
	}

	for i := 1; i <= upperSqrt; i++{
		if val, ok := factorBitMap[i]; ok && val {
			for j := i * 2 * (i + 1); j <= sieveBound; j += 2 * i + 1 {
				factorBitMap[j] = false
			}
		}
	}
	for i := 1; i <= sieveBound; i++ {
		if(factorBitMap[i]){
			fact := 2 * i + 1
			sum += fact
			factors = append(factors, fact)
		}
	}
	return
}

func New() *problem.Problem {
	return &problem.Problem{Calculate:func(args...interface{}) int {
		_, sum := ESieve(args[0].(int))
		return sum
	}}
}
