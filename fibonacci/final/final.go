package final

import "errors"

type fibonacciChannel chan int

// sender
func newFibonacciChannel(limit int) fibonacciChannel {
	newFibChan := make(fibonacciChannel)
	first, second := 0, 1

	go func() {
		for {
			if limit <= 0 {
				// close(newFibChan)
				break
			}
			newFibChan <- first
			first, second = second, first+second
			limit--
		}
	}()

	return newFibChan
}

// receiver
func (c fibonacciChannel) Next() (int, error) {
	nextNum, ok := <-c

	if !ok {
		return -1, errors.New("fibonacci channel empty")
	}

	return nextNum, nil
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
