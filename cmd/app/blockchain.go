package main

type Blockchain struct {
	blocks []*Block // a slice with each element of "pointer to block" data type
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1] // last block (before inserted of new block)
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	blockchain := Blockchain{[]*Block{NewGenesisBlock()}}
	return &blockchain
}
