// M0dify the echo program to print the index and value of
// each of its arguments one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", index, value)
	}
}
