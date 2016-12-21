package project2

import (
	"projects"
)

func fib(one, another, max int) int {
	if one <= max {
		if one % 2 == 0 {
			return one + fib(another, one + another, max)
		}else{
			return fib(another, one + another, max)
		}
	}
	if one % 2 == 0{
		return one
	}else{
		return 0
	}
}

func New() *project.Project{
	return &project.Project{Calculate:func(args...interface{})int {
		return fib(1, 2, args[0].(int))
	}}
}
