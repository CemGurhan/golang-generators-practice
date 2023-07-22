package genchannels

import (
	"errors"
)

type fibonacciChannel chan int

// fills the fibonacciChannel up to the limiting index.
func fibonacci(limit int) fibonacciChannel {
	first, second := 0, 1
	channel := make(chan int)

	go func() {
		for {
			if limit == 0 {
				close(channel)
				return
			}
			channel <- first
			first, second = second, first+second
			limit--
		}
	}()

	return channel
}

// lazily evaluates the fibonacci channel contents every time invoked.
func (fibChan fibonacciChannel) Next() (*int, error) {
	next, ok := <-fibChan
	if !ok {
		return nil, errors.New("no more solutions")
	}

	return &next, nil
}
