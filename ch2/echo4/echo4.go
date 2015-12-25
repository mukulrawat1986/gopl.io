// Echo4 prints its command line arguments
// It takes two optional flags: -n causes echo to omit the
// trailing newline that would normally be printed, and -s sep
// causes it to separate the output arguments by the content of the
// string sep instead of the default single space.

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
