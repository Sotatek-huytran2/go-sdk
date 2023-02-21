package transaction

import (
	"github.com/Sotatek-huytran2/go-sdk/types/msg"
	"github.com/Sotatek-huytran2/go-sdk/types/tx"
)

type ListMiniPairResult struct {
	tx.TxCommitResult
}

func (c *client) ListMiniPair(baseAssetSymbol string, quoteAssetSymbol string, initPrice int64, sync bool, options ...Option) (*ListMiniPairResult, error) {
	fromAddr := c.keyManager.GetAddr()

	listMsg := msg.NewListMiniMsg(fromAddr, baseAssetSymbol, quoteAssetSymbol, initPrice)
	commit, err := c.broadcastMsg(listMsg, sync, options...)
	if err != nil {
		return nil, err
	}

	return &ListMiniPairResult{*commit}, nil

}
