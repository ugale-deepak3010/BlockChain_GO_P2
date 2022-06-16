// Main.go project Main.go.go
package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp    time.Time
	transactions []string
	prevHash     []byte
	Hash         []byte
}

func main() {
	fmt.Println("\n ===== BLOCK CHAIN START ===== \n")
	abc := []string{" A sent 3 coins to BC"}
	xyz := Blocks(abc, []byte{})
	fmt.Println("This is First Block")
	Print(xyz)

	fmt.Println("\n---== 2 ==---\n")

	pqrs := []string{" pq sent 89.3444 coins to rs"}
	klmn := Blocks(pqrs, xyz.Hash)
	fmt.Println("This is Second Block")
	Print(klmn)

	fmt.Println("\n---== 3 ==---\n")

	opyu := []string{" op sent 30.5002 coins to yu"}
	asdu := Blocks(opyu, klmn.Hash)
	fmt.Println("This is Second Block")
	Print(asdu)

	fmt.Println("\n------------ THE END --------------\n")
}

func Blocks(transactions []string, prevhash []byte) *Block {
	currentTime := time.Now()

	return &Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevHash:     prevhash,
		Hash:         NewHash(currentTime, transactions, prevhash),
	}
}

func NewHash(time time.Time, transactions []string, prevhash []byte) []byte {
	input := append(prevhash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func Print(block *Block) {
	fmt.Printf("\ttime: %s \n", block.timestamp.String())
	fmt.Printf("\tprevhash: %x \n", block.prevHash)
	fmt.Printf("\thash: %x \n", block.Hash)
	Transaction(block)
}

func Transaction(block *Block) {
	fmt.Println("\tTransaction:")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}
