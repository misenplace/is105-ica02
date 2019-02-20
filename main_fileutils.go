package main

import (
	"fmt"

	"github.com/anzu/is105-ica02/fileutils"
)

func main() {
	l1 := fileutils.FileToByteslice("files/lang01.wl")
	l2 := fileutils.FileToByteslice("files/lang02.wl")
	l3 := fileutils.FileToByteslice("files/lang03.wl")

	fmt.Println("lang01.wl")
	fmt.Printf("% X\n", l1)
	fmt.Println("lang02.wl")
	fmt.Printf("% X\n", l2)
	fmt.Println("lang03.wl")
	fmt.Printf("% X\n", l3)
}