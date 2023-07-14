package mapmethod

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	actual := fibonacci(4)

	require.Equal(t, 3, actual)

	actual = fibonacci(5)

	require.Equal(t, 52, actual)
}
