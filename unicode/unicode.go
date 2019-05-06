package unicode

import (
  "fmt"
)

func Translate(expression string, language string) string {

	s := ""

	if expression == "nord og sør" {
		if language == "is" {
			s = "\x22" + expression + "\x22" + " på japansk er " + "\x22\xE5\x8C\x97\xE3\x81\xA8\xE5\x8D\x97\x22"
		} else if language == "jp" {
			s = "\x22" + expression + "\x22" + " på islandsk er " + "\x22\x6E\x6F\x72\xC3\xB0\x75\x72\x20\x6F\x67\x20\x73\x75\xC3\xB0\x75\x72\x22"
		}
	}
	return s
}

func UnicodeCodePointDemo() {
	fmt.Println("\xf0\x9F\x98\x80")
	fmt.Println("\xf0\x9F\x98\x97")
	fmt.Println("\xf0\x9F\x98")
	fmt.Println("\xf0\x9F")
	fmt.Println("\xf0")
}
