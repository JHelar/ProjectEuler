package problem6

import "problems"

func sumSquare(numberRange int, sum chan int) {
	mySum := 0
	for i := 1; i <= numberRange; i++{
		mySum += i
	}
	sum <- (mySum * mySum)
}

func squareSum(numberRange int, sum chan int) {
	mySum := 0
	for i := 1; i <= numberRange; i++{
		mySum += i * i
	}
	sum <- mySum
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{}) int {
		numberRange := args[0].(int)

		sSquare := make(chan int)
		sSum := make(chan int)

		go sumSquare(numberRange, sSquare)
		go squareSum(numberRange, sSum)

		return (<-sSquare - <-sSum)
	}}
}
