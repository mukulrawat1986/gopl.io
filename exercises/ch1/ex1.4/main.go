package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var mainMap = map[string]map[string]int{}

	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		mainMap["Args"] = counts
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts := make(map[string]int)
			countLines(f, counts)
			mainMap[filename] = counts
		}
	}

	for filename, counts := range mainMap {
		if filename != "Args" {
			fmt.Println(filename)
		}

		for line, n := range counts {
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
		fmt.Fprintf(os.Stderr, "dup2: error while scanning %v\n", input.Err())
	}
}
