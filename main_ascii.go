package main

import (
	"github.com/anzu/is105-ica02/ascii"
	"os"
	"fmt"
)

// constant var of the standard ASCII table
const asciiLiteral = "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f" +
	"\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f" +
	` !"#$%&'()*+,-./0123456789:;<=>?` +
	`@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_` +
	"`abcdefghijklmnopqrstuvwxyz{|}~\x7f"

// constant var of the extended ASCII table
const asciiExtendedLiteral = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89" +
	"\x8A\x8B\x8C\x8D\x8E\x8F" +
	"\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99" +
	"\x9A\x9B\x9C\x9D\x9E\x9F" +
	"\xA0\xA1\xA2\xA3\xA4\xA5\xA6\xA7\xA8\xA9" +
	"\xAA\xAB\xAC\xAD\xAE\xAF" +
	"\xB0\xB1\xB2\xB3\xB4\xB5\xB6\xB7\xB8\xB9" +
	"\xBA\xBB\xBC\xBD\xBE\xBF" +
	"\xC0\xC1\xC2\xC3\xC4\xC5\xC6\xC7\xC8\xC9" +
	"\xCA\xCB\xCC\xCD\xCE\xCF" +
	"\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8\xD9" +
	"\xDA\xDB\xDC\xDD\xDE\xDF" +
	"\xE0\xE1\xE2\xE3\xE4\xE5\xE6\xE7\xE8\xE9" +
	"\xEA\xEB\xEC\xED\xEE\xEF" +
	"\xF0\xF1\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9" +
	"\xFA\xFB\xFC\xFD\xFE\xFF"

func main() {

	// if arguments passed is less than 2, print usage
	// os.Args counts the actual program name (main_ascii.go) as an argument.
	// therefore, we have to check if it's more than two.
	// this check is required to not get an IndexOutOfBoundsException
	if len(os.Args) < 2 {
		printUsage()
	} else {
		// if second argument is -ex, print extended ascii table
		if os.Args[1] == "-ex" {
			printAsciiTable(asciiExtendedLiteral)
		// if second argument is -st, print standard ascii table
		} else if os.Args[1] == "-st" {
			printAsciiTable(asciiLiteral)
		// if second argument is nothing at all, print usage.
		} else {
			printUsage()
		}
	}
}

// make this a function for ease of use.
func printAsciiTable(stringLiteral string) {
	ascii.IterateOverASCIIStringLiteral(stringLiteral)
}

// usage function
func printUsage() {
	fmt.Println("usage: main_ascii [-st][-ex]")
	fmt.Println("-st	prints standard 0 to 127 ascii table")
	fmt.Println("-ex	prints exnteded ascii table")
}
