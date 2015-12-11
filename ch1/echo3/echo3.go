// Echo program to print the command line arguments
// here we use the Joins function from the strings package to
// get a simple and efficient solution.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	//  if we don't care about format but just want to see the values,
	// perhaps for debugging.
	// fmt.Println(os.Args[1:])
}
