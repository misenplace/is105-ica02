package main

import (
	"fmt"
	"github.com/anzu/is105-ica02/slice"
)

func main() {
	byteslice1 := slice.AllocateMake(512)
	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteslice1[0])
	aslice := slice.Reslice(byteslice1, 50, 100)
	fmt.Println("&aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteslice1[0])
}