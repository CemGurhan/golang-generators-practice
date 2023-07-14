package mapmethod

var cache = map[int]int{}

func fibonacci(seqIndex int) int {
	if seqIndex < 2 {
		return seqIndex
	}

	if cachedResult, isCached := cache[seqIndex]; isCached {
		return cachedResult
	}

	result := fibonacci(seqIndex-1) + fibonacci(seqIndex-2)
	cache[seqIndex] = result
	return result
}
