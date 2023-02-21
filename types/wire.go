package types

import (
	ntypes "github.com/Sotatek-huytran2/go-sdk/common/types"
	"github.com/Sotatek-huytran2/go-sdk/types/tx"
	"github.com/tendermint/go-amino"
	types "github.com/tendermint/tendermint/rpc/core/types"
)

func NewCodec() *amino.Codec {
	cdc := amino.NewCodec()
	types.RegisterAmino(cdc)
	ntypes.RegisterWire(cdc)
	tx.RegisterCodec(cdc)
	return cdc
}
