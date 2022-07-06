package urlshort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFor(t *testing.T) {
	a := For()
	b := "astana"
	require.Equal(t, a, b)
}
