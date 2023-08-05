package final

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFibonacciChannel(t *testing.T) {
	seq := []int{}
	fibChan := NewFibonacciChannel(5)
	for i := 0; i < 5; i++ {
		nextNum, err := fibChan.Next()
		require.NoError(t, err)
		seq = append(seq, nextNum)
	}

	require.Equal(t, []int{0, 1, 1, 2, 3}, seq)
}

func TestNext_ClosedChannelReturnsError(t *testing.T) {
	fibChan := NewFibonacciChannel(1)
	_, err := fibChan.Next()
	require.NoError(t, err)

	actual, err := fibChan.Next()

	require.ErrorContains(t, err, "fibonacci channel closed")
	require.Equal(t, -1, actual)
}

func TestFindFibValue(t *testing.T) {
	testCases := map[string]struct {
		seqIndex int
		expected int
		fibCache fibCache
	}{
		"seqIndex 0": {
			seqIndex: 0,
			expected: 0,
			fibCache: fibCache{},
		},
		"seqIndex 1": {
			seqIndex: 1,
			expected: 1,
			fibCache: fibCache{},
		},
		"seqIndex 2": {
			seqIndex: 2,
			expected: 1,
			fibCache: fibCache{},
		},
		"seqIndex 3": {
			seqIndex: 3,
			expected: 2,
			fibCache: fibCache{},
		},
		"seqIndex 4": {
			seqIndex: 4,
			expected: 3,
			fibCache: fibCache{},
		},
		"seqIndex 5": {
			seqIndex: 5,
			expected: 5,
			fibCache: fibCache{},
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := findFibValue(test.seqIndex, test.fibCache)

			require.Equal(t, test.expected, actual)
		})
	}
}
