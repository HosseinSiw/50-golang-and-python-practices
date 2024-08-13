package main

import (
	"bytes"
	"crypto/sha256"
	"golang.org/x/crypto/blake2b"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

const difficulty = 4

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp, []byte(strconv.Itoa(b.Nonce))}, []byte{})

	// Using blacke2b function to hash the data.
	sh := sha256.Sum256(headers)
	hash := blake2b.Sum256(sh[:])
	b.Hash = hash[:]
}
func (b *Block) CalculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp, []byte(strconv.Itoa(b.Nonce))}, []byte{})
	sh := sha256.Sum256(headers)
	hash := blake2b.Sum256(sh[:])
	return hash[:]
}
func (b *Block) ProofOfWork() {
	target := big.NewInt(1)
	target.Lsh(target, 256-difficulty)

	var hashInt big.Int
	for {
		b.SetHash()
		hashInt.SetBytes(b.Hash)
		if hashInt.Cmp(target) == -1 {
			break
		}
		b.Nonce++
	}
}
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
		Nonce:         rand.Intn(1000000),
	}
	block.SetHash()
	block.ProofOfWork()
	return block
}
