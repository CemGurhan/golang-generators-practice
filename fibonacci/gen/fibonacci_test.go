package gen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenFibonacciSequence_WithNewLimitedFibonacciGenerator_GeneratesValidFibonacciSequence(t *testing.T) {
	fibonacciGenerator := newFibonacciGenerator(5)
	actual := fibonacciGenerator.genFibonacciSequence()
	require.Equal(t, []int{0, 1, 1, 2, 3}, actual)
}
