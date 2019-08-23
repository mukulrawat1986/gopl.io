package main

import (
	"fmt"
	"io"
	"strings"
)

func echo2(args []string, w io.Writer) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(w, s)
}

func echo3(args []string, w io.Writer) {
	fmt.Fprintln(w, strings.Join(args, " "))
}

func main() {

}
