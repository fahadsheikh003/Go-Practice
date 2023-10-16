package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Transaction struct {
	TransactionID              string
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

type Block struct {
	PreviousHash    string
	Previous        *Block
	TransactionPool []*Transaction
	Nonce           int
	Hash            string
	Timestamp       int64
}

type BlockChain struct {
	Blocks []*Block
}

func CalculateHash(strings ...string) string {
	var sha256String string
	for _, str := range strings {
		sha256String += str
	}
	sha256Bytes := sha256.Sum256([]byte(sha256String))
	return hex.EncodeToString(sha256Bytes[:])
}

func SolvePuzzle(pattern string, sequence int, data string) (string, int) {
	isSolution := func(hash string, pattern string) bool {
		length := len(pattern)
		firstLengthChars := hash[:length]

		if pattern == firstLengthChars {
			return true
		}
		return false
	}

	requiredPattern := strings.Repeat(pattern, sequence)
	for {
		nonce := rand.Intn(1000000)
		newData := strconv.Itoa(nonce) + data
		hashInBytes := sha256.Sum256([]byte(newData))
		hash := hex.EncodeToString(hashInBytes[:])
		if isSolution(hash, requiredPattern) {
			return hash, nonce
		}
	}
}

func (block *Block) Mine() {
	var transactionsInfo string
	for _, transaction := range block.TransactionPool {
		transactionsInfo += transaction.TransactionID + transaction.SenderBlockchainAddress + transaction.RecipientBlockchainAddress + fmt.Sprint(transaction.Value)
	}
	blockInfo := block.PreviousHash + transactionsInfo
	hash, nonce := SolvePuzzle("0", 2, blockInfo)
	block.Hash = hash
	block.Nonce = nonce
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	hash := CalculateHash(sender, recipient, fmt.Sprint(value))
	transaction := Transaction{TransactionID: hash, SenderBlockchainAddress: sender, RecipientBlockchainAddress: recipient, Value: value}
	return &transaction
}

func (blockChain *BlockChain) NewEmptyBlock(previousHash string, previousBlock *Block) {
	block := Block{PreviousHash: previousHash, Previous: previousBlock, TransactionPool: []*Transaction{}, Nonce: 0}
	blockChain.Blocks = append(blockChain.Blocks, &block)
}

func (blockChain *BlockChain) AddTransaction(sender string, recipient string, value float32) {
	if len(blockChain.Blocks) == 0 {
		transaction := NewTransaction(sender, recipient, value)
		blockChain.NewEmptyBlock("", nil)
		blockChain.Blocks[0].TransactionPool = append(blockChain.Blocks[0].TransactionPool, transaction)
	} else {
		block := blockChain.Blocks[len(blockChain.Blocks)-1]
		if len(block.TransactionPool) < 4 {
			transaction := NewTransaction(sender, recipient, value)
			block.TransactionPool = append(block.TransactionPool, transaction)
		} else {
			block.Mine()
			blockChain.NewEmptyBlock(block.Hash, block)

			block := blockChain.Blocks[len(blockChain.Blocks)-1]

			transaction := NewTransaction(sender, recipient, value)
			block.TransactionPool = append(block.TransactionPool, transaction)
		}
	}
}

func (transaction *Transaction) TransactionToJson() string {
	transactionJson, _ := json.MarshalIndent(transaction, "", "    ")
	return string(transactionJson)
}

func (block *Block) BlockToJson() string {
	blockJson, _ := json.MarshalIndent(block, "", "    ")
	return string(blockJson)
}

func JsonToTransaction(transactionJson string) *Transaction {
	transaction := Transaction{}
	json.Unmarshal([]byte(transactionJson), &transaction)
	return &transaction
}

func JsonToBlock(blockJson string) *Block {
	block := Block{}
	json.Unmarshal([]byte(blockJson), &block)
	return &block
}

func (blockChain *BlockChain) PrintBlockChain() {
	for index, block := range blockChain.Blocks {
		fmt.Println(strings.Repeat("=", 50), "Block ", index+1, strings.Repeat("=", 50))
		fmt.Printf("%s\n", block.BlockToJson())
	}
}

func main() {
	blockChain := BlockChain{}

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.AddTransaction("B", "C", 2.0)
	blockChain.AddTransaction("C", "D", 3.0)
	blockChain.AddTransaction("D", "E", 4.0)
	blockChain.AddTransaction("E", "F", 5.0)
	blockChain.AddTransaction("F", "G", 6.0)
	blockChain.AddTransaction("G", "H", 7.0)

	// for _, block := range blockChain.Blocks {
	// 	fmt.Println("Previous Hash: ", block.PreviousHash)
	// 	fmt.Println("Hash: ", block.Hash)
	// 	fmt.Println("Nonce: ", block.Nonce)
	// 	fmt.Println("Transaction Pool: ")
	// 	for _, transaction := range block.TransactionPool {
	// 		fmt.Println("Transaction ID: ", transaction.TransactionID)
	// 		fmt.Println("Sender: ", transaction.SenderBlockchainAddress)
	// 		fmt.Println("Recipient: ", transaction.RecipientBlockchainAddress)
	// 		fmt.Println("Value: ", transaction.Value)
	// 		fmt.Println()
	// 	}
	// 	fmt.Println()
	// }

	blockChain.PrintBlockChain()
}
