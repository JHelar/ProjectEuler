package problem4

import (
	"strconv"
	"sync"
	"problems"
)

func strReverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func findProdPalindrome(check, min , max int, biggest chan int, wg *sync.WaitGroup){
	defer wg.Done()
	myBiggest := 0
	for checkMe := check; checkMe < check + min; checkMe ++ {
		for i := min; i <= max; i += 1 {
			sum := i * checkMe
			if sumStr := strconv.Itoa(sum); myBiggest < sum {
				reverseSum, _ := strconv.Atoi(strReverse(sumStr))
				if reverseSum == sum {

					myBiggest = sum
				}
			}
		}
	}
	biggest <- myBiggest
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{})int {
		min := args[0].(int)
		max := args[1].(int)

		biggest := make(chan int)
		var wg sync.WaitGroup

		theBiggest := 0
		//Create routines
		go func(){
			for{
				select {
				case i := <- biggest:
					if theBiggest < i {
						theBiggest = i
					}
				}
			}
		}()
		for i := min; i < max; i += min * 2 {
			wg.Add(1)
			go func(myCheck int){
				findProdPalindrome(myCheck, min, max, biggest, &wg)
			}(i)
		}
		wg.Wait()
		return theBiggest
	}}
}
