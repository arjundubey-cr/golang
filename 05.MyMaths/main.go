package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// using math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt((5)))
	fmt.Println(myRandomNum)
}
