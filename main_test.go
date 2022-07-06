package urlshort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFor(t *testing.T) {
	a := For()
	b := "astan"
	require.Equal(t, a, b)
}
func TestFor2(t *testing.T) {
	a := For()
	b := "astan"
	require.Equal(t, a, b)
}
