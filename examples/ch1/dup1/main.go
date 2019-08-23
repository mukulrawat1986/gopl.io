// dup1 prints the test of each line that appears more than once
// in the standard input, prceeded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
