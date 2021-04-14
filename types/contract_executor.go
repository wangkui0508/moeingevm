package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type SystemContractExecutor interface {
	Init(ctx *Context)
	IsSystemContract(addr common.Address) bool
	Execute(context Context, currBlock *BlockInfo, tx *TxToRun) (status int, logs []EvmLog, gasUsed uint64, outData []byte)
}