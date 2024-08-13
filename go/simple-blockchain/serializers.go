package main

import (
	"encoding/gob"
	"os"
)

func SaveBlockchain(bc *BlockChain, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(bc)
}

func LoadBlockchain(filename string) (*BlockChain, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bc BlockChain
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&bc)
	if err != nil {
		return nil, err
	}
	return &bc, nil
}
