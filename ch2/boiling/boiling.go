// Boiling prints the boiling point of water

package main

import "fmt"

const boiling = 212.0

func main() {
	var f = boiling
	var c = (f - 32) * 5 / 9
	fmt.Printf("%gF or %gC\n", f, c)
}
