package tx

import (
	"github.com/tendermint/go-amino"
	cryptoAmino "github.com/tendermint/tendermint/crypto/encoding/amino"

	"github.com/Sotatek-huytran2/go-sdk/types/msg"
)

// cdc global variable
var Cdc = amino.NewCodec()

func RegisterCodec(cdc *amino.Codec) {
	cdc.RegisterInterface((*Tx)(nil), nil)
	cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
	msg.RegisterCodec(cdc)
}

func init() {
	cryptoAmino.RegisterAmino(Cdc)
	RegisterCodec(Cdc)
}
