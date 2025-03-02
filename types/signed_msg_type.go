package types

import cmtproto "github.com/ben2077/cometbft/proto/tendermint/types"

// IsVoteTypeValid returns true if t is a valid vote type.
func IsVoteTypeValid(t cmtproto.SignedMsgType) bool {
	switch t {
	case cmtproto.PrevoteType, cmtproto.PrecommitType:
		return true
	default:
		return false
	}
}
