package main

import (
	"fmt"
	"github.com/anzu/is105-ica02/unicode"
)

func main() {
	jp := unicode.Translate("nord og sør", "jp")
	is := unicode.Translate("nord og sør", "is")

	fmt.Println(is)
	fmt.Println(jp)
}