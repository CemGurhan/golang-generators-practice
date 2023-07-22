package nongen

func fibonacci(endSeqNumber int) int {
	if endSeqNumber < 2 {
		return endSeqNumber
	}
	return fibonacci(endSeqNumber-1) + fibonacci(endSeqNumber-2)
}
