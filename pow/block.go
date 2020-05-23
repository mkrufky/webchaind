// Copyright 2014 The go-ethereum Authors
// Copyright 2018 Webchain project
// This file is part of Webchain.
//
// Webchain is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Webchain is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Webchain. If not, see <http://www.gnu.org/licenses/>.

package pow

import (
	"math/big"

	"github.com/mkrufky/webchaind/common"
	"github.com/mkrufky/webchaind/core/types"
)

type Block interface {
	Difficulty() *big.Int
	HashNoNonce() common.Hash
	Nonce() uint64
	NumberU64() uint64
	Header() *types.Header
}

type ChainManager interface {
	GetBlockByNumber(uint64) *types.Block
	CurrentBlock() *types.Block
}
