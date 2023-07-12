package nongen

func fibonacci(endSeqNumber int) int {
	if endSeqNumber == 1 || endSeqNumber == 0 {
		return endSeqNumber % 5
	}
	return fibonacci(endSeqNumber-1) + fibonacci(endSeqNumber-2)
}
