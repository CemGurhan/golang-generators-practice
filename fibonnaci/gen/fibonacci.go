package gen

type fibonacciGenerator struct {
	first, second, limit int
}

func (f *fibonacciGenerator) Next() int {
	next := f.first

	f.first, f.second = f.second, f.first+f.second
	f.limit--

	return next
}

func newFibonacciGenerator(seqIndex int) fibonacciGenerator {
	return fibonacciGenerator{first: 0, second: 1, limit: seqIndex}
}

func (f *fibonacciGenerator) genFibonacciSequence() []int {
	res := []int{}

	for f.limit != 0 {
		res = append(res, f.Next())
	}
	return res
}
