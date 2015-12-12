// Modify dup2 to print the names of all files
// in which each duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	store := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, store, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, store, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			if len(files) != 0 {
				fmt.Printf("%s\t%d\t%s\n", store[line], n, line)
			} else {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, store map[string]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filename != "" {
			store[input.Text()] = filename
		}
	}
	// Note: ignoring potential errors from input.Err()
}
