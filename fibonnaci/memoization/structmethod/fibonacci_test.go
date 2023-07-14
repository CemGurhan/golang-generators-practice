package structmethod

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci_Memoized(t *testing.T) {
	memoized := memoize(fibonacci)
	actual := memoized.Call(6)
	require.Equal(t, 8, actual)
}
