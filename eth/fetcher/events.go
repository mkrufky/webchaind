package fetcher

import "github.com/mkrufky/webchaind/core/types"

type FetcherInsertBlockEvent struct {
	Peer  string
	Block *types.Block
}
