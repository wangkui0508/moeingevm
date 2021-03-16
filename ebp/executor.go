package ebp

import (
	gethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/smartbch/moeingevm/types"
)

type TxExecutor interface {
	//step 1: for deliverTx, collect block txs in engine.txList
	CollectTx(tx *gethtypes.Transaction)
	//step 2: for commit, check sig, insert regular txs standbyTxQ
	Prepare()
	//step 3: for postCommit, parallel execute tx in standbyTxQ
	Execute(currBlock *types.BlockInfo)

	//set context
	SetContext(ctx *types.Context)
	Context() *types.Context
	CollectTxsCount() int
	CommittedTxs() []*types.Transaction
}
