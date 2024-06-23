package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	solutionPart1 := Solution("ckczppom", 5)
	solutionPart2 := Solution("ckczppom", 6)
	fmt.Printf("Solution Part 1: %d\n", solutionPart1)
	fmt.Printf("Solution Part 2: %d\n", solutionPart2)
}

func Solution(key string, leadingZeros int) int {
	hasher := md5.New()

	for i := 1; ; i++ {
		hasher.Reset()
		currentTry := fmt.Sprintf("%s%d", key, i)
		_, err := io.WriteString(hasher, currentTry)
		if err != nil {
			fmt.Printf("Error creating MD5 hash of %s: %s", currentTry, err)
			os.Exit(1)
		}

		hash := hex.EncodeToString(hasher.Sum(nil))
		if strings.HasPrefix(hash, strings.Repeat("0", leadingZeros)) {
			return i
		}
	}
}
