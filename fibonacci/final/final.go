package final

import "errors"

type FibonacciChannel chan int

type fibCache map[int]int

func NewFibonacciChannel(limit int) FibonacciChannel {
	newFibChan := make(FibonacciChannel)
	newFibCache := fibCache{}

	go func() {
		for i := 0; i < limit; i++ {
			newFibChan <- findFibValue(i, newFibCache)
		}
		close(newFibChan)
	}()

	return newFibChan
}

func findFibValue(seqIndex int, fibCache fibCache) int {
	if seqIndex < 2 {
		return seqIndex
	}

	if cachedFibVal, isCached := fibCache[seqIndex]; isCached {
		return cachedFibVal
	}

	fibValue := findFibValue(seqIndex-1, fibCache) +
		findFibValue(seqIndex-2, fibCache)
	fibCache[seqIndex] = fibValue

	return fibValue
}

func (c FibonacciChannel) Next() (int, error) {
	nextVal, ok := <-c
	if !ok {
		return -1, errors.New("fibonacci channel closed")
	}

	return nextVal, nil
}

// make sure to close the channel in `newFibonacciChannel`
// otherwise once the channel is emptied, if you have a receiver
// try to receive from it, it will hang, as there is nothing in the channel,
// but it is expecting something new to come due to the channel being open.

// fill the channel in the goroutine in `newFibonacciChannel`. If you don't
// have the goroutine here, and instead have the for loop, go will assume
// there is never a receiver present to receive the item you are trying to
// pass to the channel with `newFibChan <- first`. If you have the goroutine,
// go assumes that you are able to receive from the channel whilst the
// goroutine fills it, leading to no hang. Try removing the next method and
// running `newFibonacciChannel`, even without an explicit reader, we can still
// fill the channel as the filling is done in the goroutine.

// seqIndex is zero indexed. E.g if seqIndex is 5,
// we'll have 5 recursive paths + the base path of
// 0
//      5        1
//    4   3      2
//  3  2  2 1    3
// 2 11 01 0     4
//1 0            5

// 0 1 1 2 3 5
