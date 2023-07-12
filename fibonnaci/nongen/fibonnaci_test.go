package nongen

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	expected := 8
	actual := fibonacci(6)
	require.Equal(t, expected, actual)
}
