// Echo program to print the command line arguments.
// Prints the index and value of each of its arguments
// one per line.

package main

import (
	"fmt"
	"os"
)

func main() {

	for i, arg := range os.Args[1:] {
		//		fmt.Printf("%d: %s\n", i, arg)
		fmt.Println(i, arg)
	}
}
