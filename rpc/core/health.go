package core

import (
	ctypes "github.com/ben2077/cometbft/rpc/core/types"
	rpctypes "github.com/ben2077/cometbft/rpc/jsonrpc/types"
)

// Health gets node health. Returns empty result (200 OK) on success, no
// response - in case of an error.
// More: https://docs.cometbft.com/main/rpc/#/Info/health
func (env *Environment) Health(*rpctypes.Context) (*ctypes.ResultHealth, error) {
	return &ctypes.ResultHealth{}, nil
}
