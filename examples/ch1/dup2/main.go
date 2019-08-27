// dup2 prints the count and text of lines that appear more than once
// in the input. It reads from Stdin or from a list of names files

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]int)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(r io.Reader, counts map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}

	if input.Err() != nil {
		fmt.Fprintf(os.Stderr, "dup2: error while printing %v\n", input.Err())
	}
}
