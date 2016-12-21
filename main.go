package main

import (
	"log"
	"problems/problem1"
	"problems/problem2"
	"problems/problem3"
	"problems/problem4"
	"problems/problem5"
	"problems/problem6"
	"problems/problem7"
	"problems/problem8"
)

func main(){
	prob1 := problem1.New()
	log.Printf("Problem1 result: %v", prob1.Calculate(1000))

	prob2 := problem2.New()
	log.Printf("Problem2 result: %v", prob2.Calculate(4000000))

	prob3 := problem3.New()
	log.Printf("Problem3 result: %v", prob3.Calculate(600851475143))

	prob4 := problem4.New()
	log.Printf("Problem4 result: %v", prob4.Calculate(100, 999))

	prob5 := problem5.New()
	log.Printf("Problem5 result: %v", prob5.Calculate(20))

	prob6 := problem6.New()
	log.Printf("Problem6 result: %v", prob6.Calculate(100))

	prob7 := problem7.New()
	log.Printf("Problem7 result: %v", prob7.Calculate(10001))

	prob8 := problem8.New()
	log.Printf("Problem8 result: %v", prob8.Calculate("./src/problems/problem8/numbers.txt", 13))
}