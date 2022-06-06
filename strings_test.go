package client

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrings(t *testing.T) {
	t.Run("ConvertArrayOptsToString", func(t *testing.T) {
		expected := "0x0000000000000000000000000000000000000001,0x0000000000000000000000000000000000000002"

		assert.Equal(t, expected, ConvertArrayOptsToApiParam("0x0000000000000000000000000000000000000001,0x0000000000000000000000000000000000000002"))
		assert.Equal(t, expected, ConvertArrayOptsToApiParam([]common.Address{
			common.HexToAddress("0x0000000000000000000000000000000000000001"),
			common.HexToAddress("0x0000000000000000000000000000000000000002"),
		}))
		assert.Equal(t, expected, ConvertArrayOptsToApiParam([]common.Address{
			common.HexToAddress("0x0000000000000000000000000000000000000001"),
			common.HexToAddress("0x0000000000000000000000000000000000000002"),
		}))
		assert.Equal(t, nil, ConvertArrayOptsToApiParam(nil))
	})
}
