// Modify dup2 to print the name of all files in which each duplicate line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// commonFiles maps a line of text to name of all files that contain it.
	commonFiles := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, commonFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			defer f.Close()
			countLines(f, counts, commonFiles)
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\t%v\n", n, line, commonFiles[line])
	}
}

func countLines(f *os.File, counts map[string]int, commonFiles map[string][]string) {
	input := bufio.NewScanner(f)
	// name of the file
	name := f.Name()
	for input.Scan() {
		text := input.Text()
		counts[text]++

		// slice of named files
		files := commonFiles[text]
		// flag
		flag := 0

		if len(files) == 0 {
			commonFiles[text] = append(commonFiles[text], name)
		} else {
			for _, file := range files {
				if file == name {
					flag = 1
					break
				}
			}
			if flag == 0 {
				commonFiles[text] = append(commonFiles[text], name)
			}
		}
	}
}
