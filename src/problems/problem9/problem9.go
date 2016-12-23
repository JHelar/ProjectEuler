package problem9

import (
	"problems"
	"sync"
)

func getPythTriplet(n, m int) (int, int, int){
	a := (n * n) - (m * m)
	b := 2 * n * m
	c := (n * n) + (m * m)
	return a, b, c
}

func findPythTripletProd(start, end, value int, result chan int) {
	for m := start; m < end; m++ {
		for n := 1; n <= 100; n++ {
			if(m < n) {
				a, b, c := getPythTriplet(n, m)
				if sum := a + b + c; sum == value {
					result <- a * b * c
					return
				}
			}
		}
	}
	result <- -1
	return
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{}) int {
		value := args[0].(int)
		routines := args[1].(int)

		step := value / routines

		var wg sync.WaitGroup
		result := make(chan int)
		answer := -1


		go func(){
			for {
				select {
				case r := <- result:
					if r > -1 {
						answer = r
						break
					}
				}
			}
		}()

		for start := 0; start < value; start += step {
			wg.Add(1)
			go func(myStart, myEnd int){
				defer wg.Done()
				findPythTripletProd(myStart, myEnd, value, result)
			}(start, start + step)
		}
		wg.Wait()
		return answer
	}}
}

