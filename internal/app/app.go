// Package app hold blockchain app initializer
package app

import (
	"fmt"
)

// InitBlockChain to initialize blockchain application
func InitBlockChain() {
	blockChainServices := InitializeBlockchain()

	// nolint:gomnd //example transaction, should not be linted
	blockChainServices.AddBlock(
		20, "ghz", "nakomoto", "test transaction 1")
	// nolint:gomnd //example transaction, should not be linted
	blockChainServices.AddBlock(
		32, "sabini", "tommy", "test transaction 2")
	// nolint:gomnd //example transaction, should not be linted
	blockChainServices.AddBlock(
		64, "tokyo", "nairobi", "test transaction 3")

	for _, block := range blockChainServices.GetBlocks() {
		fmt.Printf("Hash : %x\nPrevHash : %x\nBlockData : %+v\n\n", block.Hash, block.PrevHash, block.Data)
	}
}
