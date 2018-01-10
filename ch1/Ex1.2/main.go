// Ex1.2 print index and value of arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", i, arg)
	}
}
