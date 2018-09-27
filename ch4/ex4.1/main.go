//  Function that counts the number of bits that are different in two SHA256 hashes
package main

import (
	"crypto/sha256"
	"fmt"
)

func popcount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= (b - 1)
	}
	return count
}

func main() {

	// create two SHA256 hashes and print them
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("c1 = % b\nc2 = % b\n%[1]T\n%[2]T\n", c1, c2)

	count := 0

	for i := 0; i < 32; i++ {
		fmt.Printf("% b  ^ % b = % b\n", c1[i], c2[i], c1[i]^c2[i])
		count += popcount(c1[i] ^ c2[i])
	}
	fmt.Println(count)
}
