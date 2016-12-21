package problem8

import (
	"problems"
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func strProd(numberStr string) int {
	sum := 1
	for i,_ := range numberStr {
		val, _ := strconv.Atoi(string(numberStr[i]))
		sum *= val
	}
	return sum
}

func getXBest(numberStr string, x int) int {
	length := len(numberStr)
	best := 0
	prevRune := '0'
	for cursor := 0; cursor < length - x; cursor ++ {
		subStr := numberStr[cursor : cursor + x]
		if int(subStr[x - 1]) > int(prevRune) {
			if sum := strProd(subStr); sum > best {
				best = sum
			}
		}
		prevRune = rune(subStr[0])
	}
	return best
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{}) int {
		filePath := args[0].(string)
		amount := args[1].(int)

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Problem %v", err)
			return -1
		}
		defer file.Close()

		numberStr := ""
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			numberStr += scanner.Text()
		}

		return getXBest(numberStr, amount)
	}}
}
