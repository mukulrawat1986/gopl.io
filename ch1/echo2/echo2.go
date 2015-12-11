// Echo2 prints its command-line arguments
// It uses another form of the for-loop and iterates over a
// range of values from a data type like string or a slice.

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
