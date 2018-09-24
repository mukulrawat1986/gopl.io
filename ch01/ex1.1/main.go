// modify the echo program to also print the name of the command that
// invoked it.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Name of command: ", os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
