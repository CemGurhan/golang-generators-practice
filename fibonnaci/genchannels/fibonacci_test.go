package genchannels

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	actual := []int{}
	for val := range fibonacci(5) {
		actual = append(actual, val)
	}

	require.Equal(t, []int{0, 1, 1, 2, 3}, actual)
}

func TestFibonacciNext(t *testing.T) {
	actual := []int{}
	fibChan := fibonacci(5)
	actual = append(actual, *fibChan.Next())
	actual = append(actual, *fibChan.Next())

	require.Equal(t, []int{0, 1}, actual)
}
