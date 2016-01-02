// Simple program to show tuple assignment in Go
// We create two functions here, once for gcd and other
// for fibonacci

package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	x := 10
	y := 14
	res := gcd(x, y)
	fmt.Printf("The gcd of %d and %d is %d\n", x, y, res)

	n := 11
	res = fib(n)
	fmt.Printf("The %dth fibonacci number is %d\n", n, res)
}
