package main

import (
	"fmt"
)

const s = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%+q\n", s)
	fmt.Printf("%c\n", s)
	fmt.Printf("%s\n", []byte(s))
}