package problem5

import "problems"

func findSmallestDivisible(numberRange int) int {
	count := 1;
	done := true
	for {
		done = true
		for div := 1; div <= numberRange; div += 1 {
			if count % div != 0{
				done = false
				break
			}
		}
		if done {
			break
		}else{
			count += 1
		}
	}
	return count
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{})int{
		return findSmallestDivisible(args[0].(int))
	}}
}
