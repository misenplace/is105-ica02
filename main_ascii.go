package main

import (
	"github.com/anzu/is105-ica02/ascii"
	"os"
	"fmt"
)

func main() {

	// if arguments passed is less than 2, print usage
	// os.Args counts the actual program name (main_ascii.go) as an argument.
	// therefore, we have to check if it's more than two.
	// this check is required to not get an IndexOutOfBoundsException
	if len(os.Args) < 2 {
		printUsage()
	} else {
//		// if second argument is -ex, print extended ascii table
//		if os.Args[1] == "-ex" {
//			ascii.IterateOverExtendedASCIIStringLiteral(ascii.GetExtendedASCIIStringLiteral())
//		// if second argument is -st, print standard ascii table
//		} else 
		if os.Args[1] == "-st" {
			ascii.IterateOverASCIIStringLiteral(ascii.GetStandardASCIIStringLiteral())
		// if second argument is nothing at all, print usage.
		} else {
			printUsage()
		}
	}
}

// usage function
func printUsage() {
	fmt.Println("usage: main_ascii [-st][-ex]")
	fmt.Println("-st	prints standard 0 to 127 ascii table")
	fmt.Println("-ex	prints exnteded ascii table")
}
