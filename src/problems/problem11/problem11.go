package problem11

import (
	"problems"
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func prod(vals...int) (sum int){
	sum = 1
	for _,v := range vals {
		sum *= v
	}
	return
}

func diagonalBest(grid [][]int, cursorX, cursorY, rows, cols int, resultChan chan int) {
	result := 0
	//Check southeast
	s1 := cursorY
	s2 := cursorY + 1
	s3 := cursorY + 2
	s4 := cursorY + 3

	t1 := cursorX
	t2 := cursorX + 1
	t3 := cursorX + 2
	t4 := cursorX + 3

	if s4 < rows && s3 < rows && s2 < rows && s1 < rows && t4 < cols && t3 < cols && t2 < cols && t1 < cols {
		if p := prod(grid[s1][t1], grid[s2][t2], grid[s3][t3], grid[s4][t4]); p > result {
			result = p
		}
	}
	//Check southwest
	s1 = cursorY
	s2 = cursorY + 1
	s3 = cursorY + 2
	s4 = cursorY + 3

	t1 = cursorX
	t2 = cursorX - 1
	t3 = cursorX - 2
	t4 = cursorX - 3

	if s4 < rows && s3 < rows && s2 < rows && s1 < rows && t4 >= 0 && t3 >= 0 && t2 >= 0 && t1 >= 0 {
		if p := prod(grid[s1][t1], grid[s2][t2], grid[s3][t3], grid[s4][t4]); p > result {
			result = p
		}
	}
	//Check northeast
	s1 = cursorY
	s2 = cursorY - 1
	s3 = cursorY - 2
	s4 = cursorY - 3

	t1 = cursorX
	t2 = cursorX + 1
	t3 = cursorX + 2
	t4 = cursorX + 3

	if s4 >= 0 && s3 >= 0 && s2 >= 0 && s1 >= 0 && t4 < cols && t3 < cols && t2 < cols && t1 < cols {
		if p := prod(grid[s1][t1], grid[s2][t2], grid[s3][t3], grid[s4][t4]); p > result {
			result = p
		}
	}
	//Check northeast
	s1 = cursorY
	s2 = cursorY - 1
	s3 = cursorY - 2
	s4 = cursorY - 3

	t1 = cursorX
	t2 = cursorX - 1
	t3 = cursorX - 2
	t4 = cursorX - 3

	if s4 >= 0 && s3 >= 0 && s2 >= 0 && s1 >= 0 && t4 >= 0 && t3 >= 0 && t2 >= 0 && t1 >= 0 {
		if p := prod(grid[s1][t1], grid[s2][t2], grid[s3][t3], grid[s4][t4]); p > result {
			result = p
		}
	}
	resultChan <- result
}

func crossBest(grid [][]int, cursorX, cursorY, rows, cols int, resultChan chan int) {
	result := 0
	//Check south
	s1 := cursorY
	s2 := cursorY + 1
	s3 := cursorY + 2
	s4 := cursorY + 3
	if s4 < rows && s3 < rows && s2 < rows && s1 < rows{
		if p := prod(grid[s1][cursorX], grid[s2][cursorX], grid[s3][cursorX], grid[s4][cursorX]); p > result {
			result = p
		}
	}
	//Check north
	s1 = cursorY
	s2 = cursorY - 1
	s3 = cursorY - 2
	s4 = cursorY - 3
	if s4 >= 0 && s3 >= 0 && s2 >= 0 && s1 >= 0{
		if p := prod(grid[s1][cursorX], grid[s2][cursorX], grid[s3][cursorX], grid[s4][cursorX]); p > result {
			result = p
		}
	}
	//Check east
	s1 = cursorX
	s2 = cursorX + 1
	s3 = cursorX + 2
	s4 = cursorX + 3
	if s4 < cols && s3 < cols && s2 < cols && s1 < cols{
		if p := prod(grid[cursorY][s1], grid[cursorY][s2], grid[cursorY][s3], grid[cursorY][s4]); p > result {
			result = p
		}
	}
	//Check west
	s1 = cursorX
	s2 = cursorX - 1
	s3 = cursorX - 2
	s4 = cursorX - 3
	if s4 >= 0 && s3 >= 0 && s2 >= 0 && s1 >= 0 {
		if p := prod(grid[cursorY][s1], grid[cursorY][s2], grid[cursorY][s3], grid[cursorY][s4]); p > result {
			result = p
		}
	}
	resultChan <- result
}

func getGridProd(grid [][]int, rows, cols int) (sum int){
	sum = 0
	cursorX := 0
	cursorY := 0
	resultChan := make(chan int)

	go func() {
		for {
			select {
			case result := <- resultChan:
				if result > sum {
					sum = result
				}
			}
		}
	}()
	for cursorY = 0; cursorY < rows; cursorY++{
		for cursorX = 0; cursorX < cols; cursorX++{
			go crossBest(grid, cursorX, cursorY, rows, cols, resultChan)
			go diagonalBest(grid, cursorX, cursorY, rows, cols, resultChan)
		}
	}
	return
}

func New() *problem.Problem{
	return &problem.Problem{Calculate:func(args...interface{}) int {
		filePath := args[0].(string)
		grid := make([][]int, 0)
		cols := 0
		rows := 0

		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Problem %v", err)
			return -1
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			strNumbers := strings.Split(scanner.Text(), " ")
			row := make([]int, len(strNumbers))
			for i,c := range strNumbers {
				row[i],_ = strconv.Atoi(c)
			}
			grid = append(grid, row)
		}
		rows = len(grid)
		cols = len(grid[0])
		return getGridProd(grid, rows, cols)
	}}
}
