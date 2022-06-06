package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings(t *testing.T) {
	t.Run("ConvertArrayOptsToString", func(t *testing.T) {
		expected := []string{"0x0000000000000000000000000000000000000001,0x0000000000000000000000000000000000000002"}

		assert.Equal(t, expected, ConvertArrayOptsToApiParam([]string{
			"0x0000000000000000000000000000000000000001",
			"0x0000000000000000000000000000000000000002",
		}))
	})
}
