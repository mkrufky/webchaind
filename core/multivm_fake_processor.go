// +build !sputnikvm

package core

import (
	"math/big"

	"github.com/mkrufky/webchaind/core/state"
	"github.com/mkrufky/webchaind/core/types"
	evm "github.com/mkrufky/webchaind/core/vm"
)

const SputnikVMExists = false

var UseSputnikVM = "false"

func ApplyMultiVmTransaction(config *ChainConfig, bc *BlockChain, gp *GasPool, statedb *state.StateDB, header *types.Header, tx *types.Transaction, totalUsedGas *big.Int) (*types.Receipt, evm.Logs, *big.Int, error) {
	panic("not implemented")
}
