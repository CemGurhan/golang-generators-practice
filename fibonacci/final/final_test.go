package final

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibChan(t *testing.T) {
	seq := []int{}
	fibChan := newFibonacciChannel(5)
	for i := 0; i < 5; i++ {
		nextNum, err := fibChan.Next()
		require.NoError(t, err)
		seq = append(seq, nextNum)
	}

	require.Equal(t, []int{0, 1, 1, 2, 3}, seq)
}
