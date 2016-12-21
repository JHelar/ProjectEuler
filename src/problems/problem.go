package problem

type Problem struct {
	Calculate func(args...interface{}) int
}
