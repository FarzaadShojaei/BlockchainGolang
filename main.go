package main

//Describing The Folder Our Main File Is Nested In as a package
import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	timestamp    time.Time
	transactions []string
	prevhash     []byte
	Hash         []byte //byte:unsigned 8-bit integer,
	//Can Represent a character from a string
}

func main() {
	abc := []string{"A sent  50 Coins to BC"}
	xyz := Blocks(abc, []byte{})
	fmt.Println("This is Our First Block")
	Print(xyz)

	pqrs := []string{"PQ sent 230 coins to RC"}
	klmn := Blocks(pqrs, xyz.Hash)
	fmt.Println("This is our Second Block")
	Print(klmn)
}

func Blocks(transactions []string, prevhash []byte) *Block {
	currentTime := time.Now()
	return &Block{ //pointed to Struct
		timestamp:    currentTime,
		transactions: transactions,
		prevhash:     prevhash,
		Hash:         NewHash(currentTime, transactions, prevhash),
	}
}
func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, time.String()...)
	for transaction := range transactions {
		input = append(input, string(rune(transaction))...)

	}
	hash := sha256.Sum256(input)
	return hash[:]
}

func Print(block *Block) {
	fmt.Printf("\ttime:%s\n", block.timestamp.String())
	fmt.Printf("\tprevhash:%x\n", block.prevhash)
	fmt.Printf("\thash:%x\n", block.Hash)
	Transaction(block)

}
func Transaction(block *Block) {
	fmt.Println("\tTransactions: ")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v:%q\n", i, transaction)
	}
}
