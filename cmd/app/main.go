package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp                 int64  // timestamp of when block is created
	Data, PrevBlockHash, Hash []byte // data contains valuable information, prevBlockHash stored prev block hash, Hash is the hash of this block
}

type Blockchain struct {
	blocks []*Block // an array with each element of "pointer to block" data type
}

func (b *Block) SetHash() {
	var timestamp = []byte(strconv.FormatInt(b.Timestamp, 10))                       // FormatInt converts Timestamp to decimal number
	var headers = bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // all properties other than current block hash concatenated
	var hash = sha256.Sum256(headers)                                                // hash of current block generated

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	var block = Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}} // block stores address of newly created(malloced) block
	block.SetHash()
	return &block // new block address returned (pointer to new block)
}

func (bc *Blockchain) AddBlock(data string) {
	var prevBlock = bc.blocks[len(bc.blocks)-1] // last block (before inserted of new block)
	var newBlock = NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // creates a Genesis block using malloc
}

func NewBlockchain() *Blockchain {
	var blockchain = Blockchain{[]*Block{NewGenesisBlock()}}
	return &blockchain
}

func main() {
	var bc = NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
