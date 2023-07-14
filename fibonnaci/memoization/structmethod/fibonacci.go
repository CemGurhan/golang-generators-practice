package structmethod

type Memoized struct {
	Procedure func(int) int
	Cache     map[int]int
}

func (m *Memoized) Call(val int) int {
	if result, isResultCached := m.Cache[val]; isResultCached {
		return result
	}

	result := m.Procedure(val)
	m.Cache[val] = result
	return result
}

func memoize(procedure func(int) int) *Memoized {
	return &Memoized{Procedure: procedure, Cache: map[int]int{}}
}

func fibonacci(seqIndex int) int {
	if seqIndex <= 1 {
		return seqIndex
	}

	return fibonacci(seqIndex-1) + fibonacci(seqIndex-2)
}
