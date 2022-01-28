// Package service hold service layer implementation
package service

import (
	"math"
	"math/big"

	"github.com/ghzmhrm/go-blockchain/internal/entity"
)

const (
	// BlockDifficulty determine the difficulty of block mining
	BlockDifficulty = 12
)

// BlockChainServices is a struct for holding
// method implementation of blockchain service
type BlockChainServices struct {
	BlockChain *entity.BlockChain
}

// NewBlockChainServices initialize blockchain service
// for dependency injection
func NewBlockChainServices() *BlockChainServices {
	// initialize int for target
	minimumHashTarget := big.NewInt(1)
	// nolint:gomnd // expected to be 256
	minimumHashTarget.Lsh(minimumHashTarget, uint(256-BlockDifficulty))

	blockChainServices := &BlockChainServices{
		BlockChain: &entity.BlockChain{
			Difficulty:        BlockDifficulty,
			MinimumHashTarget: minimumHashTarget,
		},
	}

	// Initialize Genesis Block
	block := blockChainServices.InitializeGenesisBlock()
	blockChainServices.BlockChain.Blocks = append(blockChainServices.BlockChain.Blocks, block)

	return blockChainServices
}

// AddBlock is a method for adding new block to blockchain
func (c *BlockChainServices) AddBlock(
	trxAmount uint64, sender string, receiver string, note string,
) {
	blockData := &entity.BlockData{
		Amount:   trxAmount,
		Sender:   sender,
		Receiver: receiver,
		Note:     note,
	}
	prevBlock := c.BlockChain.Blocks[len(c.BlockChain.Blocks)-1]

	newBlock := c.createBlock(prevBlock.Hash, blockData)
	c.BlockChain.Blocks = append(c.BlockChain.Blocks, newBlock)
}

// GetBlocks will return the entire blocks in blockchain struct
func (c *BlockChainServices) GetBlocks() []*entity.Block {
	return c.BlockChain.Blocks
}

// InitializeGenesisBlock to init the first block of the blockchain
func (c *BlockChainServices) InitializeGenesisBlock() *entity.Block {
	block := &entity.Block{
		Data: &entity.BlockData{
			Note: "Genesis Block",
		},
	}
	c.mine(block)
	return block
}

func (c *BlockChainServices) createBlock(
	prevHash []byte,
	blockData *entity.BlockData,
) *entity.Block {
	block := &entity.Block{
		PrevHash: prevHash,
		Data: &entity.BlockData{
			Amount:   blockData.Amount,
			Sender:   blockData.Sender,
			Receiver: blockData.Receiver,
			Note:     blockData.Note,
		},
	}
	c.mine(block)

	return block
}

func (c *BlockChainServices) mine(b *entity.Block) {
	var intHash big.Int

	nonce := 0

	for nonce < math.MaxInt64 {
		b.Nonce = uint64(nonce)
		intHash.SetBytes(b.CalculateHash())
		if intHash.Cmp(c.BlockChain.MinimumHashTarget) == -1 {
			break
		} else {
			nonce++
		}
	}
}
