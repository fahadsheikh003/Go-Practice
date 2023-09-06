package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func isSolution(hash string, pattern string) bool {
	length := len(pattern)
	firstLengthChars := hash[:length]

	if pattern == firstLengthChars {
		return true
	}
	return false
}

func solvePuzzle(pattern string, sequence int, data string) (string, int) {
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

func main() {
	hash, nonce := solvePuzzle("0", 4, "")
	fmt.Printf("Hash: %v\nNonce: %v", hash, nonce)
}
