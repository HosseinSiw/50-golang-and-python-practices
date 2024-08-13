package main

import (
	"fmt"
)

func main() {
	blockchain, err := LoadBlockchain("sample")
	if err != nil {
		return
	}
	for index, block := range blockchain.Blocks {
		fmt.Printf("%d. Data -> %s\n", index, block.Data)
		fmt.Printf("Hash %x\n", block.Hash)
	}
}
