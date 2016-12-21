package main

import (
	"projects/project1"
	"projects/project2"
	"projects/project3"
	"log"
)

func main(){
	proj1 := project1.New()
	proj2 := project2.New()
	proj3 := project3.New()


	log.Printf("Project1 result: %v", proj1.Calculate(1000))
	log.Printf("Project2 result: %v", proj2.Calculate(4000000))
	log.Printf("Project3 result: %v", proj3.Calculate(600851475143))
}