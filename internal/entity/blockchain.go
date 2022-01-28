// Package entity contains entity declaration and simple process related to the entity
package entity

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"math/big"
)

// BlockChain is a struct for blockchain entity
type BlockChain struct {
	Blocks            []*Block
	Difficulty        int
	MinimumHashTarget *big.Int
}

// Block is a struct for block entity
type Block struct {
	Hash     []byte
	Data     *BlockData
	PrevHash []byte
	Nonce    uint64
}

// BlockData is a struct for holding data
type BlockData struct {
	Amount   uint64 `json:"amount"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Note     string `json:"note"`
}

// CalculateHash will return a newly calculated hash by combining
// bytes of data, prev hash and nonce
func (b *Block) CalculateHash() []byte {
	dataByte, err := json.Marshal(b.Data)
	if err == nil {
		hash := bytes.Join([][]byte{dataByte, b.PrevHash, nonceToHex(int64(b.Nonce))}, []byte{})
		hashedData := sha256.Sum256(hash)
		b.Hash = hashedData[:]
	}

	return b.Hash
}

func nonceToHex(nonce int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, nonce)
	if err != nil {
		return nil
	}

	return buff.Bytes()
}
