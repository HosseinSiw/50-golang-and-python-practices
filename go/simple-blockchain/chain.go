package main

import "bytes"

type BlockChain struct {
	Blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlocks := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlocks.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *BlockChain) BcValidation() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		prevBlock := bc.Blocks[i-1]
		currBlock := bc.Blocks[i]

		// Check if current block's hash is correct
		if !bytes.Equal(currBlock.Hash, currBlock.CalculateHash()) {
			return false
		}
		// Check if previous block's hash is correct
		if !bytes.Equal(currBlock.PrevBlockHash, prevBlock.Hash) {
			return false
		}
	}
	return true
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
