// non-recursive version of comma
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("1"))
	fmt.Println(comma("12"))
	fmt.Println(comma("213"))
	fmt.Println(comma("3451"))
	fmt.Println(comma("23451"))
	fmt.Println(comma("234543"))
	fmt.Println(comma("1345678"))
}

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	count := len(s)

	buf := &bytes.Buffer{}

	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		count--
		if count%3 == 0 && count != 0 {
			buf.WriteRune(',')
		}
	}
	return buf.String()
}
