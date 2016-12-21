package project

type Project struct {
	Calculate func(args...interface{}) int
}
