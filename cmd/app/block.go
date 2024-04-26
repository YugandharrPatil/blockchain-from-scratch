package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp                 int64  // timestamp of when block is created
	Data, PrevBlockHash, Hash []byte // data contains valuable information, prevBlockHash stored prev block hash, Hash is the hash of this block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // FormatInt converts Timestamp to decimal number as string
	// []byte() converts the string to a byte slice like [104 101 108 108 111]
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // all properties other than current block hash concatenated
	hash := sha256.Sum256(headers)                                                // hash of current block generated

	b.Hash = hash[:] // same as b.Hash = hash
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}} // newly malloced block
	block.SetHash()
	return &block // new block address returned (pointer to new block)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}) // creates a Genesis block using malloc
}
