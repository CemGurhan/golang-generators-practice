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

	for i := 0; i < 2; i++ {
		next, err := fibChan.Next()
		require.NoError(t, err)
		actual = append(actual, *next)
	}

	require.Equal(t, []int{0, 1}, actual)
}

func TestFibonacciNext_NoMoreSolutionsError(t *testing.T) {
	fibChan := fibonacci(5)
	var err error
	for err == nil {
		_, err = fibChan.Next()
	}

	require.ErrorContains(t, err, "no more solutions")
}
