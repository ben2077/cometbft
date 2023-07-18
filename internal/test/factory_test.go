package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ben2077/cometbft/types"
)

func TestMakeHeader(t *testing.T) {
	header := MakeHeader(t, &types.Header{})
	require.NotNil(t, header)

	require.NoError(t, header.ValidateBasic())
}
