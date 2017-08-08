package core

import (
	"github.com/ethereumproject/go-ethereum/logger"
)

var mlogBlockchain = logger.MLogPostComponent("blockchain", MLogLines)

// MLogLines is an exported slice of all available mlog LINES.
// May be used for automatic mlog documentation generator, or
// for API usage/display/documentation otherwise.
var MLogLines = []logger.MLogT{
	mlogBlockchainWriteBlock,
	mlogBlockchainInsertBlocks,
}

// Collect and document available mlog lines.

var mlogBlockchainWriteBlock = logger.MLogT{
	Description: "Called when a single block is added to the chaindb without error.",
	Receiver: "BLOCKCHAIN",
	Verb: "WRITE",
	Subject: "BLOCK",
	Details: []logger.MLogDetailT{
		{"WRITE", "STATUS", "STRING"},
		{"WRITE", "ERROR", "STRING"},
		{"BLOCK", "NUMBER", "INT"},
		{"BLOCK", "HASH", "STRING"},
		{"BLOCK", "SIZE", "INT"},
		{"BLOCK", "TRANSACTIONS_COUNT", "INT"},
		{"BLOCK", "GAS_USED", "INT"},
		{"BLOCK", "COINBASE", "STRING"},
		{"BLOCK", "TIME", "INT"},
	},
}


var mlogBlockchainInsertBlocks = logger.MLogT{
	Description: "Called when a chain of blocks is inserted into our client's chain db.",
	Receiver: "BLOCKCHAIN",
	Verb: "INSERT",
	Subject: "BLOCKS",
	Details: []logger.MLogDetailT{
		{"BLOCKS", "COUNT", "INT"},
		{"BLOCKS", "QUEUED", "INT"},
		{"BLOCKS", "IGNORED", "INT"},
		{"BLOCKS", "TRANSACTIONS_COUNT", "INT"},
		{"BLOCKS", "LAST_NUMBER", "INT"},
		{"BLOCKS", "FIRST_HASH", "STRING"},
		{"BLOCKS", "LAST_HASH", "STRING"},
		{"INSERT", "TIME", "INTERVAL"},
	},
}
