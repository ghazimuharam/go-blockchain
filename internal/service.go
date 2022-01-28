// Package internal is root package for blockchain implementation
package internal

import (
	"github.com/ghzmhrm/go-blockchain/internal/entity"
)

// BlockchainSvc is an interface for holding blockchain service implementation
type BlockchainSvc interface {
	AddBlock(trxAmount uint64, sender string, receiver string, note string)
	GetBlocks() []*entity.Block
	InitializeGenesisBlock() *entity.Block
}
